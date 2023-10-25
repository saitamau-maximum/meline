#!/bin/bash

THIS_FILE_DIR="$(cd "$(dirname "$0")" && pwd)"
PROJECT_DIR="$(cd "${THIS_FILE_DIR}/.." && pwd)"
ENV_FILE="${PROJECT_DIR}/.env"

if [ ! -f "${ENV_FILE}" ]; then
    echo "環境変数ファイルが存在しません！"
    echo "環境変数ファイルを作成してから再度実行してください。"
    exit 1
fi

docker compose -f docker-compose.prod.yml down
docker compose -f docker-compose.prod.yml up -d --build
