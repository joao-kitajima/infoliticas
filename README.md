# Separação dos resultados de totalização das candidaturas

Este programa requer as seguintes variáveis de ambiente para ser executado:
* **AzureStgConnStr**: connection string do Azure Storage onde serão armazenados os dados de saída.
* **blobURL**: URL do Azure Blob que concentra todos as candidaturas extraídas.

As candidaturas são dividas em até 5 categorias:
* **Eleito**
* **Não eleito**
* **Suplente**
* **Concorrendo**
* **Demais**: referente à demais opções. Serve para filtar candidaturas inválidas e totalizações como "2º turno".
