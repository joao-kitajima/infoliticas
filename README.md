# Candidaturas TSE - Filtro de Mídias Sociais

Este programa requer as seguintes variáveis de ambiente para ser executado:
* **AzureStgConnStr**: connection string do Azure Storage onde serão armazenados os dados de saída.
* **blobURL**: Endereço do Azure Blob contendo as candidaturas de origem.
* **midia**: Nome da mídia social a ser pesquisada. (**Importante**: também será o nome da pasta onde será gravado o arquivo).
* **pattern**: Utilizado para filtrar as URLs.

#### Exemplo de Configuração de variáveis
* **blobURL** = "https://`<account>`.blob.core.windows.net/ingestion/tse/eleicoes/totalizacao/eleitos.jsonl"
* **midia** = "twitter"
* **pattern**: "twitter.com"
