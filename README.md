# Clima App by CEP

## Descrição

Este sistema, construído em Go, permite a consulta de um CEP, obtendo a localização da cidade correspondente e fornecendo as condições climáticas atuais, incluindo a temperatura em Celsius, Fahrenheit e Kelvin. Ele é implantado na plataforma Google Cloud Run para garantir escalabilidade e eficiência.

## Como Executar

### Requisitos

- Docker
- Docker Compose

### Passos para Executar Localmente

1. **Clone o repositório:**

   ```sh
   git clone https://github.com/jonasjesusamerico/goexpert-temperatura-cep-google-cloud-run.git;
   cd goexpert-temperatura-cep-google-cloud-run
   ```

2. **Gerar imagem**: Na raiz do projeto, execute: `sudo docker-compose build`

3. **Executar com Docker Compose**: `sudo docker-compose up`


4. **Acesse a aplicação:** Aplicação disponível: `http://localhost:8080`.

### Exemplos de Requisição

- **Para obter o clima de um CEP específico:** ``` curl -X GET http://localhost:8080/clima/88804050 ```

### Testando o Serviço no Google Cloud Run

- **Google cloud Run**  ` curl -X GET https://goexpert-temperatura-6sqqs2orqq-uc.a.run.app/clima/88804050 `

