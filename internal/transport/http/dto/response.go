package dto

import (
	"time"

	"github.com/alonsoF100/shotlink/internal/model"
)

type CreateShortURLResponse struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
	ShortCode   string `json:"short_code"`
	CreatedAt   string `json:"created_at"`
	IsCustom    bool   `json:"is_custom"`
}

func NewCreateShortURLResponse(shortURL *model.ShortURL, baseURL string, isCustom bool) CreateShortURLResponse {
	return CreateShortURLResponse{
		ShortURL:    baseURL + "/" + shortURL.ShortCode,
		OriginalURL: shortURL.OriginalURL,
		ShortCode:   shortURL.ShortCode,
		CreatedAt:   shortURL.CreatedAt.Format(time.RFC3339),
		IsCustom:    isCustom,
	}
}

type ShortURLInfoResponse struct {
	ID          string `json:"id"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
	ShortCode   string `json:"short_code"`
	CreatedAt   string `json:"created_at"`
	ExpiresAt   string `json:"expires_at"`
	ClickCount  int64  `json:"click_count"`
	IsActive    bool   `json:"is_active"`
}

func NewShortURLInfoResponse(shortURL *model.ShortURL, baseURL string) ShortURLInfoResponse {
	response := ShortURLInfoResponse{
		ID:          shortURL.ID.String(),
		ShortURL:    baseURL + "/" + shortURL.ShortCode,
		OriginalURL: shortURL.OriginalURL,
		ShortCode:   shortURL.ShortCode,
		CreatedAt:   shortURL.CreatedAt.Format(time.RFC3339),
		ClickCount:  shortURL.ClickCount,
		IsActive:    shortURL.IsActive,
	}

	if shortURL.ExpiresAt != nil {
		response.ExpiresAt = shortURL.ExpiresAt.Format(time.RFC3339)
	}

	return response
}
