# Envio de Mensagens para o Slack em Go

Este projeto é um simples cliente Go para enviar mensagens para um canal do Slack utilizando a API `chat.postMessage`.

## Pré-requisitos

Antes de rodar o código, você precisa:
- Criar um bot no Slack e obter um **Bot Token**.
- Conceder a permissão `chat:write` ao bot.
- Definir a variável de ambiente `SLACK_AUTHORIZER_KEY` com o token do bot.

## Como Executar

1. Clone o repositório:
   ```sh
   git clone https://github.com/seu-usuario/slack-bot-go.git
   cd slack-bot-go
   ```
2. Configure a variável de ambiente:
   ```sh
   export SLACK_AUTHORIZER_KEY="seu-bot-token-aqui"
   ```
3. Compile e execute o código:
   ```sh
   go run main.go
   ```
## 📖 Referências
- [Slack API - chat.postMessage](https://api.slack.com/methods/chat.postMessage)
- [Como criar um bot no Slack](https://api.slack.com/apps)
