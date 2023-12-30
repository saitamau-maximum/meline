package presenter

import (
	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/usecase/presenter"
)

type messagePresenter struct{}

func NewMessagePresenter() presenter.IMessagePresenter {
	return &messagePresenter{}
}

func (p *messagePresenter) GenerateGetMessageByIDResponse(message *entity.Message) *presenter.GetMessageByIDResponse {
	user := &presenter.User{
		ID:        message.User.ID,
		Name:      message.User.Name,
		ImageURL:  message.User.ImageURL,
	}

	comments := make([]*presenter.Comments, len(message.Comments))
	for _, comment := range message.Comments {
		comments = append(comments, &presenter.Comments{
			ID:        comment.ID,
			User:      &presenter.User{
				ID:       comment.User.ID,
				Name:     comment.User.Name,
				ImageURL: comment.User.ImageURL,
			},
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt.String(),
			UpdatedAt: comment.UpdatedAt.String(),
		})
	}

	return &presenter.GetMessageByIDResponse{
		Thread: &presenter.Thread{
			ID:        message.ID,
			User:      user,
			Content:   message.Content,
			Comments:  comments,
			CreatedAt: message.CreatedAt.String(),
			UpdatedAt: message.UpdatedAt.String(),
		},
	}
}
