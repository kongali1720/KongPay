# 🏛 KongPay Architecture

KongPay is built using **Clean Architecture**, **Dependency Injection**, and the **Repository Pattern**.

The goal is to keep business logic independent from infrastructure so that the application remains maintainable, testable, and scalable.

---

# 🎯 Design Principles

- API First
- Clean Architecture
- SOLID Principles
- Repository Pattern
- Dependency Injection
- Stateless Services
- Cloud Native
- Docker Ready
- Production Ready

---

# 🏗 High-Level Architecture

```mermaid
flowchart TD

A[Client]
--> B[Gin Router]

B --> C[HTTP Handlers]

C --> D[Business Services]

D --> E[Repository Layer]

E --> F[(PostgreSQL)]

D -. Future Cache .-> G[(Redis)]
```

---

# 📡 Request Flow

```mermaid
sequenceDiagram

participant Client
participant Router
participant Handler
participant Service
participant Repository
participant PostgreSQL

Client->>Router: HTTP Request
Router->>Handler: Route Request
Handler->>Service: Validate Input
Service->>Repository: Execute Business Logic
Repository->>PostgreSQL: SQL Query
PostgreSQL-->>Repository: Result
Repository-->>Service: Model
Service-->>Handler: Response
Handler-->>Client: JSON Response
```

---

# 🧱 Project Layers

```text
Client
│
▼
Gin Router
│
▼
HTTP Handler
│
▼
Service Layer
│
▼
Repository Layer
│
▼
PostgreSQL
```

---

# 📋 Layer Responsibilities

| Layer | Responsibility |
|--------|----------------|
| Client | Sends HTTP requests |
| Router | Maps endpoints to handlers |
| Handler | Request binding and response formatting |
| Service | Business logic and validation |
| Repository | Database access |
| PostgreSQL | Persistent storage |
| Redis | Future caching and coordination |

---

# 📂 Project Structure

```text
kongpay/
│
├── cmd/
│   └── kongpay/
│       └── main.go
│
├── internal/
│   ├── config/
│   ├── database/
│   ├── handlers/
│   ├── middleware/
│   ├── models/
│   ├── repositories/
│   ├── router/
│   ├── services/
│   └── utils/
│
├── migrations/
├── docker/
├── docs/
├── tests/
├── .github/
│
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

---

# 🔄 Dependency Injection

```mermaid
flowchart LR

Database

--> Repository

Repository

--> Service

Service

--> Handler

Handler

--> Router
```

The application wires dependencies in the following order:

1. Connect PostgreSQL.
2. Create repositories.
3. Create services.
4. Create handlers.
5. Register HTTP routes.
6. Start the Gin server.

---

# 🗄 Data Flow

```mermaid
flowchart LR

HTTP_Request

--> Handler

Handler

--> Service

Service

--> Repository

Repository

--> PostgreSQL

PostgreSQL

--> Repository

Repository

--> Service

Service

--> Handler

Handler

--> HTTP_Response
```

---

# 🔐 Authentication Flow (Planned)

```mermaid
flowchart LR

Client

--> JWT

JWT

--> Middleware

Middleware

--> Handler

Handler

--> Service

Service

--> Repository
```

---

# 🌐 Future Microservice Architecture

```mermaid
flowchart TD

Gateway

--> Wallet

Gateway

--> Merchant

Gateway

--> Payment

Gateway

--> Settlement

Gateway

--> Notification

Gateway

--> Authentication

Wallet --> PostgreSQL

Merchant --> PostgreSQL

Payment --> PostgreSQL

Settlement --> PostgreSQL
```

---

# 🛠 Current Components

| Component | Status |
|------------|:------:|
| Router | ✅ |
| Wallet Handler | ✅ |
| Wallet Service | ✅ |
| Wallet Repository | ✅ |
| PostgreSQL | ✅ |
| Redis | ✅ |
| Docker | ✅ |
| Health Endpoint | ✅ |

---

# 🚧 Planned Components

| Component | Status |
|------------|:------:|
| JWT Authentication | 🚧 |
| Swagger/OpenAPI | 🚧 |
| SQL Migration | 🚧 |
| Merchant Service | ⏳ |
| Payment Engine | ⏳ |
| Ledger | ⏳ |
| Settlement | ⏳ |
| QRIS Integration | ⏳ |

---

# 🚀 Scalability Vision

KongPay is designed to evolve from a modular monolith into a collection of independent services while preserving a consistent API-first approach.

Future enhancements include:

- Kubernetes Deployment
- API Gateway
- Horizontal Scaling
- Distributed Caching
- Observability
- Metrics & Monitoring
- Event-Driven Messaging
- High Availability

---

# ❤️ Architecture Philosophy

> **Keep business logic independent of infrastructure, build modular services, and scale through clean abstractions.**
