package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func main() {
	log.Println("start program")

	// validation of env var "ELECTION_NUM"
	log.Println("validating env var...")
	num, err := strconv.Atoi(os.Getenv("ELECTION_NUM"))
	if err != nil {
		log.Fatalln(err)
	}

	if err := validateElectionNum(&num); err != nil {
		log.Fatalln(err)
	}
	log.Println("env var ok")

	// create tmp file
	log.Println("create tmp file")
	f, err := os.CreateTemp("", "profiles.jsonl")
	if err != nil {
		log.Fatalln(err)
	}
	defer os.Remove(f.Name())

	candidatesListChan := make(chan string, 64)
	profilesChan := make(chan string, 64)
	doneChan := make(chan bool, 1)
	go listCandidates(candidatesListChan, profilesChan)
	go scanProfile(profilesChan, doneChan, f)

	id, year, isFederal := classifyElection(&num)
	log.Printf("[election attrs] id: %v, year: %v, isFederal: %v\n", id, year, isFederal)

	if isFederal {

		go func() {
			for _, zone := range Zones {
				// in federal level: other zones uses codes between "3" and "10" (skipping "8")
				if zone != "BR" {
					for i := 3; i <= 10; i++ {
						if i == 8 {
							continue
						}

						candidatesListChan <- fmt.Sprintf("%v/candidatura/listar/%v/%v/%v/%d/candidatos", TSE, year, zone, id, i)
					}

					continue
				}

				// in federal level: "BR" uses codes "1" and "2"
				for i := 1; i <= 2; i++ {
					candidatesListChan <- fmt.Sprintf("%v/candidatura/listar/%v/%v/%v/%d/candidatos", TSE, year, zone, id, i)
				}
			}

			close(candidatesListChan)
		}()

	} else {

		citiesListChan := make(chan string, 64)
		go listCities(citiesListChan, candidatesListChan, year)

		go func() {
			for _, zone := range Zones {
				if zone == "BR" {
					continue
				}

				citiesListChan <- fmt.Sprintf("%v/eleicao/buscar/%v/%v/municipios", TSE, zone, id)
			}

			close(citiesListChan)
		}()

	}

	<-doneChan // block until done

	wd := os.Getenv("AZURE_STG_WD")
	container, path := resolveAzureVars(&wd)
	blobName := filepath.Join(path, fmt.Sprintf("election_%v.jsonl", id))

	blob, err := azblob.NewClientFromConnectionString(os.Getenv("AZURE_STG_CONN_STR"), nil)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = blob.UploadFile(context.TODO(), container, blobName, f, nil)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("end program")
}

func listCities(citiesListChan <-chan string, candidatesListChan chan<- string, year string) {
	var wg sync.WaitGroup

	for url := range citiesListChan {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			resp, err := http.Get(url)
			if err != nil {
				log.Println(err)
			}
			defer resp.Body.Close()

			var body Cidades
			if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
				log.Println(err)
			}

			// build candidates list endpt
			for _, c := range body.Municipios {
				suffix := url[68:]
				zone := suffix[:strings.Index(suffix, "/")]
				suffix = suffix[len(zone)+1:]
				election := suffix[:strings.Index(suffix, "/")]

				for i := 11; i <= 13; i++ {
					candidatesListChan <- fmt.Sprintf("%v/candidatura/listar/%v/%v/%v/%d/candidatos", TSE, year, c.Codigo, election, i)
				}
			}
		}(url)
	}

	wg.Wait()
	close(candidatesListChan)
}

func listCandidates(candidatesListChan <-chan string, profilesChan chan<- string) {
	var wg sync.WaitGroup

	for url := range candidatesListChan {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			resp, err := http.Get(url)
			if err != nil {
				log.Println(err)
			}
			defer resp.Body.Close()

			var body Candidaturas
			if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
				log.Println(err)
			}

			// build profiles endpt
			for _, c := range body.Candidatos {
				suffix := url[72:]
				year := suffix[:4]
				suffix = suffix[len(year)+1:]
				zone := suffix[:strings.Index(suffix, "/")]
				suffix = suffix[len(zone)+1:]
				election := suffix[:strings.Index(suffix, "/")]
				profilesChan <- fmt.Sprintf("%v/candidatura/buscar/%v/%v/%v/candidato/%d", TSE, year, zone, election, c.ID)
			}
		}(url)
	}

	wg.Wait()
	close(profilesChan)
}

func scanProfile(profilesChan <-chan string, doneChan chan<- bool, f *os.File) {
	var wg sync.WaitGroup

	for url := range profilesChan {
		wg.Add(2)

		go func(url string) {
			defer wg.Done()

			resp, err := http.Get(url)
			if err != nil {
				log.Println(err)
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
			}

			content := string(body)

			if content != "" {
				// write profile to tmp file
				go func() {
					log.Printf("requesting %q\n", url)
					defer wg.Done()
					_, err = f.WriteString(content + "\n")
					if err != nil {
						log.Println(err)
					}
				}()
			}
		}(url)
	}

	wg.Wait()
	doneChan <- true
}

func validateElectionNum(num *int) error {
	if *num <= 0 || *num > len(Elections) {
		return fmt.Errorf(`environment variable "ELECTION_NUM" not valid\n`)
	}

	return nil
}

func resolveAzureVars(wd *string) (container, path string) {
	idx := strings.Index(*wd, "/")
	container = (*wd)[:idx]
	path = (*wd)[idx+1:]
	return
}

func classifyElection(num *int) (id, year string, isFederal bool) {
	id = Elections[*num-1][1]
	idx := strings.LastIndex(Elections[*num-1][0], " ")
	year = Elections[*num-1][0][idx+1:]

	// new elections must be added here too
	switch *num {
	case 1, 3, 5, 7, 9:
		isFederal = false
	case 10:
		isFederal, year = false, "2020"
	case 2, 4, 6, 8, 11:
		isFederal = true
	}

	return
}
