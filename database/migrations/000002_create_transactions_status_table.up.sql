DROP TABLE IF EXISTS transactions_status;
CREATE TABLE transactions_status (
	id SERIAL PRIMARY KEY,
	status VARCHAR(255) NOT NULL UNIQUE
);
