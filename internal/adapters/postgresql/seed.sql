-- Seed data for the ecom database
-- Run with: psql "host=127.0.0.1 port=5433 user=postgres password=postgres dbname=ecom" -f internal/adapters/postgresql/seed.sql

INSERT INTO products (name, description, price, quantity) VALUES
  ('Gaming Laptop',     'High-performance laptop with RTX 4080',   1499.99, 25),
  ('Wireless Mouse',    'Ergonomic Bluetooth mouse',                 39.99, 150),
  ('Mechanical Keyboard', 'Cherry MX Brown switches, RGB backlit',  129.99, 75),
  ('USB-C Hub',         '7-in-1 multiport adapter',                  49.99, 200),
  ('27" 4K Monitor',    'IPS panel, 144Hz refresh rate',            599.99, 40),
  ('Webcam HD',         '1080p with built-in microphone',            79.99, 100),
  ('Noise-Cancelling Headphones', 'Over-ear, 30hr battery life',   249.99, 60),
  ('Portable SSD 1TB',  'USB 3.2, read speeds up to 1050MB/s',     109.99, 90),
  ('Laptop Stand',      'Adjustable aluminum stand',                 34.99, 120),
  ('Wireless Charger',  'Qi-compatible 15W fast charge pad',         24.99, 180)
ON CONFLICT DO NOTHING;

INSERT INTO customers (name, email, phone, address) VALUES
  ('Ravi Kumar',     'ravi.kumar@email.com',     '+91-9876543210', '12 MG Road, Bangalore'),
  ('Priya Sharma',   'priya.sharma@email.com',   '+91-9123456789', '45 Linking Road, Mumbai'),
  ('Amit Patel',     'amit.patel@email.com',     '+91-9988776655', '78 Ashram Road, Ahmedabad'),
  ('Sneha Reddy',    'sneha.reddy@email.com',    '+91-9112233445', '23 Jubilee Hills, Hyderabad'),
  ('Vikram Singh',   'vikram.singh@email.com',   '+91-9001122334', '56 Connaught Place, Delhi')
ON CONFLICT DO NOTHING;
