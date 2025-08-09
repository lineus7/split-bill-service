DROP TABLE IF EXISTS transaction_statuses;
CREATE TABLE transaction_statuses (
	id SERIAL PRIMARY KEY,
	status VARCHAR(255) NOT NULL UNIQUE
);
