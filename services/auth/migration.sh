#!/bin/bash

export DSN="host=$AUTH_PSQL_CONTAINER_NAME port=${AUTH_PSQL_PORT} dbname=$AUTH_PSQL_DB user=$AUTH_PSQL_USER password=$AUTH_PSQL_PASS"
sleep 2 && goose -dir ./migrations postgres "${DSN}" up -v