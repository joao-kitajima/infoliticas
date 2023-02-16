package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

const containerName string = "ingestion"

type resultCount struct {
	name string
	f    *os.File
	blob *azblob.Client
}

func (r *resultCount) createTmpFile() {
	name := fmt.Sprintf("%v.jsonl", r.name)

	var err error
	r.f, err = os.Create(name)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Arquivo temporário local criado: %q.", name)
}

func (r *resultCount) removeTmpFile() { os.Remove(r.f.Name()) }

func (r *resultCount) writeObj(obj string, wg *sync.WaitGroup) {
	defer wg.Done()
	r.f.WriteString(obj)
}

func (r *resultCount) createBlobClient() {
	var err error
	r.blob, err = azblob.NewClientFromConnectionString(os.Getenv("AzureStgConnStr"), nil)
	if err != nil {
		log.Fatalln(err)
	}
}

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

func main() {
	// asserting var env src url input
	srcBlob := os.Getenv("blobURL")

	if !strings.HasSuffix(srcBlob, ".jsonl") {
		log.Fatalln("Endereço URL do Azure Blob não é um arquivo JSON Lines!")
	}

	// creating tmp file
	log.Println("Criando arquivo temporário em pasta local ...")
	f, err := os.Create("tmp.jsonl")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	log.Println("Arquivo criado.")

	log.Printf("Baixando arquivo da URL recebida: %q.\n", os.Getenv("blobURL"))
	// downloading az blob to local tmp file
	client, err := azblob.NewClientWithNoCredential(os.Getenv("blobURL"), nil)
	if err != nil {
		log.Fatalln(err)
	}

	blobClientOpts := azblob.DownloadFileOptions{
		Progress: func(bytesTransferred int64) {
			log.Printf("Recebendo %d bytes ...", bytesTransferred)
		},
	}

	_, err = client.DownloadFile(context.TODO(), "", "", f, &blobClientOpts)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Transferência concluída com sucesso.")

	// removing base URL
	idx := strings.LastIndex(srcBlob, "/")
	srcBlob = srcBlob[idx+1:]

	// creating result counts catgs
	// "Eleito"
	var eleitos resultCount
	eleitos.name = "eleitos"
	eleitos.createTmpFile()
	defer eleitos.removeTmpFile()
	defer eleitos.f.Close()
	eleitos.createBlobClient()

	// "Não Eleito"
	var naoEleitos resultCount
	naoEleitos.name = "NAOeleitos"
	naoEleitos.createTmpFile()
	defer naoEleitos.removeTmpFile()
	defer naoEleitos.f.Close()
	naoEleitos.createBlobClient()

	// "Suplente"
	var suplentes resultCount
	suplentes.name = "suplentes"
	suplentes.createTmpFile()
	defer suplentes.removeTmpFile()
	defer suplentes.f.Close()
	suplentes.createBlobClient()

	// "Concorrendo"
	var concorrendo resultCount
	concorrendo.name = "concorrendo"
	concorrendo.createTmpFile()
	defer concorrendo.removeTmpFile()
	defer concorrendo.f.Close()
	concorrendo.createBlobClient()

	// Demais Opções
	var demais resultCount
	demais.name = "demais"
	demais.createTmpFile()
	defer demais.removeTmpFile()
	defer demais.f.Close()
	demais.createBlobClient()

	log.Println("Iniciando leitura sobre os registros encontrados ...")

	var wg sync.WaitGroup
	scanner := bufio.NewScanner(f)
	count := 0

	for scanner.Scan() {
		count++
		obj := scanner.Bytes()

		var prf PerfilCandidato
		if err := json.Unmarshal(obj, &prf); err != nil {
			log.Println(err)
		}

		if prf.ID == 0 {
			continue
		}

		switch prf.DescricaoTotalizacao {
		case "Eleito", "Eleito por QP", "Eleito por média":
			wg.Add(1)
			go eleitos.writeObj(string(obj)+"\n", &wg)

		case "Não eleito":
			wg.Add(1)
			go naoEleitos.writeObj(string(obj)+"\n", &wg)

		case "Suplente":
			wg.Add(1)
			go suplentes.writeObj(string(obj)+"\n", &wg)

		case "Concorrendo":
			wg.Add(1)
			go concorrendo.writeObj(string(obj)+"\n", &wg)

		default:
			wg.Add(1)
			go demais.writeObj(string(obj)+"\n", &wg)

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Erro de leitura! Problema encontrado na linha %d.\n", count)
	} else {
		log.Println("Fim do arquivo. Leitura encerrada.")
	}

	blobs := []resultCount{eleitos, naoEleitos, suplentes, concorrendo, demais}

	for _, b := range blobs {
		wg.Add(1)

		go func(r resultCount) {
			defer wg.Done()
			blobName := fmt.Sprintf("tse/eleicoes/totalizacao/%v_%v", r.name, srcBlob)

			_, err := r.blob.UploadFile(context.TODO(), containerName, blobName, r.f, nil)
			if err != nil {
				log.Fatalln(err)
			}
			log.Printf("Transferindo arquivo %q para o Azure Storage.", blobName)
		}(b)
	}

	// rm tmp file
	os.Remove(f.Name())

	wg.Wait()
	log.Println("Transferência completa. Encerrando programa.")
}
