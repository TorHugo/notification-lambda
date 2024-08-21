# Notification-Lambda
O Notification Lambda é uma função AWS Lambda desenvolvida em Go, projetada para processar mensagens de um tópico Kafka. Cada mensagem recebida é utilizada para acionar um endpoint HTTP que envia notificações por e-mail ou SMS. Essa função é ideal para sistemas que exigem alta escalabilidade e processamento assíncrono de mensagens.

> 👨‍💻 A branch develop é dedicada ao desenvolvimento e testes locais do projeto. Nesta branch, você encontrará um **Kafka Consumer** que lê mensagens do tópico e aciona um endpoint HTTP. Essa configuração permite que você teste a lógica de consumo e envio de notificações em um ambiente local, sem a necessidade de uma função Lambda.
lambda

> 🚧 A branch lambda é focada na implementação e configuração da **Lambda Function** para ser implantada na **AWS**. Aqui, o código é ajustado para o ambiente AWS Lambda, incluindo a integração com o Kafka e o gerenciamento de mensagens. Esta branch é usada para preparar o código para execução em produção na **AWS**, garantindo que ele funcione conforme o esperado no ambiente **Lambda**.

# Architecture
Arquitetura
Este projeto adota a arquitetura hexagonal, também conhecida como Ports and Adapters. Essa arquitetura é projetada para isolar o núcleo da aplicação (domínio) das dependências externas, promovendo maior modularidade, testabilidade e facilidade de manutenção.

```tree 
notification-lambda/
│
├── cmd/
│   └── main.go
│
├── internal/
│   ├── core/
│   │   ├── domain/
│   │   │   └── notification.go 
│   │   │   └── parameter.go 
│   │   └── service/
│   │       └── notification_service.go
│   │
│   ├── ports/
│   │   ├── consumer.go
│   │   └── http_client.go
│   │
│   └── adapters/
│       ├── messaging/
│       │   └── kafka_consumer.go
│       └── http/
│           └── http_client.go
│
└── config/
    └── kafka_config.go
```

- **Core** (Domínio): Contém as regras de negócio e a lógica central da aplicação. O core é a parte mais importante do sistema, pois define o comportamento do domínio, sem se preocupar com detalhes de implementação externa.
- **Ports** (Interfaces): Define contratos que os adaptadores externos devem seguir para interagir com o núcleo da aplicação. Isso garante que o domínio permaneça independente das tecnologias específicas utilizadas pelos adaptadores. Exemplos de ports incluem consumer.go e http_client.go.
- **Adapters** (Implementações): Implementa as interfaces definidas pelos ports, conectando o núcleo da aplicação a serviços externos, como bancos de dados, filas de mensagens (Kafka), APIs externas, etc. Aqui, a lógica específica de integração com tecnologias externas é tratada, como em kafka_consumer.go e http_client.go.
