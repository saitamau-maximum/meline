# migration 手順

## 事前に用意するもの

### dotenvのインストール

`$ npm install -g dotenv-cli`  

このコマンドで`dotenv-cli`をインストールしてください。

## migrationファイルの作成手順

### goファイルの場合

`$ ./scripts/migrate.sh create_go {FILENAME}`

### SQLファイルの場合

`$ ./scripts/migrate.sh create_sql {FILENAME}`

## migrateの手順

### 1. dockerを立ち上げる

`$ ./scripts/setup.sh`

### 2. DBの初期化をおこなう

`$ ./scripts/migrate.sh init`  

### 3. migrationの実行

`$ ./scripts/migrate.sh migrate`

## rollbackを行う場合

`$ ./scripts/migrate.sh rollback`
