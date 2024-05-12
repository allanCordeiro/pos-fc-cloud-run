# FC POS Go Expert
## Desafio Open Telemetry

### Objetivo: 

Desenvolver um sistema em Go que receba um CEP, identifica a cidade e retorna o clima atual (temperatura em graus celsius, fahrenheit e kelvin).  Esse sistema deverá implementar OTEL(Open Telemetry) e Zipkin.

### Execução local
Subir o serviço via docker-compose:
`docker-compose up -d`

Serviço do Zipkin está configurado em `http://localhost:9411/`

Existe um arquivo, `resquest_sample.http` na raiz do repositório, com uma requisição padrão ao serviço A, para auxiliar na avaliação.