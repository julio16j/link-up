# Link-Up API

## 📋 Descrição

A **Link-Up API** é uma aplicação backend desenvolvida em **Go** que oferece suporte para verificação de disponibilidade de serviços e criação de propostas para planos de internet, TV e celular. A API utiliza MongoDB para armazenar dados e Kafka para comunicação entre serviços.

---

## 🛠️ Tecnologias Utilizadas

- **Linguagem:** Go (versão 1.23.4)
- **Banco de Dados:** MongoDB
- **Mensageria:** Apache Kafka (com a biblioteca Sarama v1.45.0)
- **Frameworks e Bibliotecas:**
  - `github.com/IBM/sarama` (Kafka)
  - `go.mongodb.org/mongo-driver` (MongoDB)

---

## 🚀 Funcionalidades

1. **Verificar Disponibilidade (`/availability`):**

   - Endpoint para verificar se há disponibilidade de serviço em um determinado endereço.

2. **Criar Proposta (`/proposal`):**
   - Endpoint para iniciar o processo de criação de uma proposta e enviar a mensagem para o Kafka.

---

## 📂 Estrutura do Projeto

```plaintext
link-up-api/
├── config/                # Configuração do MongoDB
│   └── database.go
├── controllers/           # Controladores dos endpoints
│   ├── availability_controller.go
│   └── proposal_controller.go
├── models/                # Modelos de dados
│   ├── address.go
│   └── proposal.go
├── repositories/          # Repositórios para acesso ao banco e Kafka
│   ├── availability_repository.go
│   └── proposal_repository.go
├── routes/                # Registro de rotas
│   └── routes.go
├── services/              # Camada de regras de negócio
│   ├── availability_service.go
│   └── proposal_service.go
├── go.mod                 # Arquivo de dependências do Go
├── go.sum                 # Hash das dependências
└── main.go                # Ponto de entrada da aplicação
```

---

## 🧑‍💻 Como Rodar o Projeto

### Pré-requisitos

- Go (1.23.4)
- MongoDB
- Apache Kafka

### Passos

1. **Clone o repositório:**

   ```bash
   git clone https://github.com/seu-usuario/link-up-api.git
   cd link-up-api
   ```

2. **Instale as dependências:**

   ```bash
   go mod tidy
   ```

3. **Configure o MongoDB:**

   - Certifique-se de que o MongoDB está rodando em `localhost:27017`.
   - Use as credenciais padrão:
     - **Usuário:** `root`
     - **Senha:** `root`

4. **Configure o Kafka:**

   - Certifique-se de que o Kafka está rodando em `localhost:9092`.

5. **Rode a aplicação:**
   ```bash
   go run main.go
   ```

---

## 🔥 Testando a API

### Endpoint: `/availability`

- **Método:** `POST`
- **URL:** `http://localhost:8080/availability`
- **Body:**
  ```json
  {
    "street": "Rua Desembargador Azevedo",
    "city": "Campina Grande",
    "state": "PB",
    "zipcode": "58401-160"
  }
  ```
- **Resposta:**
  ```json
  {
    "available": true
  }
  ```

### Endpoint: `/proposal`

- **Método:** `POST`
- **URL:** `http://localhost:8080/proposal`
- **Body:**
  ```json
  {
    "customer_name": "John Doe",
    "address": {
      "street": "Rua Desembargador Azevedo",
      "city": "Campina Grande",
      "state": "PB",
      "zipcode": "58401-160"
    },
    "plan": "Internet + TV",
    "email": "johndoe@example.com"
  }
  ```
- **Resposta:**
  ```json
  {
    "status": "Proposal initiated"
  }
  ```

---

## 👥 Contribuidores

- [Julio Cesar](https://github.com/julio16j)

---
