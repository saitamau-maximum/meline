# Auth

## GET /api/auth/login

OAuth認証画面へ遷移する

## GET /api/auth/callback

OAuth認証後に実行されるコールバック  
ユーザが存在しなければユーザの登録を行い、ログイン処理も同時に行う
発行されたAccess TokenはCookieに保存
