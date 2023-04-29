package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

// when new elections occur it must be added to the slice
var Elections2 = [][]string{
	{"Eleições Municipais 2004", "14431"},
	{"Eleição Geral Federal 2006", "14423"},
	{"Eleições Municipais 2008", "14422"},
	{"Eleição Geral Federal 2010", "14417"},
	{"Eleições Municipais 2012", "1699"},
	{"Eleição Geral Federal 2014", "680"},
	{"Eleições Municipais 2016", "2"},
	{"Eleição Geral Federal 2018", "2022802018"},
	{"Eleições Municipais 2020", "2030402020"},
	{"Eleições Municipais 2020 - AP", "2032002020"},
	{"Eleição Geral Federal 2022", "2040602022"},
}

var Zones2 = [28]string{
	"BR",                                     // Federal
	"AC", "AM", "AP", "PA", "RO", "RR", "TO", // North
	"AL", "BA", "CE", "MA", "PB", "PE", "PI", "RN", "SE", // Northeast
	"DF", "GO", "MS", "MT", // Central-west
	"ES", "MG", "RJ", "SP", // Southeast
	"PR", "RS", "SC", // South
}

const TSEd string = "https://divulgacandcontas.tse.jus.br/divulga/rest/v1"

// JSON Structs
type Candidaturas struct {
	UnidadeEleitoral struct {
		ID         interface{} `json:"id"`
		Sigla      string      `json:"sigla"`
		Nome       string      `json:"nome"`
		Regiao     interface{} `json:"regiao"`
		Cargos     interface{} `json:"cargos"`
		Diretorios interface{} `json:"diretorios"`
		Codigo     string      `json:"codigo"`
		Capital    bool        `json:"capital"`
		Estado     interface{} `json:"estado"`
	} `json:"unidadeEleitoral"`
	Cargo struct {
		Codigo      int         `json:"codigo"`
		Sigla       interface{} `json:"sigla"`
		Nome        string      `json:"nome"`
		CodSuperior int         `json:"codSuperior"`
		Titular     bool        `json:"titular"`
		Contagem    int         `json:"contagem"`
	} `json:"cargo"`
	Candidatos []struct {
		ID                               int64       `json:"id"`
		NomeUrna                         string      `json:"nomeUrna"`
		Numero                           int         `json:"numero"`
		IDCandidatoSuperior              interface{} `json:"idCandidatoSuperior"`
		NomeCompleto                     string      `json:"nomeCompleto"`
		DescricaoSexo                    interface{} `json:"descricaoSexo"`
		DataDeNascimento                 interface{} `json:"dataDeNascimento"`
		TituloEleitor                    interface{} `json:"tituloEleitor"`
		Cpf                              interface{} `json:"cpf"`
		DescricaoEstadoCivil             interface{} `json:"descricaoEstadoCivil"`
		DescricaoCorRaca                 interface{} `json:"descricaoCorRaca"`
		DescricaoSituacao                string      `json:"descricaoSituacao"`
		Nacionalidade                    interface{} `json:"nacionalidade"`
		GrauInstrucao                    interface{} `json:"grauInstrucao"`
		Ocupacao                         interface{} `json:"ocupacao"`
		GastoCampanha1T                  interface{} `json:"gastoCampanha1T"`
		GastoCampanha2T                  interface{} `json:"gastoCampanha2T"`
		SgUfNascimento                   interface{} `json:"sgUfNascimento"`
		NomeMunicipioNascimento          interface{} `json:"nomeMunicipioNascimento"`
		LocalCandidatura                 interface{} `json:"localCandidatura"`
		UfCandidatura                    interface{} `json:"ufCandidatura"`
		UfSuperiorCandidatura            interface{} `json:"ufSuperiorCandidatura"`
		DataUltimaAtualizacao            interface{} `json:"dataUltimaAtualizacao"`
		FotoURL                          interface{} `json:"fotoUrl"`
		FotoDataUltimaAtualizacao        interface{} `json:"fotoDataUltimaAtualizacao"`
		DescricaoTotalizacao             string      `json:"descricaoTotalizacao"`
		NomeColigacao                    string      `json:"nomeColigacao"`
		ComposicaoColigacao              interface{} `json:"composicaoColigacao"`
		DescricaoTipoDrap                interface{} `json:"descricaoTipoDrap"`
		NumeroProcessoDrap               interface{} `json:"numeroProcessoDrap"`
		NumeroProcessoDrapEncrypt        interface{} `json:"numeroProcessoDrapEncrypt"`
		NumeroProcesso                   interface{} `json:"numeroProcesso"`
		NumeroProcessoEncrypt            interface{} `json:"numeroProcessoEncrypt"`
		NumeroProcessoPrestContas        interface{} `json:"numeroProcessoPrestContas"`
		NumeroProcessoPrestContasEncrypt interface{} `json:"numeroProcessoPrestContasEncrypt"`
		NumeroProtocolo                  interface{} `json:"numeroProtocolo"`
		Cargo                            struct {
			Codigo      int         `json:"codigo"`
			Sigla       interface{} `json:"sigla"`
			Nome        string      `json:"nome"`
			CodSuperior int         `json:"codSuperior"`
			Titular     bool        `json:"titular"`
			Contagem    int         `json:"contagem"`
		} `json:"cargo"`
		Bens        interface{} `json:"bens"`
		TotalDeBens interface{} `json:"totalDeBens"`
		Vices       interface{} `json:"vices"`
		Partido     struct {
			Numero int         `json:"numero"`
			Sigla  string      `json:"sigla"`
			Nome   interface{} `json:"nome"`
		} `json:"partido"`
		Eleicao struct {
			ID                       int         `json:"id"`
			SiglaUF                  interface{} `json:"siglaUF"`
			LocalidadeSgUe           interface{} `json:"localidadeSgUe"`
			Ano                      int         `json:"ano"`
			Codigo                   interface{} `json:"codigo"`
			NomeEleicao              interface{} `json:"nomeEleicao"`
			TipoEleicao              interface{} `json:"tipoEleicao"`
			Turno                    interface{} `json:"turno"`
			TipoAbrangencia          interface{} `json:"tipoAbrangencia"`
			DataEleicao              interface{} `json:"dataEleicao"`
			CodSituacaoEleicao       interface{} `json:"codSituacaoEleicao"`
			DescricaoSituacaoEleicao interface{} `json:"descricaoSituacaoEleicao"`
			DescricaoEleicao         string      `json:"descricaoEleicao"`
		} `json:"eleicao"`
		Emails                     interface{} `json:"emails"`
		Sites                      interface{} `json:"sites"`
		Arquivos                   interface{} `json:"arquivos"`
		EleicoesAnteriores         interface{} `json:"eleicoesAnteriores"`
		Substituto                 interface{} `json:"substituto"`
		Motivos                    interface{} `json:"motivos"`
		CodigoSituacaoCandidato    interface{} `json:"codigoSituacaoCandidato"`
		DescricaoSituacaoCandidato interface{} `json:"descricaoSituacaoCandidato"`
		IsCandidatoInapto          interface{} `json:"isCandidatoInapto"`
		CodigoSituacaoPartido      interface{} `json:"codigoSituacaoPartido"`
		DescricaoSituacaoPartido   interface{} `json:"descricaoSituacaoPartido"`
		IsCandFechado              interface{} `json:"isCandFechado"`
		DescricaoNaturalidade      string      `json:"descricaoNaturalidade"`
		StSUBSTITUIDO              interface{} `json:"st_SUBSTITUIDO"`
		Cnpjcampanha               interface{} `json:"cnpjcampanha"`
		GastoCampanha              float64     `json:"gastoCampanha"`
		StMOTIVOFICHALIMPA         interface{} `json:"st_MOTIVO_FICHA_LIMPA"`
		StMOTIVOABUSOPODER         interface{} `json:"st_MOTIVO_ABUSO_PODER"`
		StMOTIVOCOMPRAVOTO         interface{} `json:"st_MOTIVO_COMPRA_VOTO"`
		StMOTIVOCONDUTAVEDADA      interface{} `json:"st_MOTIVO_CONDUTA_VEDADA"`
		StMOTIVOGASTOILICITO       interface{} `json:"st_MOTIVO_GASTO_ILICITO"`
		DsMOTIVOOUTROS             interface{} `json:"ds_MOTIVO_OUTROS"`
		StMOTIVOAUSENCIAREQUISITO  interface{} `json:"st_MOTIVO_AUSENCIA_REQUISITO"`
		StMOTIVOINDPARTIDO         interface{} `json:"st_MOTIVO_IND_PARTIDO"`
		StDIVULGA                  interface{} `json:"st_DIVULGA"`
		StDIVULGABENS              interface{} `json:"st_DIVULGA_BENS"`
		StREELEICAO                bool        `json:"st_REELEICAO"`
		StDIVULGAARQUIVOS          interface{} `json:"st_DIVULGA_ARQUIVOS"`
	} `json:"candidatos"`
}

