BINARY_DIR := bin
BINARY_NAME := okta-viewer

.PHONY: all build run test clean

all: build

build:
	go build -o $(BINARY_DIR)/$(BINARY_NAME) .

test:
	go test ./... -v

# Clean compiled artifacts
clean:
	rm -rf $(BINARY_DIR)
