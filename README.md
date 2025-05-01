# 💳 Card Validator API

A RESTful API service for validating credit card information, built with **Go** and the **Echo** framework.  
The service validates card numbers using the **Luhn algorithm** and checks expiration dates.

---

## ✨ Features

- ✅ Validates credit card number using **Luhn algorithm**
- 🗓️ Checks expiration **month** (1–12) and **year** (not expired)
- 📢 Returns clear validation results with **error codes and messages**
- 🧼 Clean architecture with **dependency injection**
- 🐳 **Docker** support
- 📚 **Swagger** documentation
- 🧪 Unit tests

---

## 🛠️ Prerequisites

- 🧑‍💻 Go `1.23.4`
- 🐋 Docker
- 🧰 Make (optional, for using Makefile commands)

---

## ⚙️ Installation

Clone the repository:

```bash
git clone https://github.com/your-username/card-validator-api.git
cd card-validator-api
```

Install dependencies:

```bash
go mod download
```

Create `.env` file:

```bash
echo "HTTP_PORT=8080" > .env
```

---

## 🚀 Running the Application

### 🖥️ Local Development

```bash
make up
# or
ENV=dev go run cmd/server/main.go
```

### 🐳 Docker

```bash
make docker-build
# or
docker build -t card-validator-api . && docker run --rm -p 8080:8080 card-validator-api
```

---

## 📡 API Endpoints

### `POST /card/validate`
Validates a credit card's number and expiration date.

**Request Body**:

```json
{
  "number": "4111111111111111",
  "exp_month": 12,
  "exp_year": 2028
}
```

**Response (✅ Success)**:

```json
{
  "valid": true
}
```

**Response (❌ Error)**:

```json
{
  "valid": false,
  "error": {
    "code": "001",
    "message": "Invalid card number"
  }
}
```

---

## ❗ Error Codes

| Code | Description               |
|------|---------------------------|
| 001  | ❌ Invalid card number     |
| 002  | 📆 Invalid expiration month |
| 003  | 🕒 Card expired             |

---

## 👨‍💻 Development

### 📖 Generate Swagger Docs

```bash
make swagger
```

### 🧪 Run Tests

```bash
make test
```
