<p align="center">
<img src="https://github.com/kongali1720/KongWallet-Payment-Gateway-API/blob/main/kop_surat.jpg" width="100%">
</p>

<p align="center">
<img src="https://raw.githubusercontent.com/kongali1720/KongPay/main/KP-Logo-01.png" width="100%">
</p>

<h1 align="center">
KONGPAY
</h1>

<p align="center">
<b>Autonomous Digital Financial Infrastructure</b><br>
Open Source • Enterprise Ready • API First • Cloud Native
</p>

<p align="center">

<img src="https://img.shields.io/github/stars/kongali1720/KongPay?style=for-the-badge&logo=github&color=FFD700">
<img src="https://img.shields.io/github/forks/kongali1720/KongPay?style=for-the-badge&logo=github&color=0EA5E9">
<img src="https://img.shields.io/github/watchers/kongali1720/KongPay?style=for-the-badge&logo=github&color=8B5CF6">

</p>

<p align="center">

<img src="https://img.shields.io/github/license/kongali1720/KongPay?style=for-the-badge&color=22C55E">
<img src="https://img.shields.io/github/last-commit/kongali1720/KongPay?style=for-the-badge&color=EF4444">
<img src="https://img.shields.io/github/repo-size/kongali1720/KongPay?style=for-the-badge&color=F97316">

</p>

<p align="center">

<img src="https://img.shields.io/badge/Go-1.26-00ADD8?style=for-the-badge&logo=go&logoColor=white">
<img src="https://img.shields.io/badge/Gin-Web_Framework-008ECF?style=for-the-badge">
<img src="https://img.shields.io/badge/PostgreSQL-17-4169E1?style=for-the-badge&logo=postgresql&logoColor=white">
<img src="https://img.shields.io/badge/Redis-7-DC382D?style=for-the-badge&logo=redis&logoColor=white">
<img src="https://img.shields.io/badge/Docker-Ready-2496ED?style=for-the-badge&logo=docker&logoColor=white">

</p>

<p align="center">

<img src="https://img.shields.io/badge/Open%20Source-Yes-success?style=for-the-badge">
<img src="https://img.shields.io/badge/Enterprise-Ready-blue?style=for-the-badge">
<img src="https://img.shields.io/badge/REST-API-orange?style=for-the-badge">
<img src="https://img.shields.io/badge/GitHub_Actions-CI-2088FF?style=for-the-badge&logo=githubactions&logoColor=white">
<img src="https://img.shields.io/badge/Clean-Architecture-success?style=for-the-badge">

</p>

---

# 🚀 Executive Summary

**KongPay** is an open-source payment infrastructure built with **Go**, designed to provide a modular, scalable, and secure foundation for modern financial applications.

The project follows a clean architecture consisting of **Handlers**, **Services**, **Repositories**, and **PostgreSQL**, enabling developers to build digital wallets, merchant systems, payment gateways, settlement services, and financial APIs on top of a production-ready backend.

KongPay is designed around modern engineering principles:

- API First
- Clean Architecture
- Repository Pattern
- Dependency Injection
- PostgreSQL Persistence
- Redis Caching
- Docker Deployment
- GitHub Actions CI
- Open Source Collaboration

---

# 🌍 Why KongPay?

Modern payment ecosystems require scalable and secure backend infrastructure.

KongPay aims to provide reusable financial service components that developers and organizations can extend without building everything from scratch.

The platform focuses on:

- Security
- Scalability
- Performance
- Simplicity
- Developer Experience
- Maintainability

---

# 🎯 Project Goals

- Build reusable payment infrastructure
- Develop a production-ready Wallet Service
- Provide Merchant APIs
- Build Transaction Ledger
- Develop Settlement Engine
- Support JWT Authentication
- Support QRIS & Virtual Account integration
- Encourage Open Source Collaboration

---

# 🏗 Architecture

```
                Internet
                    │
                    ▼
              REST API (Gin)
                    │
        ┌───────────┴───────────┐
        ▼                       ▼
   Wallet Handler         Merchant Handler
        │                       │
        ▼                       ▼
   Wallet Service        Merchant Service
        │                       │
        └───────────┬───────────┘
                    ▼
             Repository Layer
                    │
        ┌───────────┴───────────┐
        ▼                       ▼
      PostgreSQL             Redis
```

