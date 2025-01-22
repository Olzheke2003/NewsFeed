.PHONY: build
build:
	go build -v ./cmd/app

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build


.PHONY: migrate
migrate:
	migrate -database "postgres://postgres:admin@localhost:5432/news_feed?sslmode=disable" -path ./migrations up