package entity

type NotifyBroadcastCh struct {
	Message   []byte
	SenderID  uint64
	UserIDs   []uint64
	ChannelID uint64
}

func NewNotifyBroadcastChEntity(message []byte, senderID uint64, userIDs []uint64, channelID uint64) *NotifyBroadcastCh {
	return &NotifyBroadcastCh{
		Message:   message,
		SenderID:  senderID,
		UserIDs:   userIDs,
		ChannelID: channelID,
	}
}
