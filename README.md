```{=html}
<p align="center">
```
`<img src="https://github.com/kongali1720/KongWallet-Payment-Gateway-API/blob/main/kop_surat.jpg" width="100%">`{=html}

```{=html}
</p>
```
```{=html}
<p align="center">
```
`<img src="https://raw.githubusercontent.com/kongali1720/KongPay/main/KP-Logo-01.png" width="100%">`{=html}

```{=html}
</p>
```
```{=html}
<h1 align="center">
```
KONGPAY

```{=html}
</h1>
```
```{=html}
<p align="center">
```
`<b>`{=html}Autonomous Digital
FinancialInfrastructure`</b>`{=html}`<br>`{=html} Open Source • API
First • CloudNative • Modular Architecture

```{=html}
</p>
```
```{=html}
<p align="center">
```
`<img src="https://img.shields.io/github/stars/kongali1720/KongPay?style=for-the-badge&logo=github&color=FFD700">`{=html}
`<img src="https://img.shields.io/github/forks/kongali1720/KongPay?style=for-the-badge&logo=github&color=0EA5E9">`{=html}
`<img src="https://img.shields.io/github/watchers/kongali1720/KongPay?style=for-the-badge&logo=github&color=8B5CF6">`{=html}

```{=html}
</p>
```
```{=html}
<p align="center">
```
`<img src="https://img.shields.io/github/license/kongali1720/KongPay?style=for-the-badge&color=22C55E">`{=html}
`<img src="https://img.shields.io/github/last-commit/kongali1720/KongPay?style=for-the-badge&color=EF4444">`{=html}
`<img src="https://img.shields.io/github/repo-size/kongali1720/KongPay?style=for-the-badge&color=F97316">`{=html}
`<img src="https://img.shields.io/github/v/release/kongali1720/KongPay?style=for-the-badge&color=6366F1">`{=html}

```{=html}
</p>
```
```{=html}
<p align="center">
```
`<img src="https://img.shields.io/badge/Go-1.26-00ADD8?style=for-the-badge&logo=go&logoColor=white">`{=html}
`<img src="https://img.shields.io/badge/Gin-Web_Framework-008ECF?style=for-the-badge">`{=html}
`<img src="https://img.shields.io/badge/PostgreSQL-17-4169E1?style=for-the-badge&logo=postgresql&logoColor=white">`{=html}
`<img src="https://img.shields.io/badge/Redis-7-DC382D?style=for-the-badge&logo=redis&logoColor=white">`{=html}
`<img src="https://img.shields.io/badge/Docker-Ready-2496ED?style=for-the-badge&logo=docker&logoColor=white">`{=html}

```{=html}
</p>
```
```{=html}
<p align="center">
```
`<img src="https://img.shields.io/badge/Open%20Source-Yes-success?style=for-the-badge">`{=html}
`<img src="https://img.shields.io/badge/REST-API-orange?style=for-the-badge">`{=html}
`<img src="https://img.shields.io/badge/GitHub_Actions-CI-2088FF?style=for-the-badge&logo=githubactions&logoColor=white">`{=html}
`<img src="https://img.shields.io/badge/Clean-Architecture-success?style=for-the-badge">`{=html}

```{=html}
</p>
```
🚀 Executive Summary

KongPay is an open-source digital payment infrastructure projectbuilt
with Go and Gin. It is designed around modular backendcomponents that
can evolve into wallet, merchant, payment, ledger,settlement,
notification, and financial API services.

The current implementation establishes a working Wallet REST API
backedby PostgreSQL. The application is organized into handler,
service,repository, model, database, configuration, and routing layers
so thatbusiness logic and persistence concerns remain separated as the
projectgrows.

KongPay emphasizes API-first development, modularity,
maintainability,database-backed persistence, containerized development
infrastructure,and automated CI checks.

Development notice: KongPay is under active development. Featuresmarked
as planned or in progress in this README should not beinterpreted as
production-complete financial services.

