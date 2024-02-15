# Full Cycle - Go Export - Desafio Rate Limiter

Este projeto faz parte de um desafio do curso de pós gradução em Go - Go Expert - da Full Cycle.

## Introdução

Este projeto implementa um *Rate Limiter* escrito na linguagem de programação Go.

### Configurações

Na raiz deste projeto, encontra-se um arquivo denominado `app.env`, através do qual é possível definir as configurações do **Rate Limiter**.

* `DB_PROTOCOL`: Nome to protocolo para acesso ao banco de dados.
* `DB_HOST`: Endereço/IP para conexão com o banco de dados.
* `DB_PORT`: Porta para conexão com o banco de dados.
* `DB_USER`: Nome to usuário para acesso ao banco de dados.
* `DB_PASSWORD`: Senha para acesso ao banco de dados.
* `DB_DATABASE`: Nome do banco de dados.
* `RATE_LIMIT`: Número máximo de requisições por IP por segundo.
* `LIMIT_BY_TOKEN`: Quando o valor desta configuração é `true`, a limitação de requisições será baseada no token enviado pelo cliente no cabeçalho da requisição. Caso contrário, a limitação será por IP.
* `EXPIRATION_TIME`: Tempo de expiração em segundos quando um IP ou token é bloqueado.

Para que mudanças no arquivo de configuração sejam refletidas no servidor é necessário efetuar um novo *build* do Docker compose.

### Tokens disponíveis

Os tokens válidos para testes de requisições estão disponíveis em `./assets/tokens.json`. Eis os tokens disponíveis e seus respectivos limites de requisições por segundo:
* `p7eWgd0PvJcqB3ea45pw3k5thpWaqpI12RGYU3MiP91Kgao5MCXtlFtL2rwISxYL`: 100
* `65aYJmkHf8QC52s10HYjVV5xtfwaKf3qC2J79cKaFZarStmbT6Mueic195OXLXVy`: 200
* `DUc0K3ojDA0kQgCVl2SEfT5evimjAEojGs5QOxMfA3JAgdrF6I5l8hHFXDDtpxqv`: 300
* `sEUMDAjUhxGzTcXwerMiCwdSfe0vp24q9ISuZrYra0Hj65gtGpGeA5zZt4bjG0gz`: 400
* `IB9BB71TsBk8QevTskkFusRrZwfUBjmqACA7vvsc1TCpS5FMAmZZTZx1R1OhA12y`: 500

Caso necessário, é possível adicionar novos tokens no arquivo `./assets/tokens.json`.

### Como executar o servidor

Para executar o servidor, através do terminal acesse a raiz deste projeto e execute:

```sh
docker compose --env-file app.env up -d --build
```

### Como testar as requisições

As requisições devem ser enviadas para a porta `8080`, isto é, a porta onde está rodando o servidor. O token de acesso, quando necessário, deve ser enviado como valor do cabeçalho `API_KEY`.

Para testar as requisições, é possível utilizar ferramentas visuais como o [Postman](https://www.postman.com/) ou [Insomnia](https://insomnia.rest/). Também é possível usar ferramentas do tipo CLI, tais como [cURL](https://curl.se/) ou [httpie](https://httpie.io/).

Abaixo, dois exemplos de requisição usando o **httpie**. A segunda requisição inclui um token no cabeçalho:
```sh
http GET http://localhost:8080
http GET http://localhost:8080 API_KEY:p7eWgd0PvJcqB3ea45pw3k5thpWaqpI12RGYU3MiP91Kgao5MCXtlFtL2rwISxYL
```

Se tudo correr bem, o servidor deverá responder com a mensagem "Welcome" e código 200. Em caso de bloqueio do IP ou do token, o servidor retornará a mensagem "you have reached the maximum number of requests or actions allowed within a certain time frame" e código 429.

Para testar uma grande quantidade de requisições com o objetivo de simular o bloqueio de um IP ou token, sugere-se o uso da ferramenta [Apache ab](https://httpd.apache.org/docs/2.4/programs/ab.html).

Segue abaixo exemplos de uso da ferramenta **Apache ab**:
```sh
ab -n 20 http://localhost:8080/
ab -n 101 -H "API_KEY: p7eWgd0PvJcqB3ea45pw3k5thpWaqpI12RGYU3MiP91Kgao5MCXtlFtL2rwISxYL" http://localhost:8080/
```

### Executando testes de unidade e integração

Para executar os testes de unidade e integração do projeto, através do terminal acesse o diretório raiz deste repositório. Usando a ferramenta **make**, execute:

```sh
make test
```

Se a ferramenta **make** não estiver disponível, execute:

```sh
go test -v ./...
```

Para gerar um relatório de cobertura dos testes, execute:

```sh
make test-cov
```

Se a ferramenta **make** não estiver disponível, execute:
```sh
go test -v ./... -coverprofile=c.out
go tool cover -html=c.out
```