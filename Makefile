PROJECT_NAME := fsvm
GO := go

.PHONY: build clean run test

build:
	@echo "Building $(PROJECT_NAME)..."
	$(GO) build -o $(PROJECT_NAME)

clean:
	@echo "Cleaning $(PROJECT_NAME)..."
	@rm -f $(PROJECT_NAME)

run: build
	@echo "Running $(PROJECT_NAME)..."
	./$(PROJECT_NAME)

test:
	@echo "Running tests..."
	$(GO) test ./...

coverage:
	@echo "Running tests with coverage..."
	$(GO) test ./... --cover

install-deps:
	@echo "Installing dependencies..."
	$(GO) mod download