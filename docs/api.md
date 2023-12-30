## GET /api/channels

ログインしているユーザの属するChannel一覧を返す

```json
{
	"channels": [
		{
			"id": 1,
			"name": "test",
		},
	]
}
```

## GET /api/channels/:id

Channelの詳細を返す

```json
{
	"channel": {
		"name": "test",
		"channels": [
			{
				"id": 1,
				"name": "test",
			},
		],
		"users": [
			{
				"name": "test",
				"image_url": "test",
			},
		],
		"messages": [
			{
				"id": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
				"user": {
					"id": 1,
					"name": "test",
					"image_url": "test",
				},
				"content": "test",
				// リプライ先メッセージ
				"replyToMessage": {
					"id": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
					"user": {
						"id": 1,
						"name": "test",
						"image_url": "test",
					},
					"content": "test"
				},
				"created_at": "yyyy-mm-dd-hh-ii-ss",
				"updated_at": "yyyy-mm-dd-hh-ii-ss",
			},
		]
	}
}
```
