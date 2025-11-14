package config

import "fmt"

func (s ServerConfig) Addr() string {
	return fmt.Sprintf(":%d", s.Port)
}
