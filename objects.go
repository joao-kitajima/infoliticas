package main

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

const TSE string = "https://divulgacandcontas.tse.jus.br/divulga/rest/v1/"
