package model

import "time"

type ShortURL struct {
	ID          string    `json:"id" db:"id"`
	OriginalURL string    `json:"original_url" db:"original_url"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	ExpiresAt   time.Time `json:"expires_at" db:"expires_at"`
	ClickCount  int64     `json:"click_count" db:"click_count"`
	IsActive    bool      `json:"is_active" db:"is_active"`
}

type Click struct {
	ID          int64     `json:"id" db:"id"`
	ShortCode   string    `json:"short_code" db:"short_code"`
	ClickedAt   time.Time `json:"clicked_at" db:"clicked_at"`
	UserAgent   string    `json:"user_agent,omitempty" db:"user_agent"`
	IPAddress   string    `json:"ip_address,omitempty" db:"ip_address"`
	Referrer    string    `json:"referrer,omitempty" db:"referrer"`
	CountryCode string    `json:"country_code,omitempty" db:"country_code"`
	DeviceType  string    `json:"device_type,omitempty" db:"device_type"`
}
