package service

import (
	"context"
	"log/slog"

	ers "github.com/alonsoF100/shotlink/internal/error"
	"github.com/alonsoF100/shotlink/internal/model"
	"github.com/alonsoF100/shotlink/internal/transport/http/dto"
)

type URLRepository interface {
	FindByOriginalURL(ctx context.Context, originalURL string) (bool, error)
	FindByShortCode(ctx context.Context, shortCode string) (bool, error)
	CreateShortURL(ctx context.Context, originalURL, shortCode string) (*model.ShortURL, error)
	FindByShortCodeAndIncrement(ctx context.Context, shortCode string) (string, error)
}

type Service struct {
	urlRepo URLRepository
}

func New(urlRepo URLRepository) *Service {
	return &Service{urlRepo: urlRepo}
}

func (s Service) CreateShortURL(ctx context.Context, req dto.CreateShortURLRequest) (*model.ShortURL, error) {
	exists, err := s.urlRepo.FindByOriginalURL(ctx, req.OriginalURL)
	if err != nil {
		slog.Error("failed to check original URL existence", "error", err, "originalURL", req.OriginalURL)
		return nil, err
	}
	if exists {
		slog.Info("original URL already exists", "originalURL", req.OriginalURL)
		return nil, ers.ErrURLAlreadyExists
	}

	exists, err = s.urlRepo.FindByShortCode(ctx, req.ShortCode)
	if err != nil {
		slog.Error("failed to check short code existence", "error", err, "shortCode", req.ShortCode)
		return nil, err
	}
	if exists {
		slog.Info("short code is already taken", "shortCode", req.ShortCode)
		return nil, ers.ErrShortCodeTaken
	}

	shortURL, err := s.urlRepo.CreateShortURL(ctx, req.OriginalURL, req.ShortCode)
	if err != nil {
		slog.Error("failed to create short URL", "error", err, "originalURL", req.OriginalURL, "shortCode", req.ShortCode)
		return nil, err
	}

	slog.Info("successfully created short URL", "shortCode", req.ShortCode, "originalURL", req.OriginalURL, "id", shortURL.ID)
	return shortURL, nil
}

func (s Service) Redirect(ctx context.Context, req dto.RedirectRequest) (string, error) {
	exists, err := s.urlRepo.FindByShortCode(ctx, req.ShortCode)
	if err != nil {
		slog.Error("failed to check short code existence", "error", err, "shortCode", req.ShortCode)
		return "", err
	}
	if !exists {
		slog.Info("short code not found", "shortCode", req.ShortCode)
		return "", ers.ErrShortCodeNotFound
	}

	originalURL, err := s.urlRepo.FindByShortCodeAndIncrement(ctx, req.ShortCode)
	if err != nil {
		slog.Error("failed to get original URL and increment counter", "error", err, "shortCode", req.ShortCode)
		return "", err
	}

	slog.Info("successful redirect", "shortCode", req.ShortCode)
	return originalURL, nil
}
