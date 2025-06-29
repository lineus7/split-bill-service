#Migrate Development
migrate -source file://database/migrations -database "postgres://postgres:postgres@localhost:5432/split-bill?sslmode=disable" up
