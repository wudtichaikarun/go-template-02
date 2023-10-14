# Go variables
GO := go
GOFLAGS := -v

# Name of your Go binary (change this)
BINARY_NAME := bin/app

# Name of the Swag binary (change this if needed)
SWAG := swag

APP_PATH := ./cmd/app/main.go

# Run the dev
dev:
	@echo "💻 Running the service"
	@air $(APP_PATH)

# Run test dev
test:
	@echo "💻 Running unit test"
	@go test ./...

# Build the binary
build:
	@echo "📦️ Building..."
	@$(GO) build $(GOFLAGS) -o $(BINARY_NAME) $(APP_PATH)

# Clean the project
clean:
	@echo "🗑️ Cleaning..."
	@$(GO) clean
	@rm -f $(BINARY_NAME)

# Generate Swagger documentation
generate-swag-doc:
	@echo "📚 Generating Swagger documentation"
	@$(SWAG) init -g $(APP_PATH)

# Run the binary
run:
	@echo "🚀 Running the binary"
	@./$(BINARY_NAME)

.PHONY: dev build clean generate-swag-doc run