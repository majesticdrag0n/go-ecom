-- +goose Up
ALTER TABLE products ADD COLUMN quantity INT NOT NULL DEFAULT 0;

-- +goose Down
ALTER TABLE products DROP COLUMN quantity;
