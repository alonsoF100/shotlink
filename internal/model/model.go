package model

import "time"

type ShortURL struct {
	ID          string    `json:"id"`
	OriginalURL string    `json:"original_url"`
	ShortCode   string    `json:"short_code"`
	CreatedAt   time.Time `json:"created_at"`
	ExpiresAt   time.Time `json:"expires_at"`
	ClickCount  int64     `json:"click_count"`
	IsActive    bool      `json:"is_active"`
}

type Click struct {
	ID          int64     `json:"id"`
	ShortCode   string    `json:"short_code"`
	ClickedAt   time.Time `json:"clicked_at"`
	UserAgent   string    `json:"user_agent,omitempty"`
	IPAddress   string    `json:"ip_address,omitempty"`
	Referrer    string    `json:"referrer,omitempty"`
	CountryCode string    `json:"country_code,omitempty"`
	DeviceType  string    `json:"device_type,omitempty"`
}
