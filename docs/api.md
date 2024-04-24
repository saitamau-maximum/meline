## GET /api/channel

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

## GET /api/channel/:id

Channelの詳細を返す

```json
{
	"channel": {
		"name": "test",
		"users": [
			{
				"name": "test",
				"image_url": "test",
			},
		]
	}
}
```
