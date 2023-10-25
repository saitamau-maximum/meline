#!/bin/bash

THIS_FILE_DIR="$(cd "$(dirname "$0")" && pwd)"
PROJECT_DIR="$(cd "${THIS_FILE_DIR}/.." && pwd)"
CLIENT_DIR="${PROJECT_DIR}/client"
SERVER_DIR="${PROJECT_DIR}/server"
ENV_FILE="${PROJECT_DIR}/.env"

echo "Dockerを終了します..."

docker compose down

echo "Dockerを終了しました！"
