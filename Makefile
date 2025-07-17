# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=auth-service
BINARY_UNIX=$(BINARY_NAME)_unix

DOCKER_USERNAME="rayprastya"
IMAGE_NAME="pusrenk"
TAG="auth-service"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Development commands
.PHONY: run build clean test test-coverage test-race test-short deps mocks proto gen fmt add-domain

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

# Generate protobuf files (from gen.sh)
proto:
	@echo "--- Generating protobuf files"
	cd internal/protobuf/proto && protoc -I . --go_out=.. --go-grpc_out=.. *.proto

# Generate protobuf files and format (equivalent to gen.sh)
gen: proto fmt
	@echo "--- Generation complete"

# Format code with goimports (from fmt.sh)
fmt:
	@if command -v goimports >/dev/null 2>&1; then \
		echo "Using goimports for formatting..."; \
		goimports -w -local github.com/pusrenk/auth-service .; \
	else \
		echo "goimports not found, using go fmt instead..."; \
		echo "Install goimports with: go install golang.org/x/tools/cmd/goimports@latest"; \
		$(GOCMD) fmt ./...; \
	fi

# Format code with go fmt
go-fmt:
	$(GOCMD) fmt ./...

# Lint code (requires golangci-lint)
lint:
	golangci-lint run

# Install development tools
install-tools:
	$(GOCMD) install github.com/vektra/mockery/v2@latest
	$(GOCMD) install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	$(GOCMD) install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	$(GOCMD) install golang.org/x/tools/cmd/goimports@latest

# Domain management - Add new domain
add-domain:
	@if [ -z "$(DOMAIN)" ]; then \
		echo "Usage: make add-domain DOMAIN=<domain_name>"; \
		echo "Example: make add-domain DOMAIN=product"; \
		exit 1; \
	fi
	@echo "Adding domain: $(DOMAIN)"
	@powershell -ExecutionPolicy Bypass -File scripts/add-domain.ps1 -DomainName $(DOMAIN)
	@echo "Don't forget to run 'make mocks-config' to generate the mocks!"

# Docker commands
docker-build:
	docker build -t $(BINARY_NAME) .

docker-run:
	docker run -p 50051:50051 $(BINARY_NAME)

docker-push:
	@echo -e "${YELLOW}Building Docker image...${NC}"
	@docker build -t ${DOCKER_USERNAME}/${IMAGE_NAME}:${TAG} . && \
		echo -e "${GREEN}Build successful!${NC}" && \
		echo -e "${YELLOW}Tagging image...${NC}" && \
		docker tag ${DOCKER_USERNAME}/${IMAGE_NAME}:${TAG} ${DOCKER_USERNAME}/${IMAGE_NAME}:${TAG} && \
		echo -e "${YELLOW}Pushing to Docker Hub...${NC}" && \
		docker push ${DOCKER_USERNAME}/${IMAGE_NAME}:${TAG} && \
		echo -e "${GREEN}Successfully pushed to Docker Hub!${NC}" && \
		echo -e "${GREEN}Image: ${DOCKER_USERNAME}/${IMAGE_NAME}:${TAG}${NC}" || \
		echo -e "${RED}Operation failed!${NC}" 



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
	@echo "  gen            - Generate protobuf files and format (replaces gen.sh)"
	@echo "  fmt            - Format code with goimports (replaces fmt.sh)"
	@echo "  go-fmt         - Format code with go fmt"
	@echo "  lint           - Lint code"
	@echo "  install-tools  - Install development tools"
	@echo "  add-domain     - Add new domain (Usage: make add-domain DOMAIN=product)"
	@echo "  docker-build   - Build Docker image"
	@echo "  docker-run     - Run Docker container"
	@echo "  help           - Show this help message"
