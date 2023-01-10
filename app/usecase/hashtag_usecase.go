package usecase

import (
	"context"
	"net/url"

	"github.com/tmp-friends/victo-api/app/domain/models"
	"github.com/tmp-friends/victo-api/app/domain/service"
	"github.com/tmp-friends/victo-api/app/usecase/dto"
)

type IHashtagUsecase interface {
	CreateParameter(qms url.Values) dto.FindHashtagsParameter
	FindHashtags(ctx context.Context, parameter dto.FindHashtagsParameter) (models.HashtagSlice, error)
}

type hashtagUsecase struct {
	service service.IHashtagService
}

func NewHashtagUsecase(hs service.IHashtagService) IHashtagUsecase {
	return &hashtagUsecase{
		service: hs,
	}
}

func (hu *hashtagUsecase) CreateParameter(qms url.Values) dto.FindHashtagsParameter {
	return dto.Create(qms)
}

func (hu *hashtagUsecase) FindHashtags(ctx context.Context, parameter dto.FindHashtagsParameter) (models.HashtagSlice, error) {
	mhSlice, err := hu.service.FindHashtags(ctx, parameter)
	if err != nil {
		return nil, err
	}

	return mhSlice, nil
}
