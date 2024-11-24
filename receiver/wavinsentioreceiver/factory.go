package wavinsentioreceiver

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

var (
	typeStr = component.MustNewType("wavinsentio")
)

const (
	defaultInterval = 1 * time.Minute
)

func createDefaultConfig() component.Config {
	return Config{
		Interval: defaultInterval.String(),
	}
}

func createMetricsReceiver(ctx context.Context, settings receiver.Settings, baseCfg component.Config, consumer consumer.Metrics) (receiver.Metrics, error) {
	logger := settings.Logger
	config := baseCfg.(Config)
	scraper := newScraper(config.Username, config.Password, logger)

	rcvr := wavinsentioReceiver{
		logger:              logger,
		consumer:            consumer,
		config:              &config,
		scraper:             scraper,
		locationUnmarshaler: &locationUnmarshaler{logger: logger},
		roomUnmarshaler:     &roomUnmarshaler{logger: logger},
	}

	return &rcvr, nil
}

func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		typeStr,
		createDefaultConfig,
		receiver.WithMetrics(createMetricsReceiver, component.StabilityLevelAlpha),
	)
}
