CREATE TABLE accounts
(
    id         SERIAL PRIMARY KEY,
    username   VARCHAR(255) NOT NULL,
    email      VARCHAR(255) NOT NULL,
    password   VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT current_timestamp,
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT current_timestamp,
    deleted_at TIMESTAMPTZ,
    UNIQUE (email)
);

CREATE INDEX ON accounts (deleted_at)