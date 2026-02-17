-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    description TEXT,
    price NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Trigger function to prevent changes to created_at
CREATE OR REPLACE FUNCTION protect_created_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.created_at = OLD.created_at;  -- Always keep the original value
    NEW.updated_at = NOW();           -- Auto-update updated_at
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_protect_created_at
    BEFORE UPDATE ON products
    FOR EACH ROW
    EXECUTE FUNCTION protect_created_at();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS trg_protect_created_at ON products;
DROP FUNCTION IF EXISTS protect_created_at();
DROP TABLE IF EXISTS products;
-- +goose StatementEnd
