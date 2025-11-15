package service

import (
	"context"
	"errors"
	"log/slog"
	"time"

	ers "github.com/alonsoF100/shotlink/internal/error"
	"github.com/alonsoF100/shotlink/internal/model"
	"github.com/google/uuid"
)

type URLRepository interface {
	FindByOriginalURL(ctx context.Context, originalURL string) (*model.ShortURL, error)
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

func (s *Service) CreateShortURL(ctx context.Context, url string, customCode *string) (*model.ShortURL, error) {
	existingURL, err := s.urlRepo.FindByOriginalURL(ctx, url)
	if err != nil && !errors.Is(err, ers.ErrURLNotFound) {
		return nil, err
	}

	if existingURL != nil {
		if customCode == nil {
			return existingURL, nil
		}

		if *customCode == existingURL.ShortCode {
			return existingURL, nil
		}

		return nil, ers.ErrURLAlreadyExists
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, ers.ErrIDGenerationFailed
	}

	expiresAt := time.Now().AddDate(1, 0, 0)

	shortURL := &model.ShortURL{
		ID:          id,
		OriginalURL: url,
		CreatedAt:   time.Now(),
		ExpiresAt:   &expiresAt,
		ClickCount:  0,
		IsActive:    true,
	}

	if customCode != nil {
		exists, err := s.urlRepo.Exists(ctx, *customCode)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, ers.ErrShortCodeTaken
		}

		shortURL.ShortCode = *customCode
	} else {
		code, err := s.generateUniqueCode(ctx)
		if err != nil {
			return nil, err
		}
		shortURL.ShortCode = code
	}

	err = s.urlRepo.Create(ctx, shortURL)
	if err != nil {
		return nil, err
	}

	return shortURL, nil
}
func (s *Service) Redirect(ctx context.Context, customCode string) (string, error) {
	return s.RedirectWithAnalytics(ctx, customCode, nil)
}

func (s *Service) RedirectWithAnalytics(ctx context.Context, customCode string, clickData *model.Click) (string, error) {
	url, err := s.urlRepo.FindByShortCode(ctx, customCode)
	if err != nil {
		if errors.Is(err, ers.ErrURLNotFound) {
			return "", ers.ErrURLNotFound
		}
		return "", err
	}

	if !url.IsActive {
		return "", ers.ErrURLInactive
	}

	if url.ExpiresAt != nil && url.ExpiresAt.Before(time.Now()) {
		return "", ers.ErrURLExpired
	}

	err = s.urlRepo.IncrementClickCount(ctx, customCode)
	if err != nil {
		slog.Warn("error fail to incr click")
	}

	if clickData != nil {
		clickData.ShortCode = customCode
		clickData.ClickedAt = time.Now()

		go func() {
			_ = s.clickRepo.Create(context.Background(), clickData)
		}()
	}

	return url.OriginalURL, nil
}

func (s *Service) GetLinkInfo(ctx context.Context, сustomCode string) (*model.ShortURL, error) {
	if сustomCode == "" {
		return nil, ers.ErrInvalidShortCode
	}

	url, err := s.urlRepo.FindByShortCode(ctx, сustomCode)
	if err != nil {
		if errors.Is(err, ers.ErrURLNotFound) {
			return nil, ers.ErrURLNotFound
		}
		return nil, err
	}

	return url, nil
}

func (s *Service) generateUniqueCode(ctx context.Context) (string, error) {
	const maxAttempts = 5

	for i := 0; i < maxAttempts; i++ {
		code := generateShortCode()

		exists, err := s.urlRepo.Exists(ctx, code)
		if err != nil {
			return "", err
		}

		if !exists {
			return code, nil
		}
	}

	return "", ers.ErrShortCodeGenerationFailed
}

func generateShortCode() string {
	return generateBase62FromUUID()
}

func generateBase62FromUUID() string {
	uuid := uuid.New()
	return uuid.String()[:8]
}
