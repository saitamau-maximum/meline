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
	GetChannelsByName(ctx context.Context, name string) (*presenter.GetChannelsByNameResponse, error)
	CreateChannel(ctx context.Context, name string, userId uint64) error
	CreateChildChannel(ctx context.Context, name string, parentChannelId, userId uint64) error
	UpdateChannel(ctx context.Context, id uint64, name string) error
	DeleteChannel(ctx context.Context, id uint64) error
	DeleteChildChannel(ctx context.Context, childChannelId uint64) error
	JoinChannel(ctx context.Context, channelID uint64, userID uint64) error
	LeaveChannel(ctx context.Context, channelID uint64, userID uint64) error
}

type ChannelInteractor struct {
	channelRepository      repository.IChannelRepository
	channelUsersRepository repository.IChannelUsersRepository
	channelToChannelsRepository repository.IChannelToChannelsRepository
	userRepository         repository.IUserRepository
	channelPresenter       presenter.IChannelPresenter
}

func NewChannelInteractor(channelRepository repository.IChannelRepository, channelUsersRepository repository.IChannelUsersRepository, channelToChannelsRepository repository.IChannelToChannelsRepository, userRepository repository.IUserRepository, channelPresenter presenter.IChannelPresenter) *ChannelInteractor {
	return &ChannelInteractor{
		channelRepository:      channelRepository,
		channelUsersRepository: channelUsersRepository,
		channelToChannelsRepository: channelToChannelsRepository,
		userRepository:         userRepository,
		channelPresenter:       channelPresenter,
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

func (i *ChannelInteractor) GetChannelsByName(ctx context.Context, name string) (*presenter.GetChannelsByNameResponse, error) {
	channels, err := i.channelRepository.FindByName(ctx, name)
	if err != nil {
		if err == sql.ErrNoRows {
			return i.channelPresenter.GenerateGetChannelsByNameResponse([]*entity.Channel{}), nil
		}

		return i.channelPresenter.GenerateGetChannelsByNameResponse([]*entity.Channel{}), err
	}

	entitiedChannels := make([]*entity.Channel, len(channels))
	for i, channel := range channels {
		entitiedChannels[i] = channel.ToChannelEntity()
	}

	return i.channelPresenter.GenerateGetChannelsByNameResponse(entitiedChannels), nil
}

func (i *ChannelInteractor) CreateChannel(ctx context.Context, name string, userId uint64) error {
	id, err := i.channelRepository.Create(ctx, &model.Channel{Name: name})
	if err != nil {
		return err
	}

	if err := i.channelUsersRepository.Create(ctx, &model.ChannelUsers{ChannelID: id, UserID: userId}); err != nil {
		return err
	}

	return nil
}

func (i *ChannelInteractor) CreateChildChannel(ctx context.Context, name string, parentChannelId, userId uint64) error {
	id, err := i.channelRepository.Create(ctx, &model.Channel{Name: name})
	if err != nil {
		return err
	}

	if err := i.channelUsersRepository.Create(ctx, &model.ChannelUsers{ChannelID: id, UserID: userId}); err != nil {
		return err
	}

	if err := i.channelToChannelsRepository.Create(ctx, model.NewChannelChannelsModel(parentChannelId, id)); err != nil {
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

	if err := i.channelUsersRepository.DeleteByChannelID(ctx, id); err != nil {
		return err
	}

	if err := i.channelToChannelsRepository.DeleteFromParentChannelID(ctx, id); err != nil {
		return err
	}

	return nil
}

func (i *ChannelInteractor) DeleteChildChannel(ctx context.Context, childChannelId uint64) error {
	if err := i.channelToChannelsRepository.DeleteFromChildChannelID(ctx, childChannelId); err != nil {
		return err
	}

	if err := i.channelRepository.Delete(ctx, childChannelId); err != nil {
		return err
	}

	if err := i.channelUsersRepository.DeleteByChannelID(ctx, childChannelId); err != nil {
		return err
	}

	return nil
}

func (i *ChannelInteractor) JoinChannel(ctx context.Context, channelID uint64, userID uint64) error {
	if err := i.channelUsersRepository.Create(ctx, &model.ChannelUsers{ChannelID: channelID, UserID: userID}); err != nil {
		return err
	}

	return nil
}

func (i *ChannelInteractor) LeaveChannel(ctx context.Context, channelID uint64, userID uint64) error {
	if err := i.channelUsersRepository.Delete(ctx, channelID, userID); err != nil {
		return err
	}

	return nil
}
