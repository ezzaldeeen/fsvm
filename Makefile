PROJECT_NAME := fsvm
GO := go

.PHONY: all build clean run test

all: build

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
