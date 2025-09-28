DROP TABLE IF EXISTS transaction_item_add_ons;
CREATE TABLE transaction_item_add_ons (
	id SERIAL PRIMARY KEY,
    transaction_item_id INT NOT NULL REFERENCES transaction_items(id),
    item_name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP DEFAULT NULL
);

CREATE INDEX idx_transaction_item_add_ons_deleted_at ON transaction_item_add_ons (deleted_at) WHERE deleted_at IS NOT NULL;