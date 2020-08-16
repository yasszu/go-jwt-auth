CREATE TABLE accounts (
	account_id serial PRIMARY KEY,
    email VARCHAR ( 255 ) UNIQUE NOT NULL,
	password VARCHAR ( 255 ) NOT NULL,
	created_on TIMESTAMP NOT NULL DEFAULT current_timestamp
);
