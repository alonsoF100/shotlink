package error

import "errors"

var (
	ErrURLAlreadyExists   = errors.New("URL already exists")
	ErrInvalidURL         = errors.New("invalid URL format")
	ErrShortCodeTaken     = errors.New("short code already taken")
	ErrURLNotFound        = errors.New("URL not found")
	ErrURLExpired         = errors.New("URL has expired")
	ErrURLBlocked         = errors.New("URL is not allowed")
	ErrIDGenerationFailed = errors.New("failde to generate UUID")
	ErrShortCodeGenerationFailed = errors.New("failed to generate short code")
	ErrInvalidShortCode = errors.New("ivalid short code")
	ErrURLInactive = errors.New("inactive URL")
)
