UNAME_S := $(shell uname -s)
BINARY_NAME=graphdbcli
ifeq ($(UNAME_S),Darwin)
	INSTALL_PATH=/usr/local/bin/$(BINARY_NAME)
else ifeq ($(UNAME_S),Linux)
	INSTALL_PATH=/usr/local/bin/$(BINARY_NAME)
else
	INSTALL_PATH=./$(BINARY_NAME)
endif


.PHONY: all build build-linux build-macos test fmt lint clean install



all: build


build:
	go build -o $(BINARY_NAME)

build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux-amd64

build-macos-intel:
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME)-darwin-amd64

build-macos-arm64:
	GOOS=darwin GOARCH=arm64 go build -o $(BINARY_NAME)-darwin-arm64

install:
	@if [ "$(UNAME_S)" = "Darwin" ] || [ "$(UNAME_S)" = "Linux" ]; then \
		echo "Installing $(BINARY_NAME) to $(INSTALL_PATH)"; \
		sudo cp $(BINARY_NAME) $(INSTALL_PATH); \
	else \
		echo "Unsupported OS. Please install manually."; \
	fi

fmt:
	go fmt ./...

lint:
	golangci-lint run

clean:
	rm -f $(BIN_FILE)
	@if [ -f /usr/local/bin/$(BINARY_NAME) ]; then sudo rm -f /usr/local/bin/$(BINARY_NAME); fi