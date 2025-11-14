package service

import (
	"context"

	"github.com/alonsoF100/shotlink/internal/model"
)

type URLRepository interface {
	Create(ctx context.Context, url *model.ShortURL) error
	FindByShortCode(ctx context.Context, code string) (*model.ShortURL, error)
	IncrementClickCount(ctx context.Context, code string) error
	Exists(ctx context.Context, code string) (bool, error)
}

type ClickRepository interface {
	Create(ctx context.Context, click *model.Click) error
	GetStats(ctx context.Context, code string) (int64, error)
}

type Service struct {
	urlRepo   URLRepository
	clickRepo ClickRepository
}

func New(urlRepo URLRepository, clickRepo ClickRepository) *Service {
	return &Service{urlRepo: urlRepo, clickRepo: clickRepo}
}

func (s *Service) CreateShortURL(ctx context.Context, url string, CustomCode *string) (*model.ShortURL, error) {
	return nil, nil
}
func (s *Service) Redirect(ctx context.Context, CustomCode string) (string, error) {
	return "", nil
}
func (s *Service) GetLinkInfo(ctx context.Context, CustomCode string) (*model.ShortURL, error) {
	return nil, nil
}
