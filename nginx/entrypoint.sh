#!/bin/sh
set -e
envsubst < /etc/nginx/conf.d/default.conf.template > /etc/nginx/conf.d/default.conf
exec "$@"