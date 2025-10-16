.PHONY: run build clean test help

# Default target
help:
	@echo "Bedtime Calculator - Available commands:"
	@echo "  make run    - Run the server (default port 8080)"
	@echo "  make build  - Build the binary"
	@echo "  make clean  - Remove built binaries"
	@echo "  make test   - Test the server endpoints"
	@echo ""
	@echo "Environment variables:"
	@echo "  PORT        - Server port (default: 8080)"

# Run the server
run:
	@echo "🛏️  Starting Bedtime Calculator server..."
	@go run main.go

# Build the binary
build:
	@echo "🔨 Building bedtime-calculator..."
	@go build -o bedtime-calculator main.go
	@echo "✅ Build complete! Run with: ./bedtime-calculator"

# Clean built binaries
clean:
	@echo "🧹 Cleaning..."
	@rm -f bedtime-calculator
	@echo "✅ Clean complete!"

# Test the server
test:
	@echo "🧪 Testing server endpoints..."
	@echo "Health check:"
	@curl -s http://localhost:8080/health || echo "Server not running. Start with 'make run'"
	@echo ""
	@echo "Root endpoint:"
	@curl -s http://localhost:8080/ | head -5 || echo "Server not running. Start with 'make run'"

