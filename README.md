#Migrate Development
migrate create -ext sql -dir database/migrations -seq create_transactions_table

migrate -source file://database/migrations -database "postgres://postgres:postgres@localhost:5432/split-bill?sslmode=disable" up
