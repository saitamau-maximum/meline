package entity

type MessageToMessages struct {
	ParentMessageID string
	ParentMessage   *Message
	ChildMessageID  string
	ChildMessage    *Message
}

func NewMessageToMessagesEntity(parentMessageID string, parentMessage *Message, childMessageID string, childMessage *Message) *MessageToMessages {
	return &MessageToMessages{
		ParentMessageID: parentMessageID,
		ParentMessage:   parentMessage,
		ChildMessageID:  childMessageID,
		ChildMessage:    childMessage,
	}
}
