.PHONY: all build clean test

# Build directory
BUILD_DIR := build
SRC_DIR := src

# Find all Crystal source files in src directory
EXAMPLE_SRCS := $(wildcard $(SRC_DIR)/*.cr)
# Convert src/example.cr to build/example
EXAMPLE_BINS := $(patsubst $(SRC_DIR)/%.cr,$(BUILD_DIR)/%,$(EXAMPLE_SRCS))

# Default target
all: build

# Build all examples
build: $(EXAMPLE_BINS)
	@echo "Built $(words $(EXAMPLE_BINS)) examples to $(BUILD_DIR)/"

# Rule to build individual examples
$(BUILD_DIR)/%: $(SRC_DIR)/%.cr
	@mkdir -p $(BUILD_DIR)
	@echo "Building $@..."
	@crystal build --release -o $@ $<

# Clean build directory
clean:
	@rm -rf $(BUILD_DIR)
	@echo "Cleaned $(BUILD_DIR)/"

# Run tests
test:
	@crystal spec

# Help target
help:
	@echo "Available targets:"
	@echo "  make build    - Build all examples to build/ directory"
	@echo "  make clean    - Clean build directory"
	@echo "  make test     - Run tests"
	@echo "  make all      - Same as make build"
	@echo "  make help     - Show this help"