# Makefile for push-swap project

# Variables
GO_CMD=go
BUILD_CMD=$(GO_CMD) build
CLEAN_CMD=$(GO_CMD) clean
TEST_CMD=$(GO_CMD) test
FMT_CMD=$(GO_CMD) fmt
VET_CMD=$(GO_CMD) vet

# Binary names
PUSH_SWAP_BIN=push-swap
CHECKER_BIN=checker

# Source directories
PUSH_SWAP_SRC=./cmd/push-swap
CHECKER_SRC=./cmd/checker

# Default target
all: build

# Build both binaries
build: $(PUSH_SWAP_BIN) $(CHECKER_BIN)

# Build push-swap binary
$(PUSH_SWAP_BIN):
	$(BUILD_CMD) -o $(PUSH_SWAP_BIN) $(PUSH_SWAP_SRC)

# Build checker binary
$(CHECKER_BIN):
	$(BUILD_CMD) -o $(CHECKER_BIN) $(CHECKER_SRC)

# Run tests
test:
	$(TEST_CMD) ./...

# Format code
fmt:
	$(FMT_CMD) ./...

# Vet code
vet:
	$(VET_CMD) ./...

# Clean binaries
clean:
	$(CLEAN_CMD)
	rm -f $(PUSH_SWAP_BIN) $(CHECKER_BIN)

# Run quality checks
check: fmt vet test

# Install dependencies (if any)
deps:
	$(GO_CMD) mod tidy
	$(GO_CMD) mod download

# Example usage targets
demo: build
	@echo "Running demo with example input..."
	@echo "Input: 4 67 3 87 23"
	@./$(PUSH_SWAP_BIN) "4 67 3 87 23"

validate: build
	@echo "Validating push-swap output with checker..."
	@ARG="4 67 3 87 23"; ./$(PUSH_SWAP_BIN) "$$ARG" | ./$(CHECKER_BIN) "$$ARG"

# Help target
help:
	@echo "Available targets:"
	@echo "  build     - Build both push-swap and checker binaries"
	@echo "  test      - Run all tests"
	@echo "  fmt       - Format all Go code"
	@echo "  vet       - Run go vet on all packages"
	@echo "  clean     - Remove built binaries"
	@echo "  check     - Run fmt, vet, and test"
	@echo "  deps      - Install/update dependencies"
	@echo "  demo      - Run a demo with example input"
	@echo "  validate  - Validate push-swap output with checker"
	@echo "  help      - Show this help message"

.PHONY: all build test fmt vet clean check deps demo validate help