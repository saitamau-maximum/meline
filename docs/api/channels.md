# Channels

## GET /api/channel

ログインしているユーザの属するChannel一覧を返す

### Response Example

```json
{
	"channels": [
		{
			"id": 1,
			"name": "test-channel",
		},
	]
}
```

## GET /api/channel/:id

`:id`で指定したChannelの詳細を返す

### Response Example

```json
{
	"channel": {
		"name": "test-channel",
		"users": [
			{
				"name": "test-user",
				"image_url": "https://exapmle.com",
			},
		]
	}
}
```

## POST /api/channel/:id/join

`:id`で指定したChannelに参加する

## POST /api/channel

Channelを作成する

### Request Exapmle

```json
{
	"name": "test-channel"
}
```

## POST /api/channel/:id/create

`:id`で指定したChannel内に新たにChannelを作成する

### Request Exapmle

```json
{
	"name": "test-child-channel"
}
```

## PUT /api/channel/:id

Channel名の変更を行う(正直いるかは怪しい)

### Request Exapmle

```json
{
	"name": "test-update-channel"
}
```

## DELETE /api/channel/:id

Channelの削除を行う

## DELETE /api/channel/:id/leave

Channelから退出する
