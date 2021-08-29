







> status:	🚧  Challenge-Bravo 🚀 em construção..  🚧

API, que responde JSON, para conversão monetária. Ela tem uma moeda de lastro (USD) e faz conversões entre diferentes moedas verídicas e fictícias.



A API, originalmente, converte entre as seguintes moedas:

-   USD
-   BRL
-   EUR
-   BTC
-   ETH

Ex: USD para BRL, USD para BTC, ETH para BRL, etc...

Para as moedas acima, a cotação do USD é obtida via consumo das seguintes APIs externas: moeda via [HGBRASIL](https://hgbrasil.com/status/finance) e de criptomoeda via [Coinbase](https://developers.coinbase.com/api/v2#get-buy-price). Para isso deverá ser usada a rota ``` /currency/currentQuote ```

No cadastro de nova moeda, rota ``` /currency ```, deverá ser informado a cotação do dólar. A atualização do valor do dólar para a nova moeda, será realizada pelo usuário através do campo ```valueInUSD``` via rota ``` /currency/:id ```

## Indice

* <p><a href="#pré-requisitos">Pré Requisitos</a> </p>
* <p><a href="#iniciando-projeto">Iniciando Projeto</a></p>
* <p><a href="#variáveis-de--ambiente">Variáveis de Ambiente</a></p>
* <p><a href="#rotas">Rotas</a></p>
* <p><a href="#executando-os-testes">Executando os testes</a></p>
* <p><a href="#relatório-de-cobertura-de-testes">Relatório de cobertura de Testes</a></p>
* <p><a href="#documentação">Documentação</a></p>
* <p><a href="#autor">Autor</a></p>




## Pré Requisitos

Antes de começar, você precisará ter as seguintes ferramentas instaladas em sua máquina:
* [Git](https://git-scm.com)
* [Golang](https://golang.org/)
* [Redis](https://redis.io/)


Além disso, é bom ter um editor para trabalhar com o código como: [VSCode](https://code.visualstudio.com/)



## Iniciando Projeto 

### Local

```bash
# Clone este repositório
$ git clone https://github.com/WallaceMachado/challenge-bravo.git

# Acesse a pasta do projeto no terminal / cmd
$ cd challenge-bravo

# Instale as dependências
$ go run main.go

# Rode o projeto
$ go run main.go

# Server is running
```


## Variáveis de Ambiente

Após clonar o repositório, renomeie o ``` .env.example ``` no diretório raiz para ``` .env ``` e atualize com suas configurações.
Apesar de não ser uma boa prática e gerar vunerabilidade na aplicação, considerando que a aplicação não irá para produção, rodará somente local e para fins didáticos,
foram inseridas explicitamente as variáveis de ambiente na tabela abaixo

| Chave  |  Descrição  | Predefinição  |
| :---: | :---: | :---: | 
|  API_PORT | Número da porta em que o aplicativo será executado. | 5000  |
|  STRING_CONNECTION_MONGO_DB |  String de conexão remota com o mongo |  mongodb+srv://admin:root@cluster0.pamgw.mongodb.net |
|  NAME_MONGO_DB |  Nome do banco de dados criado no mongo  |  chBravoDb |
|  ADRRESS_REDIS|  Aonde o Redis está sendo executado (host:port) |  localhost:6379  |
|  PASSWORD_REDIS |  Senha do Redis |    |
|  KEY_API_HGBRASIL|  Chave de acesso a api HGBRASIL |  b9524aa8  |



## Rotas

| Rotas  |  HTTP Method  | Params  |  Desccrição  |
| :---: | :---: | :---: | :---: |
|  /currency |  POST |  Body: ``` name ```, ``` code ``` e ``` valueInUSD ``` |  Crie uma nova moeda |
|  /currency |  GET |  -  | Recupere uma lista com todas as moedas |
|  /currency/conversion |  GET |  Query: ```from ``` (moeda de origem), ``` to ``` (moeda de conversão), ``` amount ``` (valor a ser convertido)  |  Consulte uma conversão monetária |
|  /currency/currentQuote |  GET | -  |  Atualize a cotação do USD das moedas originais do sistema (USD, BRL, EUR, BTC, ETH)  |
|  /currency/:id |  PUT |  Body: ``` name ```, ``` code ``` e ``` valueInUSD ```  |  Edite uma moeda |
|  /currency/:id |  DELETE |  -  |  Exclua uma moeda |



## Autor


Feito com ❤️ por [Wallace Machado](https://github.com/WallaceMachado) 🚀🏽 Entre em contato!

[<img src="https://img.shields.io/badge/linkedin-%230077B5.svg?&style=for-the-badge&logo=linkedin&logoColor=white" />](https://www.linkedin.com/in/wallace-machado-b2054246/)
