### Go Variables
GOCMD=go
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet

### Output Variables
BINARY_NAME=lol
VERSION ?= 0.0.0
OUT_DIR := ./out/bin
TARGET := $(OUT_DIR)/$(BINARY_NAME)

### TPUT Variables
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: all test build vendor

all: help

## Build:
build: ## Build your project and put the output binary in out/bin/
	mkdir -p out/bin
	$(GOCMD) build -o $(TARGET)

.PHONEY: clean
clean: ## Remove build related file
	go clean
	rm ${TARGET}

vendor: ## Copy of all packages needed to support builds and tests in the vendor directory
	$(GOCMD) mod vendor

## Test:
test: ## Run the tests of the project
	$(GOTEST) -v -race ./... $(OUTPUT_OPTIONS)

## Lint:
lint: lint-go lint-reportcard ## Run all available linters

lint-go:
	go vet
	golint

lint-reportcard: ## Use goreportcard-cli on your project
	goreportcard-cli

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)
