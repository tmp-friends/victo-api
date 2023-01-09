package usecase

import (
	"context"

	"github.com/tmp-friends/victo-api/app/domain/models"
	"github.com/tmp-friends/victo-api/app/domain/service"
)

type IHashtagUsecase interface {
	FindHashtags(ctx context.Context) (models.HashtagSlice, error)
}

type hashtagUsecase struct {
	service service.IHashtagService
}

func NewHashtagUsecase(hs service.IHashtagService) IHashtagUsecase {
	return &hashtagUsecase{
		service: hs,
	}
}

func (hu *hashtagUsecase) FindHashtags(ctx context.Context) (models.HashtagSlice, error) {
	mhSlice, err := hu.service.FindHashtags(ctx)
	if err != nil {
		return nil, err
	}

	return mhSlice, nil
}