---

# 🧩 Current Features

- RESTful Wallet API
- PostgreSQL Integration
- Redis Integration
- Docker Environment
- Repository Pattern
- Dependency Injection
- UUID Support
- GitHub Actions CI
- Environment Configuration
- Health Check Endpoint

---

# ⚙ Technology Stack

| Layer | Technology |
|--------|------------|
| Language | Go 1.26 |
| Framework | Gin |
| Database | PostgreSQL 17 |
| Cache | Redis 7 |
| Driver | pgx v5 |
| Container | Docker |
| CI/CD | GitHub Actions |
| API | REST |
| Version Control | Git |

---

# 📂 Repository Structure

```text
kongpay/
│
├── cmd/
│   └── kongpay/
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

# 📡 REST API

## Health Check

```http
GET /health
```

## Create Wallet

```http
POST /api/v1/wallets
```

## Get Wallet

```http
GET /api/v1/wallets/{id}
```

### Coming Soon

- GET /wallets
- PUT /wallets/{id}
- DELETE /wallets/{id}

---

# 🚀 Getting Started

## Clone Repository

```bash
git clone https://github.com/kongali1720/KongPay.git

cd KongPay
```

---

## Configure Environment

```bash
cp .env.example .env
```

Edit `.env` as needed.

---

## Start Infrastructure

```bash
docker compose up -d
```

---

## Run KongPay

```bash
go run ./cmd/kongpay
```

---

## Test Health Endpoint

```bash
curl http://localhost:8080/health
```

---

## Create Wallet

```bash
curl -X POST http://localhost:8080/api/v1/wallets \
-H "Content-Type: application/json" \
-d '{
  "user_id":"550e8400-e29b-41d4-a716-446655440000",
  "currency":"IDR"
}'
```

---

# 🗺 Roadmap

## Phase 1 ✅

- Go Backend
- Gin Framework
- PostgreSQL
- Redis
- Docker
- Wallet Create API
- Wallet Persistence
- GitHub Actions

## Phase 2 🚧

- Wallet CRUD
- JWT Authentication
- Merchant API

## Phase 3

- Payment Engine
- Transaction Ledger
- Settlement Service

## Phase 4

- QRIS
- Virtual Account
- Webhook
- Notification Service

## Phase 5

- API Gateway
- Kubernetes
- Production Deployment

---

# 🏆 Project Milestones

- ✅ Docker Infrastructure
- ✅ PostgreSQL Integration
- ✅ Redis Integration
- ✅ Repository Pattern
- ✅ Service Layer
- ✅ Dependency Injection
- ✅ Wallet Persistence
- ✅ GitHub Actions CI

---

<p align="center">

<a href="https://github.com/kongali1720/KongPay/stargazers">
<img src="https://img.shields.io/badge/⭐-Star%20Repository-FFD700?style=for-the-badge">
</a>

<a href="https://github.com/kongali1720/KongPay/fork">
<img src="https://img.shields.io/badge/🍴-Fork-0EA5E9?style=for-the-badge">
</a>

<a href="https://github.com/kongali1720/KongPay/issues">
<img src="https://img.shields.io/badge/🐞-Issues-EF4444?style=for-the-badge">
</a>

<a href="https://github.com/kongali1720/KongPay/pulls">
<img src="https://img.shields.io/badge/🚀-Pull_Request-22C55E?style=for-the-badge">
</a>

</p>

---

# ❤️ Support KongPay

If KongPay helps your learning, research, or development journey, consider supporting the project.

Every contribution helps improve:

- Documentation
- Infrastructure
- Security
- Developer Experience
- Open Source Sustainability

<p align="center">

<a href="https://www.paypal.com/paypalme/bungtempong99">
<img src="https://img.shields.io/badge/☕-Support_Development-FFDD00?style=for-the-badge&logo=buymeacoffee&logoColor=000000">
</a>

</p>

---

<h2 align="center">
Built with Go ❤️ by KONGALI1720
</h2>

<p align="center">
Autonomous Digital Financial Infrastructure
</p>

<p align="center">
Go • Gin • PostgreSQL • Redis • Docker • GitHub Actions
</p>