📊 Project Status

Capability Status

Go Backend ✅ ImplementedGin REST API ✅ ImplementedPostgreSQL
Connection ✅ ImplementedWallet Persistence ✅ ImplementedWallet CRUD ✅
Implemented & TestedRedis Container / Infrastructure ✅ AvailableDocker
Development Infrastructure ✅ AvailableGitHub Actions ✅ AvailableJWT
Authentication 🚧 PlannedSwagger / OpenAPI 🚧 PlannedSQL Migration
Workflow 🚧 PlannedAutomated Unit Tests 🚧 PlannedMerchant Service ⏳
RoadmapPayment Engine ⏳ RoadmapTransaction Ledger ⏳ RoadmapSettlement
Service ⏳ RoadmapQRIS Integration ⏳ RoadmapVirtual Account Integration
⏳ Roadmap

🌍 Why KongPay?

Digital payment systems typically involve many distinct concerns:account
balances, merchants, transactions, authentication,
settlement,auditability, integrations, infrastructure, and operational
controls.

KongPay explores how those concerns can be separated into
reusableservices and clean application layers instead of being
concentrated inone tightly coupled codebase.

The project focuses on:

Modularity

Maintainability

API-first design

Clear separation of concerns

Developer experience

Extensible financial-service architecture

Open-source collaboration

🎯 Project Goals

Build reusable payment-infrastructure components.

Maintain a clean Go backend architecture.

Provide a persistent Wallet Service.

Add secure identity and authentication capabilities.

Develop merchant-management APIs.

Introduce a transaction ledger.

Develop payment and settlement workflows.

Add documented integration interfaces.

Improve testing, migrations, deployment, and CI/CD over
successivereleases.

✨ Current Features

RESTful Wallet API

Create Wallet

List Wallets

Get Wallet by ID

Update Wallet

Delete Wallet

PostgreSQL persistence

UUID wallet identifiers

Wallet balance, currency, and status fields

Repository pattern

Service layer

Handler layer

Dependency injection

Environment-based configuration

Health endpoint

Docker-based PostgreSQL and Redis development infrastructure

GitHub Actions workflow

🏗 Architecture

                       Client
                         │
                         ▼
                   REST API / Gin
                         │
                         ▼
                      Router
                         │
                         ▼
                  Wallet Handler
                         │
                         ▼
                  Wallet Service
                         │
                         ▼
                Wallet Repository
                         │
                         ▼
                    PostgreSQL

        Supporting development infrastructure
                         │
                         └──────── Redis

The Wallet request path implemented today is:

HTTP Request │ ▼ Gin Router │ ▼ Wallet Handler │ ▼ Wallet Service │ ▼
Wallet Repository │ ▼ PostgreSQL

📈 Mermaid Architecture

flowchart TD A\[Client\] --\> B\[Gin REST API\] B --\> C\[Router\] C
--\> D\[Wallet Handler\] D --\> E\[Wallet Service\] E --\> F\[Wallet
Repository\] F --\> G\[(PostgreSQL)\] H\[(Redis)\] -. Supporting
Infrastructure .-\> E

🔄 Wallet Request Flow

sequenceDiagram participant Client participant Router participant
Handler participant Service participant Repository participant
PostgreSQL

    Client->>Router: HTTP request
    Router->>Handler: Dispatch endpoint
    Handler->>Service: Validate/execute operation
    Service->>Repository: Persistence operation
    Repository->>PostgreSQL: SQL query
    PostgreSQL-->>Repository: Result
    Repository-->>Service: Domain data
    Service-->>Handler: Result
    Handler-->>Client: JSON response

⚙ Technology Stack

Layer Technology

Language Go 1.26HTTP Framework GinDatabase PostgreSQL 17PostgreSQL
Driver pgx v5Cache / Supporting Infrastructure Redis 7Identifiers Google
UUIDContainers DockerCI GitHub ActionsAPI Style RESTVersion Control Git

