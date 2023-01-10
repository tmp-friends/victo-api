package service

import (
	"context"

	"github.com/tmp-friends/victo-api/app/domain/models"
	"github.com/tmp-friends/victo-api/app/domain/repository"
	"github.com/tmp-friends/victo-api/app/usecase/dto"
)

type IHashtagService interface {
	FindHashtags(ctx context.Context, parameter dto.FindHashtagsParameter) (models.HashtagSlice, error)
}

type hashtagService struct {
	repo repository.IHashtagRepository
}

func NewHashtagService(hr repository.IHashtagRepository) IHashtagService {
	return &hashtagService{
		repo: hr,
	}
}

func (hs *hashtagService) FindHashtags(ctx context.Context, parameter dto.FindHashtagsParameter) (models.HashtagSlice, error) {
	return hs.repo.FindHashtags(ctx, parameter)
}
