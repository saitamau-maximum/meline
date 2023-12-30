package request

import (
	"errors"
	"unicode/utf8"

	"github.com/saitamau-maximum/meline/config"
)

type CreateMessageRequest struct {
	ChannelID uint64 `json:"channel_id"`
	UserID    uint64 `json:"user_id"`
	ReplyToID string `json:"reply_to_id"`
	Content   string `json:"content"`
}

func (r *CreateMessageRequest) Validate() error {
	if r.Content == "" {
		return errors.New(config.ERR_EMPTY_MESSAGE)
	}

	if utf8.RuneCountInString(r.Content) > config.MAX_MESSAGE_LENGTH {
		return errors.New(config.ERR_TOO_LONG_MESSAGE)
	}

	return nil
}
