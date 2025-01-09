# Link-Up Project Overview

## 📋 Descrição Geral

O **Link-Up Project** é uma solução completa para a gestão de serviços de internet, TV e celular. É composta por:

1. Uma **API Backend** desenvolvida em **Go**, responsável pela lógica de negócios e integração entre serviços.
2. **Três serviços independentes**, que incluem funcionalidades de verificação de disponibilidade, agendamento de instalação e envio de e-mails.
3. Um **aplicativo mobile** desenvolvido em **React Native**, que consome a API e oferece uma interface amigável para os usuários.

---

## 🛠️ Tecnologias Utilizadas

- **Backend:** Go (1.23.4)
  - MongoDB
  - Apache Kafka (Sarama v1.45.0)
- **Mobile App:** React Native
  - Expo CLI
  - Axios para requisições HTTP
- **Infraestrutura:** Docker Compose para serviços locais

---

## 📂 Estrutura Geral do Projeto

```plaintext
link-up-project/
├── application/            # Aplicação principal (API Backend)
├── services/               # Serviços independentes
│   ├── installation/       # Serviço de agendamento de instalação
│   ├── email/              # Serviço de envio de e-mails
│   └── availability/       # Serviço de verificação de disponibilidade
├── mobile-app/             # Aplicativo React Native
│   ├── App.js              # Arquivo principal
│   ├── components/         # Componentes reutilizáveis
│   ├── screens/            # Telas do aplicativo
│   ├── services/           # Configuração do Axios e chamadas para a API
│   └── assets/             # Imagens e recursos estáticos
├── kafka/                  # Configuração do Kafka
├── go.mod                  # Dependências do Go
├── README.md               # Este arquivo
└── docker-compose.yml      # Configuração para Docker Compose
```

---

## 🧑‍💻 Como Rodar o Projeto

### Pré-requisitos

- Node.js (para o aplicativo React Native)
- Expo CLI
- Go (1.23.4)
- MongoDB
- Apache Kafka
- Docker (opcional)

### Passos para Rodar o Backend e os Serviços

1. **Clone o repositório:**

   ```bash
   git clone https://github.com/seu-usuario/link-up-project.git
   cd link-up-project
   ```

2. **Rode os serviços do MongoDB e Kafka:**

   ```bash
   docker-compose up -d
   ```

3. **Rode cada serviço separadamente:**
   - **Aplicação Principal:**
     ```bash
     cd application
     go run main.go
     ```
   - **Serviço de Agendamento de Instalação:**
     ```bash
     cd services/installation
     go run main.go
     ```
   - **Serviço de Envio de E-mails:**
     ```bash
     cd services/email
     go run main.go
     ```
   - **Serviço de Disponibilidade:**
     ```bash
     cd services/availability
     go run main.go
     ```

### Passos para Rodar o Aplicativo Mobile

1. **Instale o Expo CLI:**

   ```bash
   npm install -g expo-cli
   ```

2. **Navegue até o diretório do aplicativo:**

   ```bash
   cd mobile-app
   ```

3. **Instale as dependências:**

   ```bash
   npm install
   ```

4. **Inicie o aplicativo:**

   ```bash
   expo start
   ```

5. **Teste no emulador ou dispositivo físico:**
   - Escaneie o QR code gerado pelo Expo usando o aplicativo Expo Go.

---

## 🔥 Funcionalidades do Aplicativo Mobile

1. **Tela de Login:**

   - Autenticação do usuário (a ser implementada).

2. **Verificar Disponibilidade:**

   - Formulário para verificar a disponibilidade de serviços em um endereço.

3. **Gerenciar Propostas:**
   - Criar e visualizar propostas enviadas para a API.

### Exemplos de Chamadas para a API (Usando Axios)

#### Configuração do Axios

Arquivo: `mobile-app/services/api.js`

```javascript
import axios from "axios";

const api = axios.create({
  baseURL: "http://localhost:8080",
});

export default api;
```

#### Exemplo de Verificação de Disponibilidade

Arquivo: `mobile-app/services/availability.js`

```javascript
import api from "./api";

export const checkAvailability = async (address) => {
  try {
    const response = await api.post("/availability", address);
    return response.data;
  } catch (error) {
    console.error("Error checking availability:", error);
    throw error;
  }
};
```

---

## 👥 Contribuidores

- [Julio Cesar](https://github.com/julio16j)

---
