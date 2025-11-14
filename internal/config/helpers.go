package config

import (
	"fmt"
)

func (s ServerConfig) Addr() string {
	return fmt.Sprintf(":%d", s.Port)
}

func (d DatabaseConfig) ConStr() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		d.User, d.Password, d.Host, d.Port, d.Name, d.SSLMode)
}
