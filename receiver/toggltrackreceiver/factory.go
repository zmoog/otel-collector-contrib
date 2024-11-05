package toggltrackreceiver

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

var (
	typeStr = component.MustNewType("toggltrack")
)

const (
	defaultInterval = 1 * time.Minute
	defaultLookback = 24 * 30 * time.Hour // 30 days
)

func createDefaultConfig() component.Config {
	return Config{
		Interval: defaultInterval.String(),
		Lookback: defaultLookback.String(),
	}
}

func createLogsReceiver(ctx context.Context, settings receiver.Settings, baseCfg component.Config, consumer consumer.Logs) (receiver.Logs, error) {
	logger := settings.Logger
	config := baseCfg.(Config)
	scraper := NewScraper(config.APIToken, logger)

	rcvr := togglTrackReceiver{
		logger:   logger,
		consumer: consumer,
		config:   &config,
		scraper:  scraper,
	}

	return &rcvr, nil
}

func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		typeStr,
		createDefaultConfig,
		receiver.WithLogs(createLogsReceiver, component.StabilityLevelAlpha),
	)
}
