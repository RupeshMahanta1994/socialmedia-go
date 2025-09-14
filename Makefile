build:
	@go build -o bin/socialmedia-go cmd/main.go

run: build
	@./bin/socialmedia-go
