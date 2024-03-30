package entity

type ChannelToChannels struct {
	ParentChannelID uint64
	ParentChannel   *Channel
	ChildChannelID  uint64
	ChildChannel    *Channel
}