📂 Repository Structure

kongpay-project/ │ ├── cmd/ │ └── kongpay/ │ └── main.go │ ├── internal/
│ ├── config/ │ ├── database/ │ ├── handlers/ │ ├── middleware/ │ ├──
models/ │ ├── repositories/ │ ├── router/ │ ├── services/ │ └── utils/ │
├── migrations/ ├── docker/ ├── docs/ ├── tests/ ├── .github/ │ └──
workflows/ │ ├── docker-compose.yml ├── go.mod ├── go.sum └── README.md

Layer Responsibilities

Layer Responsibility

cmd/kongpay Application entry pointinternal/config
Environment/application configurationinternal/database PostgreSQL
connection lifecycleinternal/handlers HTTP request/response
handlinginternal/models Application data modelsinternal/repositories
Database persistence operationsinternal/services Business/application
logicinternal/router Endpoint registration and dependency
wiringinternal/middleware Reserved for cross-cutting HTTP
middlewaremigrations Planned SQL migration filesdocs Extended project
documentationtests Automated testing area

🗄 Database Schema

The Wallet CRUD implementation uses the following table shape:

CREATE TABLE wallets ( id UUID PRIMARY KEY, user_id UUID NOT NULL,
balance NUMERIC(20,2) NOT NULL DEFAULT 0, currency VARCHAR(10) NOT NULL,
status VARCHAR(20) NOT NULL DEFAULT 'ACTIVE', created_at TIMESTAMP NOT
NULL DEFAULT NOW(), updated_at TIMESTAMP NOT NULL DEFAULT NOW() );

Indexes:

CREATE INDEX idx_wallets_user_id ON wallets(user_id); CREATE INDEX
idx_wallets_status ON wallets(status);

Wallet Fields

Column Type Purpose

id UUID Unique wallet identifieruser_id UUID Owner/user
identifierbalance NUMERIC(20,2) Wallet balancecurrency VARCHAR(10)
Wallet currencystatus VARCHAR(20) Wallet statuscreated_at TIMESTAMP
Creation timestampupdated_at TIMESTAMP Last update timestamp

📡 REST API

Base path:

/api/v1

Endpoint Matrix

Method Endpoint Description Status

GET /health Application health check ✅POST /api/v1/wallets Create
wallet ✅GET /api/v1/wallets List wallets ✅GET /api/v1/wallets/{id} Get
wallet by ID ✅PUT /api/v1/wallets/{id} Update wallet ✅DELETE
/api/v1/wallets/{id} Delete wallet ✅

🧪 API Examples

Create Wallet

curl -X POST http://localhost:8080/api/v1/wallets\
-H "Content-Type: application/json"\
-d '{ "user_id":"550e8400-e29b-41d4-a716-446655440000", "currency":"IDR"
}'

Example response shape:

{ "id": "11c58304-cffc-4073-90a9-a2e6b7aac7e2", "user_id":
"550e8400-e29b-41d4-a716-446655440000", "balance": 0, "currency": "IDR",
"status": "ACTIVE", "created_at": "2026-08-02T09:42:31.691793Z",
"updated_at": "2026-08-02T09:42:31.691793Z" }

List Wallets

curl http://localhost:8080/api/v1/wallets

Example response:

\[ { "id": "11c58304-cffc-4073-90a9-a2e6b7aac7e2", "user_id":
"550e8400-e29b-41d4-a716-446655440000", "balance": 0, "currency": "IDR",
"status": "ACTIVE", "created_at": "2026-08-02T09:42:31.691793Z",
"updated_at": "2026-08-02T09:42:31.691793Z" }\]

Get Wallet

Replace the UUID with an existing wallet ID.

curl
http://localhost:8080/api/v1/wallets/11c58304-cffc-4073-90a9-a2e6b7aac7e2

Update Wallet

