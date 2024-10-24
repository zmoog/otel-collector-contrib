package zcsazzurroreceiver

import (
	"fmt"
)

type Config struct {
	Interval string `mapstructure:"interval"`
	ClientID string `mapstructure:"client_id"`
	AuthKey  string `mapstructure:"auth_key"`
	ThingKey string `mapstructure:"thing_key"`
}

func (cfg *Config) Validate() error {
	if cfg.AuthKey == "" {
		return fmt.Errorf("auth_key is required")
	}
	if cfg.ClientID == "" {
		return fmt.Errorf("client_id is required")
	}
	if cfg.ThingKey == "" {
		return fmt.Errorf("thing_key is required")
	}

	return nil
}
