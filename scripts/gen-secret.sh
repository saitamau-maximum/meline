#!/bin/bash

THIS_FILE_DIR="$(cd "$(dirname "$0")" && pwd)"
PROJECT_DIR="$(cd "${THIS_FILE_DIR}/.." && pwd)"
SERVER_DIR="${PROJECT_DIR}/server"

cd "${SERVER_DIR}"

go run "${SERVER_DIR}/cmd/gen_secret/main.go"
