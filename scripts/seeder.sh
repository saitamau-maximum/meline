#!/bin/bash

THIS_FILE_DIR="$(cd "$(dirname "$0")" && pwd)"
PROJECT_DIR="$(cd "${THIS_FILE_DIR}/.." && pwd)"
SERVER_DIR="${PROJECT_DIR}/server"
ENV_FILE="${PROJECT_DIR}/.env"

which dotenv >/dev/null 2>&1
if [ $? -ne 0 ]; then
    echo "dotenv-cli がインストールされていません！"
    echo "dotenvインストールしてから再度実行してください。"
    echo "`npm install -g dotenv-cli`"
    exit 1
fi

cd "${SERVER_DIR}"
dotenv -e "${ENV_FILE}" go run "${SERVER_DIR}/cmd/seeder/main.go"
