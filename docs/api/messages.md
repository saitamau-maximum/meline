# Messages

## GET /api/channel/:channel_id/message

`:id`で指定したchannelのmessageを取得する

### Response Example

```json
{
    "messages": [
        {
            "id": "c7b2e0e2-4e4e-4a7d-9a0d-6e3e8b9c5d1f",
            "user": {
                "id": 1,
                "name": "test-user",
                "image_url": "https://example.com"
            },
            "content": "test-message",
            "reply_to_message": {
                "id": "a1b2c3d4-e5f6-g7h8-i9j0-k1l2m3n4o5p",
                "user": {
                    "id": 2,
                    "name": "test-reply-user",
                    "image_url": "https://example.com"
                },
                "content": "test-reply"
            },
            "created_at": "2024-04-25 10:19:24 +0000 UTC",
            "updated_at": "2024-04-25 10:19:24 +0000 UTC"
        },
        {
            "id": "d3f4a6b8-1c2d-3e4f-5a6b-7c8d9e0f1a2b",
            "user": {
                "id": 2,
                "name": "test-user-2",
                "image_url": "https://example.com"
            },
            "content": "test-message",
            "replay_to_message": null,
            "created_at": "2024-04-25 10:19:24 +0000 UTC",
            "updated_at": "2024-04-25 10:19:24 +0000 UTC"
        }
    ]
}
```

## POST /api/channel/:channel_id/message

`:channel_id`で指定したChannelにMessageを送信する

### Request Example

```json
{
    "content": "test-message"
}
```

- Validation
   - `content`は必須
   - 2000文字以内

## POST /api/channel/:channel_id/message/:id/reply

`:id`で指定したMessageに対してリプライを送信する

### Request Example

```json
{
    "content": "test-message"
}
```

- Validation
   - `content`は必須
   - 2000文字以内

## PUT /api/channel/:channel_id/message/:id

`:id`で指定したMessageの内容の編集を行う

### Request Example

```json
{
    "content": "test-message"
}
```

- Validation
   - `content`は必須
   - 2000文字以内

## PUT /api/channel/:channel_id/message/:id

`:id`で指定したメッセージの削除を行う
