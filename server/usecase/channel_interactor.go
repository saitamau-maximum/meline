package usecase

import (
	"context"
	"database/sql"

	"github.com/saitamau-maximum/meline/domain/entity"
	"github.com/saitamau-maximum/meline/domain/repository"
	"github.com/saitamau-maximum/meline/models"
	"github.com/saitamau-maximum/meline/usecase/presenter"
)

type IChannelInteractor interface {
	GetAllChannels(ctx context.Context, userId uint64) (*presenter.GetAllChannelsResponse, error)
	GetChannelByID(ctx context.Context, id uint64) (*presenter.GetChannelByIdResponse, error)
	CreateChannel(ctx context.Context, name string, userId uint64) error
	CreateChildChannel(ctx context.Context, name string, parentChannelId, userId uint64) error
	UpdateChannel(ctx context.Context, id uint64, name string) error
	DeleteChannel(ctx context.Context, id uint64) error
	JoinChannel(ctx context.Context, channelID uint64, userID uint64) error
	LeaveChannel(ctx context.Context, channelID uint64, userID uint64) error
}

type ChannelInteractor struct {
	channelRepository           repository.IChannelRepository
	userRepository              repository.IUserRepository
	channelPresenter            presenter.IChannelPresenter
}

func NewChannelInteractor(channelRepository repository.IChannelRepository, userRepository repository.IUserRepository, channelPresenter presenter.IChannelPresenter) *ChannelInteractor {
	return &ChannelInteractor{
		channelRepository:           channelRepository,
		userRepository:              userRepository,
		channelPresenter:            channelPresenter,
	}
}

func (i *ChannelInteractor) GetAllChannels(ctx context.Context, userId uint64) (*presenter.GetAllChannelsResponse, error) {
	channels, err := i.userRepository.FindChannelsByUserID(ctx, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return i.channelPresenter.GenerateGetAllChannelsResponse([]*entity.Channel{}), nil
		}

		return i.channelPresenter.GenerateGetAllChannelsResponse([]*entity.Channel{}), err
	}

	entitiedChannels := make([]*entity.Channel, len(channels))
	for i, channel := range channels {
		entitiedChannels[i] = channel.ToChannelEntity()
	}

	return i.channelPresenter.GenerateGetAllChannelsResponse(entitiedChannels), nil
}

func (i *ChannelInteractor) GetChannelByID(ctx context.Context, id uint64) (*presenter.GetChannelByIdResponse, error) {
	channel, err := i.channelRepository.FindByID(ctx, id)
	if err != nil {
		return &presenter.GetChannelByIdResponse{}, err
	}

	return i.channelPresenter.GenerateGetChannelByIdResponse(channel.ToChannelEntity()), nil
}

func (i *ChannelInteractor) CreateChannel(ctx context.Context, name string, userId uint64) error {
	if err := i.channelRepository.Create(ctx, &model.Channel{Name: name}, userId); err != nil {
		return err
	}

	return nil
}

func (i *ChannelInteractor) CreateChildChannel(ctx context.Context, name string, parentChannelId, userId uint64) error {
	if err := i.channelRepository.CreateChildChannel(ctx, &model.Channel{Name: name}, parentChannelId, userId); err != nil {
		return err
	}

	return nil
}

func (i *ChannelInteractor) UpdateChannel(ctx context.Context, id uint64, name string) error {
	channel, err := i.channelRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	channel.Name = name
	if err := i.channelRepository.Update(ctx, channel); err != nil {
		return err
	}

	return nil
}

func (i *ChannelInteractor) DeleteChannel(ctx context.Context, id uint64) error {
	if err := i.channelRepository.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (i *ChannelInteractor) JoinChannel(ctx context.Context, channelID uint64, userID uint64) error {
	if err := i.channelRepository.JoinChannel(ctx, channelID, userID); err != nil {
		return err
	}

	return nil
}

func (i *ChannelInteractor) LeaveChannel(ctx context.Context, channelID uint64, userID uint64) error {
	if err := i.channelRepository.LeaveChannel(ctx, channelID, userID); err != nil {
		return err
	}

	return nil
}
