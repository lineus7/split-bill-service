DROP TABLE IF EXISTS transactions;
CREATE TABLE transactions (
	id SERIAL PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users(id),
	status enum('pending', 'completed') NOT NULL,
	service_charge DECIMAL(10, 2) NOT NULL DEFAULT 0,
    tax_charge DECIMAL(10, 2) NOT NULL DEFAULT 0,
	total_price DECIMAL(10, 2) NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE INDEX idx_transactions_deleted_at ON transactions (deleted_at) WHERE deleted_at IS NOT NULL;