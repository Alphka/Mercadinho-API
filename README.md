# Mercadinho API

**English** | [Português (BR)](README.pt-BR.md)

[![CI](https://github.com/Alphka/Mercadinho-API/actions/workflows/ci.yml/badge.svg)](https://github.com/Alphka/Mercadinho-API/actions/workflows/ci.yml) [![License: ISC](https://img.shields.io/badge/license-ISC-blue.svg)](LICENSE.md)

REST API built as a learning project for the Go language and to deepen PostgreSQL knowledge. It is inspired by [an Android app](https://github.com/Alphka/MercadoCRUDJava), adapting its features to a REST architecture.

## Features

- CRUD for customers, products and purchases
- Customer purchases (create and list, paginated)
- Per-customer listing of pending and overdue payments
- Product soft deletion
- Request validation with custom rules (`date`, `price`, `phone`)
- Interactive Swagger documentation
- Database migrations applied automatically on startup
- Permissive CORS (all origins allowed)

## Tech stack

- **Go** with [Gin](https://github.com/gin-gonic/gin)
- **Swagger** via [swaggo](https://github.com/swaggo/swag)
- **golang-migrate** for database migrations
- **PostgreSQL** with [pgx](https://github.com/jackc/pgx)

## Requirements

- [Go](https://go.dev/dl/) (the `go.mod` declares `go 1.26.2`)
- [PostgreSQL](https://www.postgresql.org/) running and reachable
- [`swag`](https://github.com/swaggo/swag) CLI (used by the start scripts to generate the Swagger docs)

## Installation

```bash
git clone https://github.com/Alphka/Mercadinho-API.git
cd Mercadinho-API

go install github.com/swaggo/swag/cmd/swag@latest
go mod download
```

### Configuration

Copy the example file and fill in the environment variables:

```bash
cp .env.example .env
```

| Variable | Description | Default |
|---|---|---|
| `DATABASE_URL` | PostgreSQL connection string (required) | none |
| `HOST` | Host used by the server and by Swagger | `localhost` |
| `PORT` | Server port | `8000` |

> [!NOTE]
> If the database named in `DATABASE_URL` does not exist, the API tries to create it on startup (it connects to the `postgres` database of the same server for that), so the configured user needs permission to do so.

## Usage

The scripts generate the Swagger docs, build the binary into `output/` and start the server:

```bash
# Bash
./bin/start.sh

# Windows PowerShell
./bin/start.ps1

# CMD
./bin/start.bat
```

Once the server is running, the Swagger UI is available at:

```
http://localhost:8000/swagger/index.html
```

### Database

Migrations live in [`migrations/`](migrations) and are embedded in the binary. They are applied automatically on startup, except when running on Vercel (detected through the `VERCEL` environment variable). To apply them only, without starting the server:

```bash
go run . migrate
```

### Endpoints

All routes are prefixed with `/api/v1`. The full contract is in the Swagger UI.

| Resource | Routes |
|---|---|
| Customers | `GET /customers`, `GET /customers/:id`, `POST /customers`, `PUT /customers/:id`, `DELETE /customers/:id` |
| Payments | `GET /customers/:id/pending-payments`, `GET /customers/:id/overdue-payments` |
| Customer purchases | `GET /customers/:id/purchases`, `POST /customers/:id/purchases` |
| Products | `GET /products`, `GET /products/:id`, `POST /products`, `PUT /products/:id`, `DELETE /products/:id` |
| Purchases | `GET /purchases`, `GET /purchases/:id`, `PUT /purchases/:id`, `DELETE /purchases/:id` |

List endpoints accept the `page` (default `1`) and `pageSize` (default `20`) query parameters.

### Hot reload (optional)

[nodemon](https://nodemon.io/) can restart the server on file changes (see [`nodemon.json`](nodemon.json)). It requires [Node.js](https://nodejs.org/):

```bash
npm install -g nodemon
nodemon
```

## Project structure

```
bin/          Start scripts (Bash, PowerShell, CMD)
controllers/  Request handlers (annotated for Swagger)
database/     PostgreSQL connection and migration runner
docs/         Generated Swagger files (swag init)
dtos/         Request payloads and validation rules
helpers/      Validation, date and phone helpers
migrations/   SQL migrations (golang-migrate)
models/       Response and database models
routes/       Route registration
```

## Testing

The repository has no automated tests yet. A GitHub Actions workflow ([`.github/workflows/ci.yml`](.github/workflows/ci.yml)) runs the `gofmt` check, `go build`, `go vet` and `go test` on pushes to `main` and on pull requests. A GitHub Actions workflow ([`.github/workflows/ci.yml`](.github/workflows/ci.yml)) runs the `gofmt` check, `go build`, `go vet` and `go test` on pushes to `main` and on pull requests.

## License

Distributed under the [ISC License](LICENSE.md).
