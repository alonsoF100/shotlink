package error

import "errors"

var (
	ErrURLAlreadyExists = errors.New("URL already exists")
	ErrShortCodeTaken   = errors.New("short code already taken")
	ErrShortCodeNotFound = errors.New("short code not found")
	ErrURLNotFound      = errors.New("URL not found")
	ErrDatabaseQuery    = errors.New("database query failed")
)
