#!/bin/sh
set -e

echo "Running migrations..."
migrate -path /app/migrations -database "$APP_DSN" up

echo "Starting application..."
exec ./main
