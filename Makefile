generate:
	@echo "Generating..."
	go generate ./...


test:
	@echo "Testing..."
	go test -v ./...

build:
	@echo "Building..."
	go build -o bin/ ./...