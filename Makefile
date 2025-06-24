include .env.example
export

LOCAL_BIN:=$(CURDIR)/bin
PATH:=$(LOCAL_BIN):$(PATH)
BASE_STACK = docker compose -f docker-compose.yml

# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
.PHONY: help

run:
	go mod download && \
	CGO_ENABLED=0 go run ./cmd
.PHONY: run

compose-up: ### Run docker compose (without backend and reverse proxy)
	$(BASE_STACK) up --build -d redis redisinsight
.PHONY: compose-up

compose-down: ### Down docker compose
	$(BASE_STACK) down --remove-orphans
.PHONY: compose-down

deps: ### deps tidy + verify
	go mod tidy && go mod verify
.PHONY: deps

format: ### Run code formatter
	gofumpt -l -w .
	gci write . --skip-generated -s standard -s default
.PHONY: format

linter-golangci: ### check by golangci linter
	golangci-lint run
.PHONY: linter-golangci

linter-hadolint: ### check by hadolint linter
	git ls-files --exclude='Dockerfile*' -i -o | xargs hadolint
.PHONY: linter-hadolint

linter-dotenv: ### check by dotenv linter
	dotenv-linter
.PHONY: linter-dotenv
