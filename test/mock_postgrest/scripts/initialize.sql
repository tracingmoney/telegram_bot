CREATE SCHEMA IF NOT EXISTS tracing_money;

CREATE TABLE IF NOT EXISTS tracing_money.expense (
    record_id SERIAL PRIMARY KEY,
    piece_count INT NOT NULL,
    price FLOAT4 NOT NULL,
    description VARCHAR(255),
    date TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS tracing_money.income (
    record_id SERIAL PRIMARY KEY,
    income FLOAT4 NOT NULL,
    description VARCHAR(255),
    date TIMESTAMP NOT NULL
);