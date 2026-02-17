-- name: ListProducts :many
SELECT * FROM products
ORDER BY name;

-- name: GetProduct :one
SELECT * FROM products
WHERE id = $1;

-- name: PlaceOrder :one
INSERT INTO orders (customer_name, total_amount)
VALUES ($1, $2)
RETURNING *;

-- name: AddOrderItem :one
INSERT INTO order_items (order_id, product_id, quantity, unit_price)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateProductQuantity :one
UPDATE products
SET quantity = quantity - $2
WHERE id = $1 AND quantity >= $2
RETURNING *;

-- name: CreateCustomer :one
INSERT INTO customers (name, email, phone, address)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetCustomer :one
SELECT * FROM customers
WHERE id = $1;

-- name: ListCustomers :many
SELECT * FROM customers
ORDER BY name;

-- name: UpdateCustomer :one
UPDATE customers
SET name = $2, email = $3, phone = $4, address = $5
WHERE id = $1
RETURNING *;

-- name: DeleteCustomer :exec
DELETE FROM customers
WHERE id = $1;