package model

type ChannelToChannels struct {
	ParentChannelID uint64   `bun:"parent_channel_id,pk"`
	ParentChannel   *Channel `bun:"rel:belongs-to,join:parent_channel_id=id"`
	ChildChannelID  uint64   `bun:"child_channel_id,pk"`
	ChildChannel    *Channel `bun:"rel:belongs-to,join:child_channel_id=id"`
}

func NewChannelChannelsModel(parentChannelID uint64, childChannelID uint64) *ChannelToChannels {
	return &ChannelToChannels{
		ParentChannelID: parentChannelID,
		ChildChannelID:  childChannelID,
	}
}
