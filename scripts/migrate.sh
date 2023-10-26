#!/bin/bash

THIS_FILE_DIR="$(cd "$(dirname "$0")" && pwd)"
PROJECT_DIR="$(cd "${THIS_FILE_DIR}/.." && pwd)"
SERVER_DIR="${PROJECT_DIR}/server"
ENV_FILE="${PROJECT_DIR}/.env"

cd "${SERVER_DIR}"

dotenv -e "${ENV_FILE}" go run "${SERVER_DIR}/cmd/migrate/main.go" db $1 $2