run:
	@go run ./cmd/rental-app/main.go

build:
	@go build -o bin/rental-app ./cmd/rental-app/main.go