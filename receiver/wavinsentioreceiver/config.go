package wavinsentioreceiver

import (
	"fmt"
	"time"
)

type Config struct {
	Interval string `mapstructure:"interval"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func (cfg Config) Validate() error {
	interval, err := time.ParseDuration(cfg.Interval)
	if err != nil {
		return fmt.Errorf("invalid interval: %w", err)
	}
	if interval < 1*time.Minute {
		return fmt.Errorf("interval must be at least 1 minute")
	}

	if cfg.Username == "" {
		return fmt.Errorf("username is required")
	}

	if cfg.Password == "" {
		return fmt.Errorf("password is required")
	}

	return nil
}
