# Websocket

## GET /api/ws/:channel_id

`:channel_id`で指定したChannelに送信されたMessageを受信するWebsocket

### Response Example

```json
{
    "message": {
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
    "channel_id": 1
}
```

## GET /api/ws/notify

自分が所属しているChannelに送信されたメッセージを受信する

### 通知の種類

- 着信通知(Message)
    - TypeID: 0 (`config.NOTIFY_MESSAGE`)

### Response Example

```json
{
    "notify_meta": {
        "type_id": 0
    },
    "payload": {
        "channel_id": 1,
        "message": {
            "id": "c7b2e0e2-4e4e-4a7d-9a0d-6e3e8b9c5d1f",
            "user": {
                "id": 1,
                "name": "test-user",
                "image_url": "https://example.com"
            },
            "content": "test-message",
            "channel_id": 1
        }
    }
}
```
