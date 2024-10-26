package zcsazzurroreceiver

import (
	"fmt"
	"time"
)

type Config struct {
	Interval string `mapstructure:"interval"`
	ClientID string `mapstructure:"client_id"`
	AuthKey  string `mapstructure:"auth_key"`
	ThingKey string `mapstructure:"thing_key"`
}

func (cfg *Config) Validate() error {
	// Validate that min interval is 5 minutes
	interval, err := time.ParseDuration(cfg.Interval)
	if err != nil {
		return fmt.Errorf("invalid interval: %w", err)
	}
	if interval < 5*time.Minute {
		// ZCS updates data every 5 minutes, so it makes no sense
		// to have a smaller interval.
		return fmt.Errorf("interval must be at least 5 minutes")
	}

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
