CREATE TABLE Accounts (
	account_id serial PRIMARY KEY,
	username VARCHAR ( 255 ) NOT NULL,
    email VARCHAR ( 255 ) NOT NULL,
	password VARCHAR ( 255 ) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
	updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
	UNIQUE(email)
);
