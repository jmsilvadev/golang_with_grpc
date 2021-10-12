.DEFAULT_GOAL := help
.PHONY: up rebuild down ssh test test.coverage build logs docker-cleanup help lint vet fmt deps vendor doc proto go-build cert test.api

DOCKER_C := docker-compose
DOCKER_RUN := docker-compose run 
APP_NAME := app
DOCK_X_APP := $(DOCKER_C) exec $(APP_NAME)

OUTPUT_COVERAGE := app/tests/coverage/

up: ## Start docker container
	$(DOCKER_C) pull
	$(DOCKER_C) up -d

rebuild: ## Rebuild docker container
	$(DOCKER_C) pull
	$(DOCKER_C) up --build -d

down: ## Stop docker container
	$(DOCKER_C) down --remove-orphans

ssh: ## Interactive access to container
	$(DOCKER_C) up -d
	$(DOCK_X_APP) /bin/sh

deps: ## Install dependencies
	$(DOCKER_C) up -d
	$(DOCK_X_APP) go mod download

tidy: ## Dowload and Clean dependencies
	$(DOCKER_C) up -d
	$(DOCK_X_APP) go mod tidy

vendor: ## Install vendor
	$(DOCKER_C) up -d
	$(DOCK_X_APP) go mod vendor

lint: ## Checks Code Style
	$(DOCKER_C) up -d
	$(DOCK_X_APP) ./run-lint.sh

vet: ## Finds issues in code
	$(DOCKER_C) up -d
	$(DOCK_X_APP) go vet ./...

fmt: ## Applies standard formatting
	$(DOCKER_C) up -d
	$(DOCK_X_APP) go fmt ./...

doc: ## Show package documentation
	$(DOCKER_C) up -d
	$(DOCK_X_APP) go doc github.com/jmsilvadev/golangtechtask/

test: ## Run all available tests
	$(DOCKER_C) down --remove-orphans
	$(DOCKER_C) up -d
	$(DOCK_X_APP) go test ./...

test.api: ## Run end to end tests
	$(DOCKER_C) down --remove-orphans
	$(DOCKER_C) up -d
	$(DOCK_X_APP) ./run-tests.sh

test.coverage: ## Run all available tests in a separate conteiner with coverage
	$(DOCKER_C) down --remove-orphans
	$(DOCKER_RUN) --entrypoint="./run-tests-coverage.sh" $(APP_NAME)
	open $(OUTPUT_COVERAGE)coverage.html >&- 2>&- || \
	xdg-open $(OUTPUT_COVERAGE)coverage.html >&- 2>&- || \
	gnome-open $(OUTPUT_COVERAGE)coverage.html >&- 2>&-
	$(DOCKER_C) down --remove-orphans

build: ## Build docker image in daemon mode
	$(DOCKER_C) build

go-build: ## Build a new binary server
	$(DOCKER_RUN) --entrypoint="./run-build-server.sh" $(APP_NAME)
	$(DOCKER_C) down --remove-orphans

cert: ## Build a certificate to use with ssl
	$(DOCKER_RUN) --entrypoint="./run-build-cert.sh" $(APP_NAME)
	$(DOCKER_C) down --remove-orphans

logs: ## Watch docker log files
	$(DOCKER_C) logs --tail 100 -f

proto: ## Generate protobuffers
	$(DOCK_X_APP) protoc --go_out=. --go_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false --go-grpc_opt=paths=source_relative --go-grpc_out=. ./api/service.proto

help:
	@grep -E '^[a-zA-Z._-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
