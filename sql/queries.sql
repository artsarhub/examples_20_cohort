CREATE TABLE orders
(
    order_id      INTEGER PRIMARY KEY,
    customer_name TEXT,
    customer_city TEXT,
    product_name  TEXT,
    product_price REAL,
    order_date    DATE
);

INSERT INTO orders (order_id, customer_name, customer_city, product_name, product_price, order_date)
VALUES (1, 'Alice', 'New York', 'Laptop', 1000.00, '2024-11-01'),
       (2, 'Bob', 'Los Angeles', 'Smartphone', 500.00, '2024-11-02'),
       (3, 'Alice', 'New York', 'Tablet', 300.00, '2024-11-03'),
       (4, 'Charlie', 'Chicago', 'Laptop', 1000.00, '2024-11-04'),
       (5, 'Bob', 'Los Angeles', 'Tablet', 300.00, '2024-11-05');

CREATE TABLE customers
(
    customer_id INTEGER PRIMARY KEY,
    name       TEXT,
    city       TEXT
);

INSERT INTO customers (customer_id, name, city)
VALUES (1, 'Alice', 'New York'),
       (2, 'Bob', 'Los Angeles'),
       (3, 'Charlie', 'Chicago');

CREATE TABLE products
(
    product_id INTEGER PRIMARY KEY,
    name      TEXT,
    price     REAL
);

INSERT INTO products (product_id, name, price)
VALUES (1, 'Laptop', 1000.00),
       (2, 'Smartphone', 500.00),
       (3, 'Tablet', 300.00);

CREATE TABLE new_orders
(
    order_id    INTEGER PRIMARY KEY,
    customer_id INTEGER,
    product_id  INTEGER,
    order_date  DATE,
    FOREIGN KEY (customer_id) REFERENCES customers (customer_id),
    FOREIGN KEY (product_id) REFERENCES products (product_id)
);

INSERT INTO new_orders (order_id, customer_id, product_id, order_date)
VALUES (1, 1, 1, '2024-11-01'),
       (2, 2, 2, '2024-11-02'),
       (3, 1, 3, '2024-11-03'),
       (4, 3, 1, '2024-11-04'),
       (5, 2, 3, '2024-11-05');

DROP TABLE orders;

ALTER TABLE new_orders
    RENAME TO orders;

------------------------------------------------------------------------------------------

select * from orders;

SELECT COUNT(*) AS total_orders FROM orders;

SELECT o.order_id, c.name, p.name AS product_name, o.order_date
FROM orders o
         JOIN customers c ON o.customer_id = c.customer_id
         JOIN products p ON o.product_id = p.product_id
WHERE c.name = 'Alice';

SELECT c.name, COUNT(o.order_id) AS total_orders
FROM orders o
         JOIN customers c ON o.customer_id = c.customer_id
GROUP BY c.name;

with filtered_customers as (SELECT customer_id,
                    name
             from customers
             where name in ('Charlie', 'Bob'))
SELECT c.name, SUM(p.price) AS total_spent
FROM orders o
         JOIN filtered_customers c ON o.customer_id = c.customer_id
         JOIN products p ON o.product_id = p.product_id
GROUP BY c.name
order by total_spent desc
limit 1;

-- Индекс на CustomerID в таблице orders
CREATE INDEX idx_orders_customer_id ON orders (customer_id);

-- Индекс на ProductID в таблице orders
CREATE INDEX idx_orders_product_id ON orders (product_id);

-- Индекс на Name в таблице customers
CREATE INDEX idx_customers_name ON customers (name);

EXPLAIN QUERY PLAN
SELECT c.name, SUM(p.price) AS total_spent
FROM orders o
         JOIN customers c ON o.customer_id = c.customer_id
         JOIN products p ON o.product_id = p.product_id
GROUP BY c.name;

-- acid
