BINARY           ?= bin
CONTAINER_ENGINE ?= podman
GOPATH           := $(shell go env GOPATH)
GOBIN            := $(shell go env GOBIN)

ifeq ($(GOBIN),)
	INSTALL_DIR=$(GOPATH)/bin
else
	INSTALL_DIR=$(GOBIN)
endif

.PHONY: clean install uninstall lint

$(BINARY): $(shell find . -type f -name '*.go')
	go build -o $(BINARY) .

clean:
	rm -f $(BINARY)
	go clean

install:
	go install .

uninstall:
	rm -f $(INSTALL_DIR)/$(BINARY)

lint:
	$(CONTAINER_ENGINE) run --rm -v "$(PWD):/app:Z" -w /app \
		docker.io/golangci/golangci-lint:v2.14.0 \
		golangci-lint run