type PerfilCandidato struct {
	ID                               int         `json:"id"`
	NomeUrna                         string      `json:"nomeUrna"`
	Numero                           int         `json:"numero"`
	IDCandidatoSuperior              int         `json:"idCandidatoSuperior"`
	NomeCompleto                     string      `json:"nomeCompleto"`
	DescricaoSexo                    string      `json:"descricaoSexo"`
	DataDeNascimento                 string      `json:"dataDeNascimento"`
	TituloEleitor                    string      `json:"tituloEleitor"`
	Cpf                              string      `json:"cpf"`
	DescricaoEstadoCivil             string      `json:"descricaoEstadoCivil"`
	DescricaoCorRaca                 interface{} `json:"descricaoCorRaca"`
	DescricaoSituacao                string      `json:"descricaoSituacao"`
	Nacionalidade                    string      `json:"nacionalidade"`
	GrauInstrucao                    string      `json:"grauInstrucao"`
	Ocupacao                         string      `json:"ocupacao"`
	GastoCampanha1T                  interface{} `json:"gastoCampanha1T"`
	GastoCampanha2T                  interface{} `json:"gastoCampanha2T"`
	SgUfNascimento                   string      `json:"sgUfNascimento"`
	NomeMunicipioNascimento          string      `json:"nomeMunicipioNascimento"`
	LocalCandidatura                 string      `json:"localCandidatura"`
	UfCandidatura                    string      `json:"ufCandidatura"`
	UfSuperiorCandidatura            interface{} `json:"ufSuperiorCandidatura"`
	DataUltimaAtualizacao            string      `json:"dataUltimaAtualizacao"`
	FotoURL                          string      `json:"fotoUrl"`
	FotoDataUltimaAtualizacao        string      `json:"fotoDataUltimaAtualizacao"`
	DescricaoTotalizacao             string      `json:"descricaoTotalizacao"`
	NomeColigacao                    string      `json:"nomeColigacao"`
	ComposicaoColigacao              string      `json:"composicaoColigacao"`
	DescricaoTipoDrap                string      `json:"descricaoTipoDrap"`
	NumeroProcessoDrap               interface{} `json:"numeroProcessoDrap"`
	NumeroProcessoDrapEncrypt        interface{} `json:"numeroProcessoDrapEncrypt"`
	NumeroProcesso                   string      `json:"numeroProcesso"`
	NumeroProcessoEncrypt            string      `json:"numeroProcessoEncrypt"`
	NumeroProcessoPrestContas        interface{} `json:"numeroProcessoPrestContas"`
	NumeroProcessoPrestContasEncrypt interface{} `json:"numeroProcessoPrestContasEncrypt"`
	NumeroProtocolo                  string      `json:"numeroProtocolo"`
	Cargo                            struct {
		Codigo      int         `json:"codigo"`
		Sigla       interface{} `json:"sigla"`
		Nome        string      `json:"nome"`
		CodSuperior int         `json:"codSuperior"`
		Titular     bool        `json:"titular"`
		Contagem    int         `json:"contagem"`
	} `json:"cargo"`
	Bens []struct {
		Ordem                 int     `json:"ordem"`
		Descricao             string  `json:"descricao"`
		DescricaoDeTipoDeBem  string  `json:"descricaoDeTipoDeBem"`
		Valor                 float64 `json:"valor"`
		DataUltimaAtualizacao string  `json:"dataUltimaAtualizacao"`
	} `json:"bens"`
	TotalDeBens float64 `json:"totalDeBens"`
	Vices       []struct {
		DtUltimaAtualizacao string      `json:"DT_ULTIMA_ATUALIZACAO"`
		NomeColigacao       interface{} `json:"nomeColigacao"`
		ComposicaoColigacao interface{} `json:"composicaoColigacao"`
		StRegistro          interface{} `json:"stRegistro"`
		SituacaoCandidato   interface{} `json:"situacaoCandidato"`
		URLFoto             string      `json:"urlFoto"`
		DtULTIMAATUALIZACAO int64       `json:"dt_ULTIMA_ATUALIZACAO"`
		SqCANDIDATO         int         `json:"sq_CANDIDATO"`
		SgUE                string      `json:"sg_UE"`
		SqCANDIDATOSUPERIOR interface{} `json:"sq_CANDIDATO_SUPERIOR"`
		NrCANDIDATO         string      `json:"nr_CANDIDATO"`
		NmURNA              string      `json:"nm_URNA"`
		NmCANDIDATO         string      `json:"nm_CANDIDATO"`
		DsCARGO             string      `json:"ds_CARGO"`
		NmPARTIDO           string      `json:"nm_PARTIDO"`
		SgPARTIDO           string      `json:"sg_PARTIDO"`
		SqELEICAO           int         `json:"sq_ELEICAO"`
	} `json:"vices"`
	Partido struct {
		Numero int    `json:"numero"`
		Sigla  string `json:"sigla"`
		Nome   string `json:"nome"`
	} `json:"partido"`
	Eleicao struct {
		ID                       int         `json:"id"`
		SiglaUF                  interface{} `json:"siglaUF"`
		LocalidadeSgUe           interface{} `json:"localidadeSgUe"`
		Ano                      int         `json:"ano"`
		Codigo                   interface{} `json:"codigo"`
		NomeEleicao              interface{} `json:"nomeEleicao"`
		TipoEleicao              interface{} `json:"tipoEleicao"`
		Turno                    interface{} `json:"turno"`
		TipoAbrangencia          interface{} `json:"tipoAbrangencia"`
		DataEleicao              interface{} `json:"dataEleicao"`
		CodSituacaoEleicao       interface{} `json:"codSituacaoEleicao"`
		DescricaoSituacaoEleicao interface{} `json:"descricaoSituacaoEleicao"`
		DescricaoEleicao         string      `json:"descricaoEleicao"`
	} `json:"eleicao"`
	Emails             []interface{} `json:"emails"`
	Sites              []interface{} `json:"sites"`
	Arquivos           []interface{} `json:"arquivos"`
	EleicoesAnteriores []struct {
		NrAno               int    `json:"nrAno"`
		ID                  string `json:"id"`
		NomeUrna            string `json:"nomeUrna"`
		NomeCandidato       string `json:"nomeCandidato"`
		IDEleicao           string `json:"idEleicao"`
		SgUe                string `json:"sgUe"`
		Local               string `json:"local"`
		Cargo               string `json:"cargo"`
		Partido             string `json:"partido"`
		SituacaoTotalizacao string `json:"situacaoTotalizacao"`
		TxLink              string `json:"txLink"`
	} `json:"eleicoesAnteriores"`
	Substituto                 interface{} `json:"substituto"`
	Motivos                    interface{} `json:"motivos"`
	CodigoSituacaoCandidato    int         `json:"codigoSituacaoCandidato"`
	DescricaoSituacaoCandidato interface{} `json:"descricaoSituacaoCandidato"`
	IsCandidatoInapto          bool        `json:"isCandidatoInapto"`
	CodigoSituacaoPartido      string      `json:"codigoSituacaoPartido"`
	DescricaoSituacaoPartido   string      `json:"descricaoSituacaoPartido"`
	IsCandFechado              bool        `json:"isCandFechado"`
	DescricaoNaturalidade      string      `json:"descricaoNaturalidade"`
	StSUBSTITUIDO              bool        `json:"st_SUBSTITUIDO"`
	StMOTIVOAUSENCIAREQUISITO  bool        `json:"st_MOTIVO_AUSENCIA_REQUISITO"`
	StMOTIVOINDPARTIDO         bool        `json:"st_MOTIVO_IND_PARTIDO"`
	StDIVULGA                  bool        `json:"st_DIVULGA"`
	StDIVULGABENS              bool        `json:"st_DIVULGA_BENS"`
	StREELEICAO                bool        `json:"st_REELEICAO"`
	StDIVULGAARQUIVOS          bool        `json:"st_DIVULGA_ARQUIVOS"`
	StMOTIVOFICHALIMPA         bool        `json:"st_MOTIVO_FICHA_LIMPA"`
	StMOTIVOABUSOPODER         bool        `json:"st_MOTIVO_ABUSO_PODER"`
	StMOTIVOCOMPRAVOTO         bool        `json:"st_MOTIVO_COMPRA_VOTO"`
	StMOTIVOCONDUTAVEDADA      bool        `json:"st_MOTIVO_CONDUTA_VEDADA"`
	StMOTIVOGASTOILICITO       bool        `json:"st_MOTIVO_GASTO_ILICITO"`
	DsMOTIVOOUTROS             interface{} `json:"ds_MOTIVO_OUTROS"`
	Cnpjcampanha               interface{} `json:"cnpjcampanha"`
	GastoCampanha              float64     `json:"gastoCampanha"`
}

