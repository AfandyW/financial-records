CREATE DATABASE financial;
CREATE TABLE transactions (
    id VARCHAR(50) PRIMARY KEY,
    title VARCHAR(100),
    description VARCHAR(200),
    type_transaction VARCHAR(10),
    amount integer,
    currency VARCHAR(50),
    category VARCHAR(50),
    sub_category VARCHAR(100),
    transaction_at TIMESTAMP NOT NULL,
    create_at TIMESTAMP NOT NULL,
    update_at TIMESTAMP NOT NULL
);