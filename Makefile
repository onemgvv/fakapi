include .env

build:
	go build -o .bin/main cmd/main/app.go

prod: build
	./.bin/main production

dev:
	go run cmd/main/app.go development

db_up:
	migrate -path ./schema -database "postgresql://postgres:root@localhost:5432/fakapi?sslmode=disable" up

db_down:
	migrate -path ./schema -database "postgresql://postgres:root@localhost:5432/fakapi?sslmode=disable" down