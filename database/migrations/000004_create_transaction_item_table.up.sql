DROP TABLE IF EXISTS transaction_items;
CREATE TABLE transaction_items (
	id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
	transaction_id INT NOT NULL REFERENCES transactions(id),
    alternative_customer_name VARCHAR(255),
    item_name VARCHAR(255) NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    add_on JSON NOT NULL DEFAULT '[]',
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP DEFAULT NULL
);

CREATE INDEX idx_transaction_items_deleted_at ON transaction_items (deleted_at) WHERE deleted_at IS NOT NULL;