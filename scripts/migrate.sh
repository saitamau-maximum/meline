#!/bin/bash

THIS_FILE_DIR="$(cd "$(dirname "$0")" && pwd)"
PROJECT_DIR="$(cd "${THIS_FILE_DIR}/.." && pwd)"
SERVER_DIR="${PROJECT_DIR}/server"
ENV_FILE="${PROJECT_DIR}/.env"

which dotenv >/dev/null 2>&1
if [ $? -ne 0 ]; then
    echo "dotenv-cli がインストールされていません！"
    echo "dotenvインストールしてから再度実行してください。"
    echo "\`npm install -g dotenv-cli\`"
    exit 1
fi

# docker containerが起動しているか確認
docker compose ps | grep server >/dev/null 2>&1
if [ $? -ne 0 ]; then
    echo "docker containerが起動していません！"
    echo "docker containerを起動してから再度実行してください。"
    echo "\`./scripts/start.sh\`"
    exit 1
fi

cd "${SERVER_DIR}"

dotenv -e "${ENV_FILE}" docker compose exec server go run /app/cmd/migrate/main.go db $1 $2
