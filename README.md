# Sistema de Consulta de Clima por CEP com Otel

Este projeto consiste em dois serviços desenvolvidos em Go: Serviço A e Serviço B.

Serviço A: Recebe um CEP via POST, valida e encaminha ao Serviço B.
Serviço B: Consulta a localização e o clima do CEP fornecido e retorna as temperaturas em Celsius, Fahrenheit e Kelvin.
Tecnologias Utilizadas
Go
Docker/Docker Compose
OpenTelemetry (OTEL)
Zipkin

# Executando o Projeto

### Para executar o projeto, siga os passos abaixo:
Clone o repositório:

execute o docker-compose:
docker compose up --build

#### Acesse o serviço A em http://localhost:8080/cep

# Observabilidade com OTEL e Zipkin
#### Porta http://localhost:9411


# Exemplo de Requisição para testar a resposta

curl -X POST http://localhost:8080/cep -d '{"cep": "04127040"}' -H "Content-Type: application/json"

### Postman ou insomnia
POST http://localhost:8080/cep

Body:
```json
{
    "cep": "04127040"
}
```
