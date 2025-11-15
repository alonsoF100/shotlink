package handler

import (
	"context"
	"log/slog"
	"time"

	ers "github.com/alonsoF100/shotlink/internal/error"
	"github.com/alonsoF100/shotlink/internal/model"
	"github.com/alonsoF100/shotlink/internal/transport/http/dto"
	"github.com/gin-gonic/gin"
)

type Service interface {
	CreateShortURL(ctx context.Context, req dto.CreateShortURLRequest) (*model.ShortURL, error)
	Redirect(ctx context.Context, req dto.RedirectRequest) (string, error)
}

type Handler struct {
	service Service
	baseURL string
}

func New(service Service, baseURL string) *Handler {
	return &Handler{service: service, baseURL: baseURL}
}

func (h *Handler) CreateShortURL(c *gin.Context) {
	var req dto.CreateShortURLRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Warn("invalid request", "error", err)
		c.JSON(400, gin.H{"error": "invalid request format"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 4*time.Second)
	defer cancel()

	shortURL, err := h.service.CreateShortURL(ctx, req)
	if err != nil {
		slog.Error("Service error", "error", err, "URL", req.URL)

		switch err {
		case ers.ErrURLAlreadyExists:
			c.JSON(409, gin.H{"error": "URL already exists"})
			return
		case ers.ErrInvalidURL:
			c.JSON(400, gin.H{"error": "invalid URL"})
			return
		default:
			c.JSON(500, gin.H{"error": "internal server error"})
		}
		return
	}

	response := shortURL

	c.JSON(201, response)
}

func (h *Handler) Redirect(c *gin.Context) {
	var req dto.RedirectRequest

	if err := c.ShouldBindUri(&req); err != nil {
		slog.Warn("Invalid short code in URI", "error", err)
		c.JSON(400, gin.H{"error": "Invalid short code format"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	originalURL, err := h.service.Redirect(ctx, req)
	if err != nil {
		slog.Error("Redirect error", "error", err, "short code", req.ShortCode)

		switch err {
		case ers.ErrURLNotFound:
			c.JSON(404, gin.H{"error": "Short URL not found"})
		case ers.ErrURLExpired:
			c.JSON(410, gin.H{"error": "Short URL has expired"})
		case ers.ErrURLBlocked:
			c.JSON(403, gin.H{"error": "URL is not allowed"})
		default:
			c.JSON(500, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.Redirect(302, originalURL)
}
