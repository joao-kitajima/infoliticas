package main

const TSE string = "https://divulgacandcontas.tse.jus.br/divulga/rest/v1"

// append new elections at the end
var Elections = [][]string{
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

var Zones = [28]string{
	"BR",                                     // Federal
	"AC", "AM", "AP", "PA", "RO", "RR", "TO", // North
	"AL", "BA", "CE", "MA", "PB", "PE", "PI", "RN", "SE", // Northeast
	"DF", "GO", "MS", "MT", // Central-west
	"ES", "MG", "RJ", "SP", // Southeast
	"PR", "RS", "SC", // South
}

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
