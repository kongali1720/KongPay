# 📘 KongPay REST API Documentation

KongPay exposes a RESTful API built with **Go**, **Gin**, and **PostgreSQL**.

The API is designed around JSON request and response bodies with predictable HTTP status codes.

---

# 🌐 Base URL

Development

```text
http://localhost:8080
```

API Version

```text
/api/v1
```

Content Type

```http
Content-Type: application/json
```

---

# 📋 HTTP Status Codes

| Code | Meaning |
|------|---------|
| 200 | OK |
| 201 | Created |
| 400 | Bad Request |
| 404 | Not Found |
| 500 | Internal Server Error |

---

# ❤️ Health Check

## Endpoint

```http
GET /health
```

### Example

```bash
curl http://localhost:8080/health
```

### Success Response

```json
{
  "status": "ok"
}
```

---

# 👛 Wallet API

## Create Wallet

### Endpoint

```http
POST /api/v1/wallets
```

### Request

```json
{
  "user_id":"550e8400-e29b-41d4-a716-446655440000",
  "currency":"IDR"
}
```

### Example

```bash
curl -X POST http://localhost:8080/api/v1/wallets \
-H "Content-Type: application/json" \
-d '{
  "user_id":"550e8400-e29b-41d4-a716-446655440000",
  "currency":"IDR"
}'
```

### Response

```json
{
  "id":"11c58304-cffc-4073-90a9-a2e6b7aac7e2",
  "user_id":"550e8400-e29b-41d4-a716-446655440000",
  "balance":0,
  "currency":"IDR",
  "status":"ACTIVE",
  "created_at":"2026-08-02T09:42:31Z",
  "updated_at":"2026-08-02T09:42:31Z"
}
```

---

## List Wallets

### Endpoint

```http
GET /api/v1/wallets
```

### Example

```bash
curl http://localhost:8080/api/v1/wallets
```

### Response

```json
[
  {
    "id":"11c58304-cffc-4073-90a9-a2e6b7aac7e2",
    "user_id":"550e8400-e29b-41d4-a716-446655440000",
    "balance":0,
    "currency":"IDR",
    "status":"ACTIVE",
    "created_at":"2026-08-02T09:42:31Z",
    "updated_at":"2026-08-02T09:42:31Z"
  }
]
```

---

## Get Wallet

### Endpoint

```http
GET /api/v1/wallets/{id}
```

### Example

```bash
curl http://localhost:8080/api/v1/wallets/11c58304-cffc-4073-90a9-a2e6b7aac7e2
```

---

## Update Wallet

### Endpoint

```http
PUT /api/v1/wallets/{id}
```

### Request

```json
{
  "currency":"USD",
  "balance":100000,
  "status":"ACTIVE"
}
```

### Example

```bash
curl -X PUT http://localhost:8080/api/v1/wallets/11c58304-cffc-4073-90a9-a2e6b7aac7e2 \
-H "Content-Type: application/json" \
-d '{
  "currency":"USD",
  "balance":100000,
  "status":"ACTIVE"
}'
```

---

## Delete Wallet

### Endpoint

```http
DELETE /api/v1/wallets/{id}
```

### Example

```bash
curl -X DELETE http://localhost:8080/api/v1/wallets/11c58304-cffc-4073-90a9-a2e6b7aac7e2
```

### Response

```json
{
  "message":"wallet deleted successfully"
}
```

---

# 📊 Current Endpoint Matrix

| Method | Endpoint | Description | Status |
|---------|----------|-------------|:------:|
| GET | /health | Health Check | ✅ |
| POST | /api/v1/wallets | Create Wallet | ✅ |
| GET | /api/v1/wallets | List Wallets | ✅ |
| GET | /api/v1/wallets/{id} | Get Wallet | ✅ |
| PUT | /api/v1/wallets/{id} | Update Wallet | ✅ |
| DELETE | /api/v1/wallets/{id} | Delete Wallet | ✅ |

---

# 🔐 Authentication (Planned)

Future API versions will support:

- JWT Authentication
- Refresh Tokens
- Role-Based Access Control (RBAC)
- API Keys for Merchants

---

# 🚧 Planned API Modules

| Module | Status |
|---------|:------:|
| Authentication | 🚧 |
| Merchant | ⏳ |
| Payments | ⏳ |
| Transactions | ⏳ |
| Settlement | ⏳ |
| Notifications | ⏳ |
| QRIS | ⏳ |
| Virtual Account | ⏳ |

---

# 📈 API Design Principles

- RESTful
- JSON-based
- Stateless
- Versioned API
- Predictable Status Codes
- Consistent Response Format

---

# 🚀 Future Enhancements

- Swagger / OpenAPI
- Request Validation
- Pagination
- Filtering
- Sorting
- Rate Limiting
- API Version Negotiation
- Idempotency Support

---

# ❤️ API Philosophy

> **Simple, predictable, secure, and developer-friendly APIs built for modern digital financial infrastructure.**
