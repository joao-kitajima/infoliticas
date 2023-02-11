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
	var input uint8

	if _, err := fmt.Scanln(&input); err != nil {
		log.Fatalln(err)
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

	// zones
	// Zones := [28]string{"BR", "AC", "AM", "AP", "PA", "RO", "RR", "TO", "AL", "BA", "CE", "MA", "PB", "PE", "PI", "RN", "SE", "DF", "GO", "MS", "MT", "ES", "MG", "RJ", "SP", "PR", "RS", "SC"}
	// fmt.Println(Zones)

	if isFederal {
		// federal
		fmt.Println()
	} else {
		fmt.Println()
		// municipal
	}

}