curl -X PUT
http://localhost:8080/api/v1/wallets/11c58304-cffc-4073-90a9-a2e6b7aac7e2\
-H "Content-Type: application/json"\
-d '{ "currency":"USD", "balance":100000, "status":"ACTIVE" }'

Delete Wallet

curl -X DELETE
http://localhost:8080/api/v1/wallets/11c58304-cffc-4073-90a9-a2e6b7aac7e2

Successful deletion returns a message such as:

{ "message": "wallet deleted successfully" }

🚀 Getting Started

Requirements

Recommended local development environment:

Go

Git

Docker

Docker Compose

curl

Clone Repository

git clone https://github.com/kongali1720/KongPay.git cd KongPay

Install Go Dependencies

go mod download

Configure Environment

Create or configure .env with values appropriate for your environment.

Example development configuration:

APP_NAME=KongPay APP_ENV=development PORT=8080

DB_HOST=localhost DB_PORT=5432 DB_NAME=kongpay DB_USER=postgres
DB_PASSWORD=postgres DB_SSLMODE=disable

Do not commit real production passwords or secrets to Git.

Start Infrastructure

docker compose up -d

Check containers:

docker ps

Run KongPay

go run ./cmd/kongpay

Expected service address:

http://localhost:8080

Test Health

curl http://localhost:8080/health

🐳 Docker Development Infrastructure

The development environment includes PostgreSQL and Redis containers.

flowchart LR A\[KongPay Go API\] --\> B\[(PostgreSQL 17)\] A -. future
caching / coordination .-\> C\[(Redis 7)\]

Useful commands:

docker compose up -d docker ps docker compose logs docker compose down

🛠 Development Workflow

Format source code:

go fmt ./...

Run tests/build checks currently available through Go tooling:

go test ./...

Build packages:

go build ./...

Run the API:

go run ./cmd/kongpay

Recommended pre-commit sequence:

go fmt ./... go test ./... go build ./... git status

🔁 Git Workflow

Check repository state:

git status

Synchronize before pushing when the remote branch has changed:

git pull --rebase origin main

Commit changes:

git add . git commit -m "feat(wallet): complete wallet CRUD"

Push:

git push origin main

⚙ GitHub Actions

The repository uses GitHub Actions for automated Go checks.

The CI roadmap is to progressively enforce:

Checkout │ ▼ Setup Go │ ▼ Format Check │ ▼ Vet │ ▼ Tests │ ▼ Build

Future improvements may include linting, coverage reporting,
containerbuilds, security scanning, and automated release workflows.

🔐 Security Direction

Security capabilities will be added incrementally as the
platformmatures.

Planned areas include:

JWT authentication

Refresh-token lifecycle

Authorization middleware

Role-based access control

Request validation

Secure secret management

Audit logging

Rate limiting

TLS deployment

Dependency scanning

Database transaction controls

KongPay should not be treated as a finished regulated payment
platformmerely because individual API capabilities are implemented.

🗃 Database Migration Roadmap

Database migrations are planned so deployments do not depend on
manuallycreating tables through psql.

Proposed structure:

migrations/ ├── 001_create_wallets.sql ├── 002_create_merchants.sql ├──
003_create_transactions.sql ├── 004_create_ledger_entries.sql └──
005_create_settlements.sql

🧪 Testing Roadmap

The current project compiles successfully with:

go test ./...

Dedicated automated tests are a planned improvement.

Target coverage areas:

Wallet service tests

Wallet repository tests

Wallet handler tests

Validation tests

Integration tests

Database tests

Authentication tests

🧭 Roadmap

v0.1.0 --- Initial Foundation ✅

Foundation milestone covering the initial Go project structure
anddevelopment infrastructure.

v0.2.0 --- Wallet CRUD ✅

Implemented and manually exercised:

Create Wallet

List Wallets

Get Wallet

Update Wallet

Delete Wallet

PostgreSQL persistence

Handler → Service → Repository flow

