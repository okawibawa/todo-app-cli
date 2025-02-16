.PHONY: build run-bin run-dev

# Commands
build:
	@echo "Building binary..."
	@go build -o  ./tasks -v main.go && echo "Binary built successfully!" || { echo "Failed building binary!"; exit 1; }

clean:
	@echo "Cleaning binary..."
	@rm -rf ./tasks && echo "Binary cleaned!"

