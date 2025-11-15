package service

import (
	"context"

	"github.com/alonsoF100/shotlink/internal/model"
	"github.com/alonsoF100/shotlink/internal/transport/http/dto"
)

type URLRepository interface {
}

type Service struct {
	urlRepo URLRepository
}

func New(urlRepo URLRepository) *Service {
	return &Service{urlRepo: urlRepo}
}

func (s Service) CreateShortURL(ctx context.Context, req dto.CreateShortURLRequest) (*model.ShortURL, error) {
	// Проверить есть ли для ссылки короткий код
	// Проверить занят ли короткий код
	// Добавить ссылку + короткий код в базу и вернуть структуру
	return nil, nil
}
func (s Service) Redirect(ctx context.Context, req dto.RedirectRequest) (string, error) {
	// Проверить есть ли такой статус код
	// Вернуть оригинальную ссылку
	return "", nil
}
