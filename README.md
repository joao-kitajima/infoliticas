# Candidaturas TSE

Programa desenvolvido para extrair os dados das candidaturas registradas no TSE (Tribunal Superior Eleitoral). A fonte de origem dos dados pode ser encontrada [aqui](https://divulgacandcontas.tse.jus.br/divulga/).

## Exemplo de requisição de um candidato

https://divulgacandcontas.tse.jus.br/divulga/rest/v1/candidatura/buscar/2020/38490/2030402020/candidato/50000674223

Endpoint segue o seguinte padrão: `<host>/divulga/rest/v1/candidatura/buscar/<anoEleicao>/<codRegiao>/<codEleicao>/candidato/<codCandidato>`

## Variáveis de ambiente necessárias

Este programa requer as seguintes variáveis de ambiente para ser executado:

- **AZURE_STG_CONN_STR**: connection string do Azure Storage onde serão armazenados os dados de saída.
- **AZURE_STG_WD**: caminho de pastas dentro do Azure Storage (exemplo: container/grandparent/parent/myblob.jsonl).
- **ELECTION_NUM**: número da eleição que se deseja realizar a extração das candituras.

### Tabela de Eleições

1. Eleições Municipais 2004
2. Eleição Geral Federal 2006
3. Eleições Municipais 2008
4. Eleição Geral Federal 2010
5. Eleições Municipais 2012
6. Eleição Geral Federal 2014
7. Eleições Municipais 2016
8. Eleição Geral Federal 2018
9. Eleições Municipais 2020
10. Eleições Municipais 2020 - AP
11. Eleição Geral Federal 2022
