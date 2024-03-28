package model

type MessageToMessages struct {
	ParentMessageID string   `bun:"parent_message_id,pk"`
	ParentMessage   *Message `bun:"rel:belongs-to,join:parent_message_id=id"`
	ChildMessageID  string   `bun:"child_message_id,pk"`
	ChildMessage    *Message `bun:"rel:belongs-to,join:child_message_id=id"`
}

func NewMessageToMessagesModel(parentMessageID string, childMessageID string) *MessageToMessages {
	return &MessageToMessages{
		ParentMessageID: parentMessageID,
		ChildMessageID:  childMessageID,
	}
}
