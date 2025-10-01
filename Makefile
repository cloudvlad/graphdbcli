UNAME_S := $(shell uname -s)
BINARY_NAME=graphdbcli
ifeq ($(UNAME_S),Darwin)
	INSTALL_PATH=/usr/local/bin/$(BINARY_NAME)
else ifeq ($(UNAME_S),Linux)
	INSTALL_PATH=/usr/local/bin/$(BINARY_NAME)
else
	INSTALL_PATH=./$(BINARY_NAME)
endif


.PHONY: all package test fmt lint clean install build



all: build


package:
	@if [ -z "$(GOOS)" ] || [ -z "$(GOARCH)" ]; then \
        if [ -n "$(VERSION)" ]; then \
            go build -ldflags "-X graphdbcli/cmd.Version=$(VERSION)" -o $(BINARY_NAME); \
        else \
            go build -o $(BINARY_NAME); \
        fi; \
    else \
        if [ -n "$(VERSION)" ]; then \
            GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-X graphdbcli/cmd.Version=$(VERSION)" -o $(BINARY_NAME); \
        else \
            GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BINARY_NAME); \
        fi; \
        chmod +x $(BINARY_NAME); \
        TAR_NAME=$(BINARY_NAME)_$(GOOS)_$(GOARCH).tar.gz; \
        tar -czvf $$TAR_NAME $(BINARY_NAME) LICENSE README.md CHANGELOG.md; \
    fi

install:
	@if [ "$(UNAME_S)" = "Darwin" ] || [ "$(UNAME_S)" = "Linux" ]; then \
		echo "Installing $(BINARY_NAME) to $(INSTALL_PATH)"; \
		sudo cp $(BINARY_NAME) $(INSTALL_PATH); \
	else \
		echo "Unsupported OS. Please install manually."; \
	fi

build:
	go build -o $(BINARY_NAME)

fmt:
	go fmt ./...

lint:
	golangci-lint run

clean:
	rm -f $(BIN_FILE)
	@if [ -f /usr/local/bin/$(BINARY_NAME) ]; then sudo rm -f /usr/local/bin/$(BINARY_NAME); fi

test:
	go test ./cmd
