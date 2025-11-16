CREATE TABLE product_fiscal_type (
    id SERIAL PRIMARY KEY,
    product_code VARCHAR(100) NOT NULL,
    fiscal_name VARCHAR(500) NOT NULL
);
