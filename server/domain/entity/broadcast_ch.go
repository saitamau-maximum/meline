package entity

type BroadcastCh struct {
	Message   []byte
	ChannelID uint64
}

func NewBroadcastChEntity(message []byte, channel uint64) *BroadcastCh {
	return &BroadcastCh{
		Message:   message,
		ChannelID: channel,
	}
}
