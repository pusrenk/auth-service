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
.PHONY: run build clean test test-coverage test-race test-short deps mocks proto gen fmt add-domain
.PHONY: docker-build docker-build-push docker-push docker-run docker-build-tag docker-push-tag
.PHONY: docker-build-auto docker-build-auto-full docker-push-auto docker-push-auto-full docker-build-push-auto docker-show-tags
.PHONY: docker-compose-build docker-compose-up docker-compose-up-env docker-compose-down docker-compose-down-env
.PHONY: docker-compose-dev docker-compose-dev-up docker-compose-staging-up docker-compose-prod-up

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
DOCKER_USERNAME ?= your-dockerhub-username
DOCKER_IMAGE = $(DOCKER_USERNAME)/$(BINARY_NAME)
DOCKER_TAG ?= latest
ENV ?= dev

# Auto-generate tags from Git info
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD 2>/dev/null || echo "main")
GIT_COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
GIT_TAG := $(shell git describe --tags --exact-match 2>/dev/null || echo "")

# Clean branch name for Docker (replace invalid characters)
CLEAN_BRANCH := $(shell echo $(GIT_BRANCH) | sed 's/[^a-zA-Z0-9._-]/-/g' | tr '[:upper:]' '[:lower:]')

# Auto-tag logic: use git tag if exists, otherwise use branch name
AUTO_TAG := $(if $(GIT_TAG),$(GIT_TAG),$(CLEAN_BRANCH))
FULL_TAG := $(AUTO_TAG)-$(GIT_COMMIT)

# Method 1: Using environment variables with defaults
docker-build:
	docker build -t $(BINARY_NAME):$(DOCKER_TAG) .
	docker tag $(BINARY_NAME):$(DOCKER_TAG) $(DOCKER_IMAGE):$(DOCKER_TAG)

docker-build-push: docker-build
	docker push $(DOCKER_IMAGE):$(DOCKER_TAG)

docker-push:
	docker push $(DOCKER_IMAGE):$(DOCKER_TAG)

docker-run:
	docker run -p 8080:8080 $(BINARY_NAME):$(DOCKER_TAG)

# Method 2: Enhanced docker-compose commands with environment support
docker-compose-build:
	@if [ -z "$(ENV)" ]; then \
		echo "Usage: make docker-compose-build ENV=<environment>"; \
		echo "Example: make docker-compose-build ENV=dev"; \
		echo "Available: dev, staging, prod"; \
		exit 1; \
	fi
	@echo "Building for environment: $(ENV)"
	docker-compose build --build-arg ENV=$(ENV)
	docker-compose -p auth-$(ENV) build

docker-compose-up:
	docker-compose -p auth-$(ENV) up --build

docker-compose-up-env:
	@if [ -z "$(ENV)" ]; then \
		echo "Usage: make docker-compose-up-env ENV=<environment>"; \
		echo "Example: make docker-compose-up-env ENV=dev"; \
		exit 1; \
	fi
	@echo "Starting services for environment: $(ENV)"
	docker-compose -p auth-$(ENV) up --build

docker-compose-down:
	docker-compose -p auth-$(ENV) down

docker-compose-down-env:
	@if [ -z "$(ENV)" ]; then \
		echo "Usage: make docker-compose-down-env ENV=<environment>"; \
		echo "Example: make docker-compose-down-env ENV=dev"; \
		exit 1; \
	fi
	@echo "Stopping services for environment: $(ENV)"
	docker-compose -p auth-$(ENV) down

docker-compose-dev:
	docker-compose -p auth-dev up --build -d

# Method 3: Specific environment shortcuts
docker-compose-dev-up:
	@echo "Starting DEV environment..."
	docker-compose -p auth-dev up --build

docker-compose-staging-up:
	@echo "Starting STAGING environment..."
	docker-compose -p auth-staging up --build

docker-compose-prod-up:
	@echo "Starting PROD environment..."
	docker-compose -p auth-prod up --build

# Method 4: Advanced - Build with custom tag and push
docker-build-tag:
	@if [ -z "$(TAG)" ]; then \
		echo "Usage: make docker-build-tag TAG=<tag_name>"; \
		echo "Example: make docker-build-tag TAG=v1.0.0"; \
		exit 1; \
	fi
	@echo "Building with tag: $(TAG)"
	docker build -t $(BINARY_NAME):$(TAG) .
	docker tag $(BINARY_NAME):$(TAG) $(DOCKER_IMAGE):$(TAG)

