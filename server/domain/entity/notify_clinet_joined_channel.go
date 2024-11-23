package entity

type NotifyClientJoinedChannel struct {
	// NOTE: チャンネルごとに配信設定等を行う
	IsNotify bool // 将来的に通知の有無を設定できるようにする
}

func NewNotifyClientJoinedChannelEntity(isNotify bool) *NotifyClientJoinedChannel {
	return &NotifyClientJoinedChannel{
		IsNotify: isNotify,
	}
}
