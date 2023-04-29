package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// MOCK >
	//

	os.Setenv("ELECTION_NUM", "11")
	// azure blob storage connection string
	fmt.Printf("%s\n", os.Getenv("ELECTION_NUM"))

	//
	// < MOCK

	// check env var > ELECTION_NUM
	num, err := strconv.Atoi(os.Getenv("ELECTION_NUM"))
	if err != nil {
		log.Fatalln(err)
	}

	if err := validateElectionNum(&num); err != nil {
		log.Fatalln(err)
	}

	classifyElection(&num)

}

func validateElectionNum(num *int) error {
	if *num <= 0 || *num > len(Elections) {
		return fmt.Errorf(`Environment variable "ELECTION_NUM" not valid.\n`)
	}

	return nil
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
