package postgres

import (
	"context"
	"log/slog"

	ers "github.com/alonsoF100/shotlink/internal/error"
	"github.com/alonsoF100/shotlink/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type URLRepository struct {
	pool *pgxpool.Pool
}

func NewURLRepository(pool *pgxpool.Pool) *URLRepository {
	return &URLRepository{pool: pool}
}

func (r URLRepository) FindByOriginalURL(ctx context.Context, originalURL string) (bool, error) {
	const query = `
		SELECT EXISTS(
			SELECT 1 
			FROM short_urls 
			WHERE original_url = $1 
			AND is_active = true)`

	var exists bool
	err := r.pool.QueryRow(ctx, query, originalURL).Scan(&exists)
	if err != nil {
		slog.Error("failed to check URL existence in database", "error", err, "originalURL", originalURL, "query", query)
		return false, ers.ErrDatabaseQuery
	}

	return exists, nil
}

func (r URLRepository) FindByShortCode(ctx context.Context, shortCode string) (bool, error) {
	const query = `
		SELECT EXISTS(
			SELECT 1 
			FROM short_urls 
			WHERE short_code = $1 
			AND is_active = true)`

	var exists bool
	err := r.pool.QueryRow(ctx, query, shortCode).Scan(&exists)
	if err != nil {
		slog.Error("failed to check shortCode existence in database", "error", err, "shortCode", shortCode, "query", query)
		return false, ers.ErrDatabaseQuery
	}

	return exists, nil
}

func (r *URLRepository) CreateShortURL(ctx context.Context, originalURL, shortCode string) (*model.ShortURL, error) {
	const query = `
        INSERT INTO short_urls (original_url, short_code) 
        VALUES ($1, $2)
        RETURNING id, original_url, short_code, created_at, click_count, is_active`

	var shortURL model.ShortURL
	err := r.pool.QueryRow(ctx, query, originalURL, shortCode).Scan(
		&shortURL.ID,
		&shortURL.OriginalURL,
		&shortURL.ShortCode,
		&shortURL.CreatedAt,
		&shortURL.ClickCount,
		&shortURL.IsActive,
	)

	if err != nil {
		slog.Error("failed to create short URL in database", "error", err, "originalURL", originalURL, "shortCode", shortCode)
		return nil, ers.ErrDatabaseQuery
	}

	slog.Info("successfully created short URL", "shortCode", shortCode, "originalURL", originalURL, "id", shortURL.ID)

	return &shortURL, nil
}

func (r URLRepository) FindByShortCodeAndIncrement(ctx context.Context, shortCode string) (string, error) {
	const query = `
		UPDATE short_urls 
		SET click_count = click_count + 1 
		WHERE short_code = $1 AND is_active = true
		RETURNING original_url, click_count`

	var originalURL string
	var clickCount int
	err := r.pool.QueryRow(ctx, query, shortCode).Scan(&originalURL, &clickCount)
	if err != nil {
		slog.Error("failed to find and increment short URL", "error", err, "shortCode", shortCode)
		return "", ers.ErrDatabaseQuery
	}

	slog.Info("successfully found and incremented click count", "shortCode", shortCode, "newClickCount", clickCount)

	return originalURL, nil
}
