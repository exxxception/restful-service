run:
	go run cmd/main.go

migrate:
	goose -dir ./migrations postgres "postgresql://postgres:0000@localhost:5432/postgres?sslmode=disable" up

swag:
	swag init -g cmd/main.go