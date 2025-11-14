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
	CreateShortURL(context.Context, string, *string) (*model.ShortURL, error)
	Redirect(context.Context, string) (string, error)
	GetLinkInfo(context.Context, string) (*model.ShortURL, error)
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
		slog.Warn("Invalid request", "error", err)
		c.JSON(400, gin.H{"error": "Invalid request format"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 4*time.Second)
	defer cancel()

	shortURL, err := h.service.CreateShortURL(ctx, req.URL, req.CustomCode)
	if err != nil {
		slog.Error("Service error", "error", err, "URL", req.URL)

		switch err {
		case ers.ErrURLAlreadyExists:
			c.JSON(409, gin.H{"error": "URL already exists"})
			return
		case ers.ErrInvalidURL:
			c.JSON(400, gin.H{"error": "Invalid URL"})
			return
		default:
			c.JSON(500, gin.H{"error": "internal server error"})
		}
		return
	}

	response := dto.NewCreateShortURLResponse(shortURL, h.baseURL, req.CustomCode != nil)

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

	originalURL, err := h.service.Redirect(ctx, req.CustomCode)
	if err != nil {
		slog.Error("Redirect error", "error", err, "short code", req.CustomCode)

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

func (h *Handler) GetLinkInfo(c *gin.Context) {
	var req dto.GetLinkInfoRequest // 👈 твоя DTO для этого хендлера

	if err := c.ShouldBindUri(&req); err != nil {
		slog.Warn("Invalid short code in URI", "error", err)
		c.JSON(400, gin.H{"error": "Invalid short code format"})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	shortURL, err := h.service.GetLinkInfo(ctx, req.CustomCode)
	if err != nil {
		slog.Error("Get link info error", "error", err, "short code", req.CustomCode)

		switch err {
		case ers.ErrURLNotFound:
			c.JSON(404, gin.H{"error": "Short URL not found"})
		case ers.ErrURLExpired:
			c.JSON(410, gin.H{"error": "Short URL has expired"})
		default:
			c.JSON(500, gin.H{"error": "Internal server error"})
		}
		return
	}

	response := dto.NewShortURLInfoResponse(shortURL, h.baseURL)
	c.JSON(200, response)
}
