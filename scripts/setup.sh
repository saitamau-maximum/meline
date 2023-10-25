#!/bin/bash

THIS_FILE_DIR="$(cd "$(dirname "$0")" && pwd)"
PROJECT_DIR="$(cd "${THIS_FILE_DIR}/.." && pwd)"
CLIENT_DIR="${PROJECT_DIR}/client"
SERVER_DIR="${PROJECT_DIR}/server"
ENV_FILE="${PROJECT_DIR}/.env"

if [ ! -f "${ENV_FILE}" ]; then
    echo "環境変数ファイルが存在しません！"
    echo "環境変数ファイルを作成してから再度実行してください。"
    exit 1
fi

echo "フロントエンドのセットアップを開始します..."

echo "Installing dependencies..."
which npm >/dev/null 2>&1
if [ $? -ne 0 ]; then
    echo "npm がインストールされていません！"
    echo "インストールしてから再度実行してください。"
    exit 1
fi

cd "${CLIENT_DIR}"
pnpm i

echo "フロントエンドのセットアップが完了しました！"

echo "バックエンドのセットアップを開始します..."

echo "Installing dependencies..."
which go >/dev/null 2>&1
if [ $? -ne 0 ]; then
    echo "go がインストールされていません！"
    echo "インストールしてから再度実行してください。"
    exit 1
fi

cd "${SERVER_DIR}"
go mod tidy

echo "バックエンドのセットアップが完了しました！"

echo "Dockerのセットアップを開始します..."

docker compose down
docker compose build
docker compose up -d

echo "Dockerのセットアップが完了しました！"
