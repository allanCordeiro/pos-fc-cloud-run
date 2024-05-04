# FC POS Go Expert
## Desafio Google Cloud Run

### Objetivo: 

Desenvolver um sistema em Go que receba um CEP, identifica a cidade e retorna o clima atual (temperatura em graus celsius, fahrenheit e kelvin). Esse sistema deverá ser publicado no Google Cloud Run.

### URL do GCP Cloud Run:

https://fc-go-expert-cloud-run-63a2rdbceq-uc.a.run.app/weather/{zip_code}

Onde *zip_code* é o número do CEP que se deseja realizar a busca.

### Testes e2e:

estão contidos no diretório `test`
- cloud_run_weather_get.http: testes no Cloud Run
- local_weather_get.http: testes locais (necessário subir o serviço atravé do `docker-compose.yaml`)

### Testes unitários/integração
Usar a própria ferramenta de testes do Go
`go test ./...`

### Execução local
Subir o serviço via docker-compose:
`docker-compose up -d`
