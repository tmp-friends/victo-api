package service

import (
	"context"

	"github.com/tmp-friends/victo-api/app/domain/models"
	"github.com/tmp-friends/victo-api/app/domain/repository"
)

type IHashtagService interface {
	FindHashtags(ctx context.Context) (models.HashtagSlice, error)
}

type hashtagService struct {
	repo repository.IHashtagRepository
}

func NewHashtagService(hr repository.IHashtagRepository) IHashtagService {
	return &hashtagService{
		repo: hr,
	}
}

func (hs *hashtagService) FindHashtags(ctx context.Context) (models.HashtagSlice, error) {
	return hs.repo.FindHashtags(ctx)
}
