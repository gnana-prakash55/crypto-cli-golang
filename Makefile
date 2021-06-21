.PHONY: client

client:
	@echo "Building the client libraries"
	go build -o bin/client main.go