.PHONY: start-prod start-dev stop lint
.SILENT:

run-dev: stop
	@ENV="dev" docker compose \
		-f deploy/docker-compose/docker-compose.yaml \
		--env-file env/.env \
		up --build -d

run-prod: stop
	@ENV="prod" docker compose \
		-f deploy/docker-compose/docker-compose.yaml \
		--env-file env/.env \
		up --build -d

stop:
	@if [ "$(RUNNING_CONTAINERS)" != "" ]; then \
		ENV="dev" docker compose \
			-f deploy/docker-compose/docker-compose.yaml \
			--env-file env/.env \
			down; \
	fi

lint:
	@cd $(CORE_PATH) && $(LOCAL_BIN_PATH)/golangci-lint run -c $(CONFIG) --path-prefix $(CORE_PATH)
	@for SERVICE in $(SERVICES); do \
		if [ "$$SERVICE" != "$(SERVICE_PATH)/service" ]; then \
			cd $$SERVICE && $(LOCAL_BIN_PATH)/golangci-lint run -c $(CONFIG) --path-prefix $$SERVICE; \
		fi; \
	done

install-deps:
	@GOBIN=$(LOCAL_BIN_PATH) go install github.com/pressly/goose/v3/cmd/goose@v3.17.0
	@GOBIN=$(LOCAL_BIN_PATH) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2
	@cd $(CORE_PATH) && go mod tidy
	@for SERVICE in $(SERVICES); do \
		if [ "$$SERVICE" != "$(SERVICE_PATH)/service" ]; then \
			cd $$SERVICE && go mod tidy; \
		fi; \
	done

# -------

LOCAL_BIN_PATH := $(CURDIR)/bin
ROOT_PATH := $(realpath .)
CONFIG := $(ROOT_PATH)/.golangci.yaml
CORE_PATH := $(ROOT_PATH)/core
SERVICE_PATH := $(ROOT_PATH)/services
SERVICES := $(filter-out $(SERVICE_PATH)/.,$(wildcard $(SERVICE_PATH)/*))

RUNNING_CONTAINERS := $(shell ENV=dev docker compose \
	-f deploy/docker-compose/docker-compose.yaml \
	--env-file env/.env \
	ps --status running -q \
)