# 🤝 Contributing to KongPay

Thank you for your interest in contributing to **KongPay**.

We welcome developers, students, researchers, and the open-source community to help improve this project.

---

## 🚀 Development Setup

Clone the repository:

```bash
git clone https://github.com/kongali1720/KongPay.git
cd KongPay
```

Download dependencies:

```bash
go mod download
```

Start infrastructure:

```bash
docker compose up -d
```

Run the application:

```bash
go run ./cmd/kongpay
```

---

## 📌 Coding Standards

Please follow these principles:

- Idiomatic Go
- Clean Architecture
- Repository Pattern
- Dependency Injection
- RESTful API Design

Before submitting changes, run:

```bash
go fmt ./...
go test ./...
go build ./...
```

---

## 🌿 Branch Strategy

- `main` → Stable
- `develop` → Development
- `feature/*` → New features
- `hotfix/*` → Critical fixes

---

## ✅ Pull Request Checklist

Before opening a Pull Request:

- Update your branch with the latest `main`
- Run `go fmt ./...`
- Run `go test ./...`
- Update documentation if necessary
- Update CHANGELOG.md for significant changes

---

## 🐛 Reporting Bugs

When reporting bugs, please include:

- Operating system
- Go version
- Docker version
- Steps to reproduce
- Expected behavior
- Actual behavior

---

## ❤️ Thank You

Every contribution helps improve KongPay.
