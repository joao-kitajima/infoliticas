package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

const containerName string = "ingestion"

// JSON Struct
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

type MidiasCandidato struct {
	ID       int      `json:"id"`
	NomeUrna string   `json:"nomeUrna"`
	Midias   []string `json:"midia"`
}

func main() {
	log.Println("Lendo variáveis de ambiente ...")
	log.Printf("Endereço do arquivo de origem: %q.", os.Getenv("blobURL"))

	if !strings.HasSuffix(os.Getenv("blobURL"), ".jsonl") {
		log.Fatalln("Endereço não é um arquivo JSON Lines (.jsonl)!")
	}

	log.Printf("Nome da mídia social a ser pesquisada: %q.", os.Getenv("midia"))
	log.Printf("Padrão de busca nos endereços de mídias sociais: %q.", os.Getenv("pattern"))

	// downloading azblob
	blob := createTmp("filtro_midias_sociais_*.jsonl")
	defer closeRemoveTmp(blob)

	client, err := azblob.NewClientWithNoCredential(os.Getenv("blobURL"), nil)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Transferindo conteúdo para arquivo local ...")
	_, err = client.DownloadFile(context.TODO(), "", "", blob, nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Transferência concluída.")

	// creating tmp file
	log.Println("Criando arquivo temporário para registro das candidaturas com contas associadas ...")
	tmp := createTmp(fmt.Sprintf("%v.jsonl", os.Getenv("midia")))
	defer closeRemoveTmp(tmp)
	log.Println("Arquivo temporário criado com sucesso.")

	// reading file
	log.Println("Iterando sobre candidaturas ...")
	reader := bufio.NewReader(blob)
	var line []byte
	var wg sync.WaitGroup

	for {
		l, pfx, err := reader.ReadLine()

		// err and EOF check
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
		}

		if pfx {
			line = append(line, l...)
		} else {
			if len(line) == 0 {
				line = l
			} else {
				line = append(line, l...)
			}

			var prf PerfilCandidato
			if err := json.Unmarshal(line, &prf); err != nil {
				log.Println(err)
				continue
			}

			// collecting midias
			var t MidiasCandidato
			t.ID, t.NomeUrna = prf.ID, prf.NomeUrna

			for _, site := range prf.Sites {
				s := fmt.Sprint(site)

				if strings.Contains(s, os.Getenv("pattern")) {
					log.Printf("Mídia social encontrada! (%v)\n", s)
					// removing query params
					idx := strings.LastIndex(s, "?")
					if idx != -1 {
						log.Printf("Removendo parâmetros desnecessários da URL. (%v)\n", s[idx:])
						s = s[:idx]
					}

					t.Midias = append(t.Midias, s)
				}
			}

			if len(t.Midias) > 0 {
				data, err := json.Marshal(&t)
				if err != nil {
					log.Println(err)
				}

				// write to file
				wg.Add(1)
				go func() {
					defer wg.Done()

					_, err = tmp.WriteString(string(data) + "\n")
					if err != nil {
						log.Println(err)
					}
				}()
			}

			line = []byte("")
		}
	}

	wg.Wait()

	// upload blob
	log.Println("Criando conexão com Azure Storage ...")
	idx := strings.LastIndex(os.Getenv("blobURL"), "/")
	blobName := os.Getenv("blobURL")[idx+1:]
	blobName = fmt.Sprintf("tse/eleicoes/midias_sociais/%v/%v", os.Getenv("midia"), blobName)
	log.Printf("Caminho do Azure Blob configurado como: %q.\n", blobName)

	client, err = azblob.NewClientFromConnectionString(os.Getenv("AzureStgConnStr"), nil)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Transferindo arquivos para nuvem ...")
	_, err = client.UploadFile(context.TODO(), containerName, blobName, tmp, nil)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Transferência realizada. Encerrando execução do programa.")
}

func createTmp(pattern string) *os.File {
	f, err := os.CreateTemp("", pattern)

	if err != nil {
		log.Println(err)
	}

	return f
}

func closeRemoveTmp(f *os.File) {
	if err := f.Close(); err != nil {
		log.Println(err)
	}

	if err := os.Remove(f.Name()); err != nil {
		log.Println(err)
	}
}
