#help target to list all available commands
help:
	@echo "Available Commands:"
	@echo " run :   Run application"
	@echo " swag-init :   Run generated"
	@echo " test : Run all tests"
	@echo " test-v : Run all tests with verbose output"
	@echo " test-c : Run all tests with coverage"
	@echo " test-cv : Run all tests with coverage, verbose output, and detailed coverage profile"
	@echo " test-html : Generate HTML coverage report"
	@echo " test-all : Run all tests with coverage, detailed output, and generate HTML report"
	@echo " coverage : Display coverage function details"

run:
	go run cmd/main.go


# Run all tests
test:
	go test ./...

# Run all tests with verbose output
test-v:
	go test ./... -v

# Run all tests with coverage
test-c:
	go test -cover ./...

# Run all tests with coverage, verbose output, and detailed coverage profile
test-cv:
	go test ./... -cover -v -covermode=count -coverprofile=cover.out

# Generate HTML coverage report
test-html:
	go tool cover -html=cover.out -o cover.html

# Run all tests with coverage, detailed output, and generate HTML report
test-all:
	go test -cover ./... -v | grep statements && \
	go test -v -coverprofile cover.out ./... && \
	go tool cover -html=cover.out -o cover.html 

# Display coverage function details
coverage:
	go tool cover -func cover.out
