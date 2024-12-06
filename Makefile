build:
	go build -o literalcheck ./cmd/literalcheck/main.go
test:
	go test ./...