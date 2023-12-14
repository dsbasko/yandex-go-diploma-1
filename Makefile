.PHONY: start-prod start-dev stop
.SILENT:

start-dev: stop
	@ENV="dev" docker compose \
		-f deploy/docker-compose/docker-compose.yaml \
		--env-file env/.env \
		up --build -d

start-prod: stop
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

# -------

RUNNING_CONTAINERS := $(shell ENV=dev docker compose \
	-f deploy/docker-compose/docker-compose.yaml \
	--env-file env/.env \
	ps --status running -q \
)