type Cidades struct {
	Estado struct {
		ID         interface{} `json:"id"`
		Sigla      string      `json:"sigla"`
		Nome       string      `json:"nome"`
		Regiao     interface{} `json:"regiao"`
		Cargos     interface{} `json:"cargos"`
		Diretorios interface{} `json:"diretorios"`
		Codigo     string      `json:"codigo"`
		Capital    bool        `json:"capital"`
		Estado     string      `json:"estado"`
	} `json:"estado"`
	Municipios []struct {
		ID         int         `json:"id"`
		Sigla      interface{} `json:"sigla"`
		Nome       string      `json:"nome"`
		Regiao     interface{} `json:"regiao"`
		Cargos     interface{} `json:"cargos"`
		Diretorios interface{} `json:"diretorios"`
		Codigo     string      `json:"codigo"`
		Capital    bool        `json:"capital"`
		Estado     interface{} `json:"estado"`
	} `json:"municipios"`
}

func main2() {
	fmt.Println("Eleições cadastradas no Tribunal Superior Eleitoral (TSE)")

	for i, v := range Elections {
		i++
		fmt.Printf("%v. %v\n", i, v[0])
	}

	input, err := strconv.Atoi(os.Getenv("ElectionID"))
	if err != nil {
		log.Fatalln(err)
	}

	isFederal, id, year := buildEndpoint(uint8(input), Elections)
	fmt.Println("Iniciando extração de dados sobre as candidaturas ...")

	// file handler
	var wg sync.WaitGroup

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	cwd = fmt.Sprintf("%v/%v.jsonl", cwd, Elections[input-1][0])

	f, err := os.Create(cwd)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	if isFederal {
		for _, zone := range Zones {
			if zone == "BR" {
				// in federal level: "BR" uses codes "1" and "2"
				for i := 1; i <= 2; i++ {
					endpt := fmt.Sprintf("%v/candidatura/listar/%v/%v/%v/%v/candidatos", TSE, year, zone, id, i)
					urlChan := sendURL(listCandidatosID(endpt), year, zone, id)

					for url := range urlChan {
						respChan := get(url)
						bodyChan := readResponse(<-respChan)

						var prf PerfilCandidato
						if err := json.Unmarshal(<-bodyChan, &prf); err != nil {
							log.Println(err)
						}

						// persist
						wg.Add(1)
						go func() {
							defer wg.Done()

							f.WriteString(fmt.Sprintf("%v\n", stringify(prf)))
						}()
					}
				}
			} else {
				// in federal level: other zones uses codes between "3" and "10" (skipping "8")
				for i := 3; i <= 10; i++ {
					if i == 8 {
						continue
					}

					endpt := fmt.Sprintf("%v/candidatura/listar/%v/%v/%v/%v/candidatos", TSE, year, zone, id, i)
					urlChan := sendURL(listCandidatosID(endpt), year, zone, id)

					for url := range urlChan {
						respChan := get(url)
						bodyChan := readResponse(<-respChan)

						var prf PerfilCandidato
						if err := json.Unmarshal(<-bodyChan, &prf); err != nil {
							log.Println(err)
						}

						// persist
						wg.Add(1)
						go func() {
							defer wg.Done()

							f.WriteString(fmt.Sprintf("%v\n", stringify(prf)))
						}()
					}
				}
			}
		}
	} else {
		for _, zone := range Zones {
			if zone == "BR" {
				continue
			}

			for _, cod := range listCidadesID(fmt.Sprintf("%v/eleicao/buscar/%v/%v/municipios", TSE, zone, id)) {

				for i := 11; i <= 13; i++ {
					endpt := fmt.Sprintf("%v/candidatura/listar/%v/%v/%v/%v/candidatos", TSE, year, cod, id, i)
					urlChan := sendURL(listCandidatosID(endpt), year, cod, id)

					for url := range urlChan {
						respChan := get(url)
						bodyChan := readResponse(<-respChan)

						var prf PerfilCandidato
						if err := json.Unmarshal(<-bodyChan, &prf); err != nil {
							log.Println(err)
						}

						// persist
						wg.Add(1)
						go func() {
							defer wg.Done()

							f.WriteString(fmt.Sprintf("%v\n", stringify(prf)))
						}()
					}
				}

			}
		}
	}

	wg.Wait()

	// upload to Az stg
	client, err := azblob.NewClientFromConnectionString(os.Getenv("AzureStgConnStr"), nil)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = client.UploadFile(context.TODO(), "ingestion", f.Name(), f, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// delete file
	if err := os.Remove(f.Name()); err != nil {
		log.Fatalln(err)
	}
}

// Given the selected user option, return properly API endpoint vars.
func buildEndpoint(in uint8, opt [][]string) (isFederal bool, id string, year string) {
	idx := strings.LastIndex(opt[in-1][0], " ")
	year = opt[in-1][0][idx+1:]
	id = opt[in-1][1]

	// new elections must be added here too
	switch in {
	case 1, 3, 5, 7, 9:
		isFederal = false
	case 10:
		isFederal, year = false, "2020"
	case 2, 4, 6, 8, 11:
		isFederal = true
	}

	return
}

// Request HTTP GET the URL and receive the channel for reading the response.
func get(url string) <-chan *http.Response {
	log.Printf("[GET] %q\n", url)
	respChan := make(chan *http.Response)

	go func(url string) {
		defer close(respChan)

		resp, err := http.Get(url)
		respChan <- resp

		if err != nil {
			log.Println(err)
		}
	}(url)

	return respChan
}

// Reads Response object and receive the channel containing its content.
func readResponse(resp *http.Response) <-chan []byte {
	bodyChan := make(chan []byte)

	go func(resp *http.Response) {
		defer close(bodyChan)
		// body, err := io.ReadAll(resp.Body)

		// if err != nil {
		// 	log.Println(err)
		// }

		defer resp.Body.Close()

		var body []byte
		if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
			log.Fatalf(">>>>>> %q\n", err)
		}

		bodyChan <- body
	}(resp)

	return bodyChan
}

func stringify(prf PerfilCandidato) string {
	str, err := json.Marshal(prf)

	if err != nil {
		log.Println(err)
	}

	return string(str)
}

func listCandidatosID(url string) []int64 {
	respChan := get(url)
	bodyChan := readResponse(<-respChan)

	var cand Candidaturas
	if err := json.Unmarshal(<-bodyChan, &cand); err != nil {
		log.Println(err)
	}

	var s []int64
	for _, c := range cand.Candidatos {
		s = append(s, c.ID)
	}

	return s
}

func listCidadesID(url string) []string {
	respChan := get(url)
	bodyChan := readResponse(<-respChan)

	var cid Cidades
	if err := json.Unmarshal(<-bodyChan, &cid); err != nil {
		log.Println(err)
	}

	var s []string
	for _, c := range cid.Municipios {
		s = append(s, c.Codigo)
	}

	return s
}

func sendURL(list []int64, year, zone, id string) <-chan string {
	urlChan := make(chan string)

	go func(y, z, i string) {
		defer close(urlChan)

		for _, v := range list {
			urlChan <- fmt.Sprintf("%v/candidatura/buscar/%v/%v/%v/candidato/%v", TSE, y, z, i, v)
		}
	}(year, zone, id)

	return urlChan
}