Next Development Milestones 🚧

Planned development areas include:

JWT Authentication

Swagger / OpenAPI

SQL migrations

Automated tests

Production-oriented Docker build

CI/CD enhancements

Merchant Service

Payment Engine

Transaction Ledger

Settlement Service

Webhooks and notifications

QRIS / Virtual Account integration research and implementation
whereapplicable

Administration interfaces

Version numbers for future milestones may be assigned as
thosecapabilities are implemented and released.

🌐 Future KongPay Ecosystem

KongPay │ ├── Wallet Service ✅ Current foundation │ ├── Authentication
🚧 Planned ├── Merchant Service ⏳ Planned ├── Payment Engine ⏳ Planned
├── Transaction Ledger ⏳ Planned ├── Settlement Service ⏳ Planned ├──
Webhook Service ⏳ Planned ├── Notification Service ⏳ Planned ├── QRIS
Integration ⏳ Planned ├── Virtual Account ⏳ Planned ├── API Gateway ⏳
Planned ├── Administration Dashboard ⏳ Planned └── Developer SDKs ⏳
Planned

🤝 Contributing

Contributions, bug reports, documentation improvements, and
engineeringdiscussions are welcome.

Suggested workflow:

git checkout -b feature/my-feature

Make changes and verify them:

go fmt ./... go test ./... go build ./...

Commit:

git commit -m "feat: describe the change"

Then push the branch and open a pull request.

🛡 Responsible Security Reporting

Please avoid publishing exploitable security findings in public
issuesbefore maintainers have an opportunity to investigate them.

A dedicated SECURITY.md policy is recommended as the project expands.

📜 License

KongPay is distributed under the license included in this repository.

Refer to the repository LICENSE file for the authoritative licenseterms.

⭐ Support the Project

If KongPay is useful for your learning, research, or development
work,you can support the project by starring the repository,
contributingcode or documentation, reporting reproducible issues,
reviewing pullrequests, or helping improve tests and security.

```{=html}
<p align="center">
```
`<a href="https://github.com/kongali1720/KongPay/stargazers">`{=html}
`<img src="https://img.shields.io/badge/⭐-Star%20Repository-FFD700?style=for-the-badge">`{=html}`</a>`{=html}

`<a href="https://github.com/kongali1720/KongPay/fork">`{=html}
`<img src="https://img.shields.io/badge/🍴-Fork-0EA5E9?style=for-the-badge">`{=html}`</a>`{=html}

`<a href="https://github.com/kongali1720/KongPay/issues">`{=html}
`<img src="https://img.shields.io/badge/🐞-Issues-EF4444?style=for-the-badge">`{=html}`</a>`{=html}

`<a href="https://github.com/kongali1720/KongPay/pulls">`{=html}
`<img src="https://img.shields.io/badge/🚀-Pull_Request-22C55E?style=for-the-badge">`{=html}`</a>`{=html}

```{=html}
</p>
```
```{=html}
<p align="center">
```
`<a href="https://www.paypal.com/paypalme/bungtempong99">`{=html}
`<img src="https://img.shields.io/badge/☕-Support_Development-FFDD00?style=for-the-badge&logo=buymeacoffee&logoColor=000000">`{=html}`</a>`{=html}

```{=html}
</p>
```
👨‍💻 Project Author

```{=html}
<h2 align="center">
```
Built with Go ❤️ by KONGALI1720

```{=html}
</h2>
```
```{=html}
<p align="center">
```
`<b>`{=html}Autonomous Digital Financial Infrastructure`</b>`{=html}

```{=html}
</p>
```
```{=html}
<p align="center">
```
Go • Gin • PostgreSQL • Redis • Docker • GitHub Actions

```{=html}
</p>
```
```{=html}
<p align="center">
```
`<b>`{=html}KongPay --- Building modular foundations for modern
digitalpayment infrastructure.`</b>`{=html}

```{=html}
</p>
```