docker-push-tag:
	@if [ -z "$(TAG)" ]; then \
		echo "Usage: make docker-push-tag TAG=<tag_name>"; \
		echo "Example: make docker-push-tag TAG=v1.0.0"; \
		exit 1; \
	fi
	@echo "Pushing tag: $(TAG)"
	docker push $(DOCKER_IMAGE):$(TAG)

# Method 5: Auto-tag from Git branch/commit
docker-build-auto:
	@echo "Git Info:"
	@echo "  Branch: $(GIT_BRANCH)"
	@echo "  Commit: $(GIT_COMMIT)"
	@echo "  Git Tag: $(if $(GIT_TAG),$(GIT_TAG),none)"
	@echo "  Auto Tag: $(AUTO_TAG)"
	@echo "  Full Tag: $(FULL_TAG)"
	@echo ""
	@echo "Building with auto-generated tag: $(AUTO_TAG)"
	docker build -t $(BINARY_NAME):$(AUTO_TAG) .
	docker tag $(BINARY_NAME):$(AUTO_TAG) $(DOCKER_IMAGE):$(AUTO_TAG)

docker-build-auto-full:
	@echo "Building with full tag (includes commit): $(FULL_TAG)"
	docker build -t $(BINARY_NAME):$(FULL_TAG) .
	docker tag $(BINARY_NAME):$(FULL_TAG) $(DOCKER_IMAGE):$(FULL_TAG)

docker-push-auto:
	@echo "Pushing auto-generated tag: $(AUTO_TAG)"
	docker push $(DOCKER_IMAGE):$(AUTO_TAG)

docker-push-auto-full:
	@echo "Pushing full tag: $(FULL_TAG)"
	docker push $(DOCKER_IMAGE):$(FULL_TAG)

docker-build-push-auto: docker-build-auto docker-push-auto
	@echo "Build and push completed for tag: $(AUTO_TAG)"

# Method 6: Show available tags
docker-show-tags:
	@echo "Available Docker tags:"
	@echo "  Manual tag: $(DOCKER_TAG) (default: latest)"
	@echo "  Auto tag: $(AUTO_TAG) (from branch: $(GIT_BRANCH))"
	@echo "  Full tag: $(FULL_TAG) (includes commit: $(GIT_COMMIT))"
	@echo "  Git tag: $(if $(GIT_TAG),$(GIT_TAG),none)"
	@echo ""
	@echo "Usage examples:"
	@echo "  make docker-build                    # Uses DOCKER_TAG ($(DOCKER_TAG))"
	@echo "  make docker-build DOCKER_TAG=v1.0.0 # Uses manual tag"
	@echo "  make docker-build-auto               # Uses auto tag ($(AUTO_TAG))"
	@echo "  make docker-build-auto-full          # Uses full tag ($(FULL_TAG))"

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
	@echo "Docker Commands (Manual Tags):"
	@echo "  docker-build                        - Build Docker image with tag (DOCKER_TAG=latest)"
	@echo "  docker-build-push                   - Build and push Docker image to registry"
	@echo "  docker-push                         - Push Docker image to registry"
	@echo "  docker-run                          - Run Docker container with tag"
	@echo "  docker-build-tag TAG=<>             - Build with custom tag"
	@echo "  docker-push-tag TAG=<>              - Push with custom tag"
	@echo ""
	@echo "Docker Commands (Auto Tags from Git):"
	@echo "  docker-build-auto                   - Build with branch-based tag"
	@echo "  docker-build-auto-full              - Build with branch+commit tag"
	@echo "  docker-push-auto                    - Push branch-based tag"
	@echo "  docker-push-auto-full               - Push branch+commit tag"
	@echo "  docker-build-push-auto              - Build and push auto tag"
	@echo "  docker-show-tags                    - Show all available tag options"
	@echo ""
	@echo "Docker Compose Commands:"
	@echo "  docker-compose-build ENV=<>    - Build for specific environment"
	@echo "  docker-compose-up               - Start services (ENV=dev default)"
	@echo "  docker-compose-up-env ENV=<>    - Start services for environment"
	@echo "  docker-compose-down             - Stop services (ENV=dev default)"
	@echo "  docker-compose-down-env ENV=<>  - Stop services for environment"
	@echo "  docker-compose-dev              - Start dev services in background"
	@echo "  docker-compose-dev-up           - Start dev environment"
	@echo "  docker-compose-staging-up       - Start staging environment"
	@echo "  docker-compose-prod-up          - Start prod environment"
	@echo "  help                - Show this help message"
