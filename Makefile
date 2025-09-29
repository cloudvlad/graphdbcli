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
	@if [ -n "$(VERSION)" ]; then \
		go build -ldflags "-X graphdbcli/cmd.Version=$(VERSION)" -o $(BINARY_NAME); \
	else \
		go build -o $(BINARY_NAME); \
	fi

build-linux:
	@if [ -n "$(VERSION)" ]; then \
	  GOOS=linux GOARCH=amd64 go build -ldflags "-X graphdbcli/cmd.Version=$(VERSION)" -o $(BINARY_NAME)-linux-amd64; \
	else \
	  GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux-amd64; \
	fi

build-macos-intel:
	@if [ -n "$(VERSION)" ]; then \
	  GOOS=darwin GOARCH=amd64 go build -ldflags "-X graphdbcli/cmd.Version=$(VERSION)" -o $(BINARY_NAME)-darwin-amd64; \
	else \
	  GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME)-darwin-amd64; \
	fi

build-macos-arm64:
	@if [ -n "$(VERSION)" ]; then \
	  GOOS=darwin GOARCH=arm64 go build -ldflags "-X graphdbcli/cmd.Version=$(VERSION)" -o $(BINARY_NAME)-darwin-arm64; \
	else \
	  GOOS=darwin GOARCH=arm64 go build -o $(BINARY_NAME)-darwin-arm64; \
	fi

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

test:
	go test ./cmd
