package entity

type NotifyBroadcastCh struct {
	Message []byte
	UserIDs []uint64
}

func NewNotifyBroadcastChEntity(message []byte, userIDs []uint64) *NotifyBroadcastCh {
	return &NotifyBroadcastCh{
		Message: message,
		UserIDs: userIDs,
	}
}
