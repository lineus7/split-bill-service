DROP TABLE IF EXISTS transaction_item_users;
CREATE TABLE transaction_item_users (
	id SERIAL PRIMARY KEY,
    transaction_item_id INT NOT NULL REFERENCES transaction_items(id),
    user_id INT REFERENCES users(id),
    alternative_customer_name VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE INDEX idx_transaction_item_users_transaction_item_id ON transaction_item_users (transaction_item_id);
