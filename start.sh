#!/bin/sh

set -e

echo "Running database migrations..."
/usr/local/bin/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "Starting the application..."
exec "$@"
