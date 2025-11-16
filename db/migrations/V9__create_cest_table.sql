CREATE TABLE cest (
    id SERIAL PRIMARY KEY,
    cest_code VARCHAR(255) NOT NULL,
    cest_item VARCHAR(255) NOT NULL,
    cest_name VARCHAR(1000) NOT NULL,
    cest_segment VARCHAR(255) NOT NULL
);