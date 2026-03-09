BINARY_NAME := kubectl-debug_pvc
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS := -ldflags "-X main.version=$(VERSION)"

.PHONY: build install clean lint check

## Build the binary
build:
	go build $(LDFLAGS) -o $(BINARY_NAME) .

## Install to /usr/local/bin (as kubectl plugin)
install: build
	install -m 755 $(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)

## Install to user's go bin
install-gobin: build
	install -m 755 $(BINARY_NAME) $(shell go env GOPATH)/bin/$(BINARY_NAME)

## Remove the binary
clean:
	rm -f $(BINARY_NAME)

## Run go fmt
fmt:
	go fmt ./...

## Run go vet
vet:
	go vet ./...

## Run golangci-lint
lint:
	golangci-lint run ./...

## Run all checks (vet + lint)
check: vet lint

## Download dependencies
deps:
	go mod tidy

## Show help
help:
	@echo "kubectl debug-pvc - Debug pods with PVC volume access"
	@echo ""
	@echo "Targets:"
	@grep -E '^## ' Makefile | sed 's/## /  /'
	@echo ""
	@echo "Usage:"
	@echo "  make build     - Build the binary"
	@echo "  make install   - Build and install to /usr/local/bin"
	@echo "  make clean     - Remove built binary"
