# 🛒 Ecom API

A RESTful e-commerce API built with Go, featuring product and order management backed by PostgreSQL.

## Tech Stack

- **Go** — [Chi](https://github.com/go-chi/chi) router with middleware
- **PostgreSQL 16** — via [pgx](https://github.com/jackc/pgx) driver
- **sqlc** — Type-safe SQL query generation
- **Goose** — Database migrations
- **Docker Compose** — Local development environment

## Project Structure

```
├── cmd/                    # Application entrypoint & server setup
├── internal/
│   ├── adapters/
│   │   └── postgresql/
│   │       ├── migrations/ # Goose SQL migrations
│   │       └── sqlc/       # Generated type-safe query code
│   ├── env/                # Environment config loader
│   ├── json/               # JSON response helpers
│   └── products/           # Product domain (handler + service)
├── docker-compose.yaml
├── sqlc.yaml
└── .ENV
```

## Getting Started

### Prerequisites

- [Go 1.25+](https://go.dev/dl/)
- [Docker](https://docs.docker.com/get-docker/)
- [Goose](https://github.com/pressly/goose) (for migrations)
- [sqlc](https://sqlc.dev/) (for code generation)

### 1. Start the database

```bash
docker compose up -d
```

This starts PostgreSQL on **port 5433** with database `ecom`.

### 2. Run migrations

```bash
goose -dir internal/adapters/postgresql/migrations postgres \
  "host=127.0.0.1 port=5433 user=postgres password=postgres dbname=ecom sslmode=disable" up
```

### 3. Run the server

```bash
go run ./cmd/
```

The API starts at **http://localhost:8080**.

## API Endpoints

| Method | Endpoint          | Description         |
|--------|-------------------|---------------------|
| GET    | `/health`         | Health check        |
| GET    | `/products`       | List all products   |
| GET    | `/products/{id}`  | Get product by UUID |

## Database Schema

### Products

| Column       | Type           | Details                     |
|--------------|----------------|-----------------------------|
| id           | UUID           | Primary key, auto-generated |
| name         | TEXT           | Required                    |
| description  | TEXT           | Optional                    |
| price        | NUMERIC(10,2)  | Required                    |
| created_at   | TIMESTAMPTZ    | Immutable, set on creation  |
| updated_at   | TIMESTAMPTZ    | Auto-updated on changes     |

### Orders

| Column       | Type           | Details                          |
|--------------|----------------|----------------------------------|
| id           | UUID           | Primary key, auto-generated      |
| customer_name| TEXT           | Required                         |
| status       | TEXT           | Defaults to `pending`            |
| total_amount | NUMERIC(10,2)  | Order total                      |
| created_at   | TIMESTAMPTZ    | Immutable, set on creation       |
| updated_at   | TIMESTAMPTZ    | Auto-updated on changes          |

### Order Items

| Column     | Type           | Details                    |
|------------|----------------|----------------------------|
| id         | UUID           | Primary key, auto-generated|
| order_id   | UUID           | FK → orders(id), cascades  |
| product_id | UUID           | FK → products(id)          |
| quantity   | INT            | Must be > 0                |
| unit_price | NUMERIC(10,2)  | Price at time of order     |
| created_at | TIMESTAMPTZ    | Set on creation            |

## Development

### Regenerate sqlc code

```bash
sqlc generate
```

### Create a new migration

```bash
goose -s create <migration_name> sql
```

## License

MIT
