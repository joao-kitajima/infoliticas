package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Saudações! Bem-vindo(a) ao Infolíticas!\nEste programa realiza a extração de informações das mídias sociais das candidaturas divulgadas pelo TSE (Tribunal Superior Eleitoral) sobre as eleições brasileiras.\n\nSobre qual eleição você deseja mais informações?\n(Digite o valor na linha de comando)")

	Elections := [][]string{
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

	for i, v := range Elections {
		i++
		fmt.Printf("%v. %v\n", i, v[0])
	}

	// scanning user input
	var input uint8 = 0

	for input <= 0 || int(input) > len(Elections) {
		if _, err := fmt.Scanln(&input); err != nil {
			log.Fatalln(err)
		}

		if input <= 0 || int(input) > len(Elections) {
			fmt.Println("Opção inválida!")
		} else {
			break
		}
	}

	// building API endpoint
	var isFederal bool
	APITSE := "https://divulgacandcontas.tse.jus.br/divulga/rest/v1"

	switch input {
	case 1:
		fmt.Printf(`Você selecionou a opção "%v. %v".`+"\n", input, Elections[input-1][0])
		isFederal = false
		APITSE += "/"

	case 2:
		fmt.Printf(`Você selecionou a opção "%v. %v".`+"\n", input, Elections[input-1][0])
		isFederal = true
	}

	fmt.Println(isFederal)

	// OQ DESEJA EXTRAIR ESTARIA AQ
	//
	//
	//
	//

	// zones
	Zones := [28]string{"BR", "AC", "AM", "AP", "PA", "RO", "RR", "TO", "AL", "BA", "CE", "MA", "PB", "PE", "PI", "RN", "SE", "DF", "GO", "MS", "MT", "ES", "MG", "RJ", "SP", "PR", "RS", "SC"}
	fmt.Println(Zones)

	if isFederal {
		// federal
		fmt.Println()

		// 1 - DESNECESSARIO
		// GET em cada regiao
		// "https://divulgacandcontas.tse.jus.br/divulga/rest/v1/eleicao/listar/municipios/14417/BR/cargos" <- .../codELEICAO/REGIAO/cargos
		// usar o de baixo direto
		// if regiao == BR > codigos 1 e 2
		// if regiao == DEMAIS > codigos 3 ao 10

		// 2
		// GET em cada CARGO de cada regiao
		// "https://divulgacandcontas.tse.jus.br/divulga/rest/v1/candidatura/listar/2010/BR/14417/1/candidatos" <- .../listar/ANO/codREGIAO/codELEICAO/codCARGO/candidatos
		// pegar ID candidato para o passo 3

		// 3
		// pagina do candidato
		// /buscar/ANO/codREGIAO/codELEICAO/candidato/idCANDIDATO
		// "https://divulgacandcontas.tse.jus.br/divulga/rest/v1/candidatura/buscar/2006/BR/14423/candidato/61"

	} else {
		// municipal
		fmt.Println()

		// 1
		// ignorar regiao "BR"
		// [GET] ""https://divulgacandcontas.tse.jus.br/divulga/rest/v1/eleicao/buscar/AC/2030402020/municipios
		// retorna municipios daquele estado
		// pegar codigos dos municipios para passo 2

		// 2
		// iterar sobre os codigos 11, 12 e 13 (prefeito, vp e vereador) para cada municipio
		// "https://divulgacandcontas.tse.jus.br/divulga/rest/v1/candidatura/listar/2020/01120/2030402020/11/candidatos"
		// coletar IDs dos candidatos para o passo 3

		// 3
		// pagina do candidato
		// "https://divulgacandcontas.tse.jus.br/divulga/rest/v1/candidatura/buscar/2020/01120/2030402020/candidato/10000854328"

	}

}
