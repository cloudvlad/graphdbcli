UNAME_S := $(shell uname -s)
BINARY_NAME=graphdbcli
ifeq ($(UNAME_S),Darwin)
	BINARY_EXT=
	INSTALL_PATH=/usr/local/bin/$(BINARY_NAME)
else ifeq ($(UNAME_S),Linux)
	BINARY_EXT=
	INSTALL_PATH=/usr/local/bin/$(BINARY_NAME)
else
	BINARY_EXT=
	INSTALL_PATH=./$(BINARY_NAME)
endif
BIN_FILE=$(BINARY_NAME)$(BINARY_EXT)

.PHONY: all build test fmt lint clean install


all: build

build:
	go build -o $(BIN_FILE)

install:
	@if [ "$(UNAME_S)" = "Darwin" ] || [ "$(UNAME_S)" = "Linux" ]; then \
		echo "Installing $(BIN_FILE) to $(INSTALL_PATH)"; \
		sudo cp $(BIN_FILE) $(INSTALL_PATH); \
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