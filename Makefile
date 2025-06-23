# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=auth-service
BINARY_UNIX=$(BINARY_NAME)_unix

# Development commands
.PHONY: run build clean test test-coverage test-race test-short deps mocks proto

# Build the application
build:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/app

# Build for Linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v ./cmd/app

# Run the application
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/app
	./$(BINARY_NAME)

# Clean build artifacts
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

# Run tests
test:
	$(GOTEST) -v ./...

# Run tests with coverage
test-coverage:
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

# Run tests with race detection
test-race:
	$(GOTEST) -v -race ./...

# Run tests (skip integration tests)
test-short:
	$(GOTEST) -v -short ./...

# Install dependencies
deps:
	$(GOMOD) download
	$(GOMOD) tidy

# Generate mocks
mocks:
	mockery --all

# Generate mocks from config
mocks-config:
	mockery

# Generate protobuf files
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		internal/protobuf/proto/*.proto

# Format code
fmt:
	$(GOCMD) fmt ./...

# Lint code (requires golangci-lint)
lint:
	golangci-lint run

# Install development tools
install-tools:
	$(GOCMD) install github.com/vektra/mockery/v2@latest
	$(GOCMD) install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	$(GOCMD) install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Docker commands
docker-build:
	docker build -t $(BINARY_NAME) .

docker-run:
	docker run -p 50051:50051 $(BINARY_NAME)

# Help
help:
	@echo "Available commands:"
	@echo "  build          - Build the application"
	@echo "  build-linux    - Build for Linux"
	@echo "  run            - Build and run the application"
	@echo "  clean          - Clean build artifacts"
	@echo "  test           - Run all tests"
	@echo "  test-coverage  - Run tests with coverage report"
	@echo "  test-race      - Run tests with race detection"
	@echo "  test-short     - Run tests (skip integration tests)"
	@echo "  deps           - Install dependencies"
	@echo "  mocks          - Generate all mocks"
	@echo "  mocks-config   - Generate mocks from config"
	@echo "  proto          - Generate protobuf files"
	@echo "  fmt            - Format code"
	@echo "  lint           - Lint code"
	@echo "  install-tools  - Install development tools"
	@echo "  docker-build   - Build Docker image"
	@echo "  docker-run     - Run Docker container"
	@echo "  help           - Show this help message"
