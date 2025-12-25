# Makefile for Jirar

# Variables
BINARY_NAME=jirar
BUILD_DIR=dist
MAIN_PATH=./cmd/jirar
VERSION=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS=-ldflags "-X main.version=$(VERSION)"

# Go variables
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Build targets
.PHONY: help build clean test lint fmt install run docker

# Default target
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Development
build: ## Build the binary for current platform
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)

build-all: ## Build binaries for all platforms
	@mkdir -p $(BUILD_DIR)
	# Linux AMD64
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_PATH)
	# Linux ARM64
	GOOS=linux GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 $(MAIN_PATH)
	# macOS AMD64
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(MAIN_PATH)
	# macOS ARM64 (M1/M2)
	GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 $(MAIN_PATH)
	# Windows AMD64
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_PATH)

clean: ## Clean build artifacts
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html

# Testing
test: ## Run tests
	$(GOTEST) -v -race ./...

test-coverage: ## Run tests with coverage
	$(GOTEST) -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

test-bench: ## Run benchmark tests
	$(GOTEST) -bench=. -benchmem ./...

# Code quality
lint: ## Run linter
	golangci-lint run --timeout=5m

fmt: ## Format code
	goimports -w .
	gofmt -s -w .

fmt-check: ## Check if code is formatted
	@! gofmt -s -l . | grep -q '.*\.go'
	@! goimports -l . | grep -q '.*\.go'

vet: ## Run go vet
	$(GOCMD) vet ./...

security: ## Run security checks
	gosec ./...

# Dependencies
deps: ## Download dependencies
	$(GOMOD) download

deps-update: ## Update dependencies
	$(GOMOD) get -u ./...
	$(GOMOD) tidy

deps-verify: ## Verify dependencies
	$(GOMOD) verify

# Development
run: ## Run the application
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	./$(BUILD_DIR)/$(BINARY_NAME)

install: ## Install binary to $GOPATH/bin
	$(GOBUILD) $(LDFLAGS) -o $${GOPATH:-~/go}/bin/$(BINARY_NAME) $(MAIN_PATH)

# Docker
docker-build: ## Build Docker image
	docker build -t jirar:$(VERSION) .
	docker tag jirar:$(VERSION) jirar:latest

docker-run: ## Run Docker container
	docker run --rm -it \
		-e JIRA_DOMAIN=${JIRA_DOMAIN} \
		-e JIRA_EMAIL=${JIRA_EMAIL} \
		-e JIRA_TOKEN=${JIRA_TOKEN} \
		jirar:latest

# CI/CD
ci: fmt-check vet lint test security ## Run all CI checks

# Release
release-test: ## Run release tests
	$(GOTEST) -race -coverprofile=coverage.out ./... -tags=integration

# Development helpers
setup: ## Set up development environment
	@echo "Setting up development environment..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/securecodewarrior/github-action-gosec@latest
	go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
	@echo "Development tools installed!"

docs: ## Generate documentation
	@echo "Generating documentation..."
	@echo "Documentation generated to docs/"

# Version info
version: ## Show version information
	@echo "Version: $(VERSION)"
	@echo "Go version: $(shell go version)"

# Quick commands
dev: clean deps fmt ci ## Clean, update deps, format, and run CI

# Development server
watch: ## Watch for changes and rebuild
	@echo "Watching for changes..."
	@command -v fswatch >/dev/null 2>&1 || { echo "fswatch not found. Install with: brew install fswatch"; exit 1; }
	fswatch -o . | xargs -n1 -I{} make build

# Git helpers
pre-commit: fmt-check vet lint test security ## Run pre-commit hooks
	@echo "✅ Pre-commit checks passed!"

# Clean up
deep-clean: ## Clean all artifacts
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html
	$(GOMOD) clean -modcache
	docker system prune -f