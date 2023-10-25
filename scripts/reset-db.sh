#!/bin/bash

THIS_FILE_DIR="$(cd "$(dirname "$0")" && pwd)"
PROJECT_DIR="$(cd "${THIS_FILE_DIR}/.." && pwd)"
CLIENT_DIR="${PROJECT_DIR}/client"
SERVER_DIR="${PROJECT_DIR}/server"
DB_DATA_DIR="${PROJECT_DIR}/etc/mysql/dbdata"
ENV_FILE="${PROJECT_DIR}/.env"

echo "DBデータを削除してもよろしいですか？"
echo "削除する場合は y を入力してください。"
echo "削除しない場合は n を入力してください。"

read -p "y/n: " yn

case "$yn" in [yY]*) ;; *)
    echo "削除せずに終了します..."
    exit
    ;;
esac

echo "Dockerを停止します..."

docker compose down

echo "Dockerを停止しました！"

echo "DBデータを削除します..."

sudo rm -rf "${DB_DATA_DIR}"

echo "DBデータを削除しました！"

echo "Dockerを立ち上げます..."

docker compose up -d --build

echo "Dockerを立ち上げました！"
