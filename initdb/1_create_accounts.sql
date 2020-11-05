CREATE TABLE accounts
(
    id            SERIAL PRIMARY KEY,
    username      VARCHAR(255) NOT NULL,
    email         VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at    TIMESTAMPTZ  NOT NULL DEFAULT current_timestamp,
    updated_at    TIMESTAMPTZ  NOT NULL DEFAULT current_timestamp,
    UNIQUE (email)
);
