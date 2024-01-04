package request

import (
	"errors"
	"unicode/utf8"

	"github.com/saitamau-maximum/meline/config"
)

type CreateMessageRequest struct {
	Content string `json:"content"`
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
