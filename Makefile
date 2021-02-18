.PHONY:
.SILENT:

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

swag:
	swag init -g internal/app/app.go
