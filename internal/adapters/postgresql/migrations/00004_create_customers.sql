-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS customers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    phone TEXT,
    address TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER trg_protect_customers_created_at
    BEFORE UPDATE ON customers
    FOR EACH ROW
    EXECUTE FUNCTION protect_created_at();

-- Add customer_id FK to orders
ALTER TABLE orders ADD COLUMN customer_id UUID REFERENCES customers(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE orders DROP COLUMN IF EXISTS customer_id;
DROP TRIGGER IF EXISTS trg_protect_customers_created_at ON customers;
DROP TABLE IF EXISTS customers;
-- +goose StatementEnd
