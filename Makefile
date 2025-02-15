.PHONY: build run-bin run-dev

# Commands
build:
	@echo "Building binary..."
	@go build -o  build/ -v main.go && echo "Binary built successfully!" || { echo "Failed building binary!"; exit 1; }

clean:
	@echo "Cleaning binary..."
	@rm -rf build/

run-bin:
	@echo "Running binary..."
	@./build/main

run-dev:
	@echo "Running development server..."
	@go run main.go || { echo "Failed running development server"; exit 1; }
