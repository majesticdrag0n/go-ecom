# 🛒 Go Ecom API

A RESTful e-commerce API built with Go, featuring **products**, **orders** (with transactional rollback), and **customers** management backed by PostgreSQL.

## Tech Stack

- **Go** — [Chi](https://github.com/go-chi/chi) router with middleware
- **PostgreSQL 16** — via [pgx](https://github.com/jackc/pgx) driver
- **sqlc** — Type-safe SQL query generation
- **Goose** — Database migrations
- **Docker Compose** — Local development environment

## Project Structure

```
├── cmd/                        # Application entrypoint & server setup
│   ├── main.go                 # Database connection & app bootstrap
│   └── api.go                  # Router, middleware & route registration
├── internal/
│   ├── adapters/
│   │   └── postgresql/
│   │       ├── migrations/     # Goose SQL migrations
│   │       ├── sqlc/           # Generated type-safe query code
│   │       └── seed.sql        # Sample data for development
│   ├── customers/              # Customer domain (handler + service)
│   ├── orders/                 # Order domain (handler + service + types)
│   ├── products/               # Product domain (handler + service)
│   ├── env/                    # Environment config loader
│   └── json/                   # JSON response helpers
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
goose up
```

### 3. Seed sample data (optional)

```bash
# On Windows (PowerShell)
Get-Content internal/adapters/postgresql/seed.sql | docker exec -i <container_id> psql -U postgres -d ecom

# On Linux/Mac
psql "host=127.0.0.1 port=5433 user=postgres password=postgres dbname=ecom" -f internal/adapters/postgresql/seed.sql
```

This inserts 10 sample products and 5 sample customers.

### 4. Run the server

```bash
go run ./cmd/
```

The API starts at **http://localhost:8080**.

## API Endpoints

### Health

| Method | Endpoint  | Description  |
|--------|-----------|--------------|
| GET    | `/health` | Health check |

### Products

| Method | Endpoint         | Description         |
|--------|------------------|---------------------|
| GET    | `/products`      | List all products   |
| GET    | `/products/{id}` | Get product by UUID |

### Orders

| Method | Endpoint  | Description                                        |
|--------|-----------|----------------------------------------------------|
| POST   | `/orders` | Place an order (with items, transactional rollback) |

**Example — Place an order:**

```json
POST /orders
{
  "customer_name": "Ravi Kumar",
  "total_amount": 1539.98,
  "items": [
    { "product_id": "<uuid>", "quantity": 1, "unit_price": 1499.99 },
    { "product_id": "<uuid>", "quantity": 1, "unit_price": 39.99 }
  ]
}
```

> Order placement runs inside a **database transaction** — if any item fails (e.g. insufficient stock), the entire order is rolled back.

### Customers

| Method | Endpoint            | Description         |
|--------|---------------------|---------------------|
| POST   | `/customers`        | Create a customer   |
| GET    | `/customers`        | List all customers  |
| GET    | `/customers/{id}`   | Get customer by ID  |
| PUT    | `/customers/{id}`   | Update a customer   |
| DELETE | `/customers/{id}`   | Delete a customer   |

**Example — Create a customer:**

```json
POST /customers
{
  "name": "Ravi Kumar",
  "email": "ravi@email.com",
  "phone": "+91-9876543210",
  "address": "12 MG Road, Bangalore"
}
```

## Database Schema

### Products

| Column      | Type          | Details                       |
|-------------|---------------|-------------------------------|
| id          | UUID          | Primary key, auto-generated   |
| name        | TEXT          | Required                      |
| description | TEXT          | Optional                      |
| price       | NUMERIC(10,2) | Required                     |
| quantity    | INT           | Stock count, defaults to 0    |
| created_at  | TIMESTAMPTZ   | Immutable, set on creation    |
| updated_at  | TIMESTAMPTZ   | Auto-updated on changes       |

### Customers

| Column     | Type          | Details                       |
|------------|---------------|-------------------------------|
| id         | UUID          | Primary key, auto-generated   |
| name       | TEXT          | Required                      |
| email      | TEXT          | Required, unique              |
| phone      | TEXT          | Optional                      |
| address    | TEXT          | Optional                      |
| created_at | TIMESTAMPTZ   | Immutable, set on creation    |
| updated_at | TIMESTAMPTZ   | Auto-updated on changes       |

### Orders

| Column        | Type          | Details                     |
|---------------|---------------|-----------------------------|
| id            | UUID          | Primary key, auto-generated |
| customer_name | TEXT          | Required                    |
| customer_id   | UUID          | FK → customers(id)          |
| status        | TEXT          | Defaults to `pending`       |
| total_amount  | NUMERIC(10,2) | Order total                |
| created_at    | TIMESTAMPTZ   | Immutable, set on creation  |
| updated_at    | TIMESTAMPTZ   | Auto-updated on changes     |

### Order Items

| Column     | Type          | Details                     |
|------------|---------------|-----------------------------|
| id         | UUID          | Primary key, auto-generated |
| order_id   | UUID          | FK → orders(id), cascades   |
| product_id | UUID          | FK → products(id)           |
| quantity   | INT           | Must be > 0                 |
| unit_price | NUMERIC(10,2) | Price at time of order      |
| created_at | TIMESTAMPTZ   | Set on creation             |

## Key Features

- **Transactional Orders** — Order placement creates the order, adds items, and deducts product stock in a single transaction with automatic rollback on failure
- **Stock Protection** — `UpdateProductQuantity` prevents overselling with a `quantity >= requested` guard
- **UUID Primary Keys** — All entities use `gen_random_uuid()` for globally unique IDs
- **Immutable Timestamps** — Database triggers protect `created_at` from modification
- **Type-safe SQL** — All queries generated by sqlc with full Go type safety

## Development

### Regenerate sqlc code

```bash
sqlc generate
```

### Create a new migration

```bash
goose create <migration_name> sql
```

## License

MIT
