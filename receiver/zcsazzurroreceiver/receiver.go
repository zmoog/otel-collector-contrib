package zcsazzurroreceiver

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.uber.org/zap"
)

type zcsazzurroReceiver struct {
	cancel    context.CancelFunc
	logger    *zap.Logger
	consumer  consumer.Metrics
	config    *Config
	scraper   *Scraper
	marshaler *azzurroRealtimeDataMarshaler
}

func (z *zcsazzurroReceiver) Start(ctx context.Context, host component.Host) error {
	z.logger.Info("Starting zcsazzurro receiver")

	_ctx, cancel := context.WithCancel(ctx)
	z.cancel = cancel

	interval, _ := time.ParseDuration(z.config.Interval)
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-_ctx.Done():
				return
			case <-ticker.C:
				// Do something
				z.logger.Info("Doing something")

				realtimeDataResponse, err := z.scraper.Scrape(z.config.ThingKey)
				if err != nil {
					z.logger.Error("Error scraping zcsazzurro", zap.Error(err))
					continue
				}

				metrics, err := z.marshaler.UnmarshalMetrics(realtimeDataResponse)
				if err != nil {
					z.logger.Error("Error unmarshalling zcsazzurro metrics", zap.Error(err))
					continue
				}

				z.logger.Info("Metrics", zap.Any("metrics", metrics))
				z.consumer.ConsumeMetrics(_ctx, metrics)
			}
		}
	}()

	return nil
}

func (z *zcsazzurroReceiver) Shutdown(ctx context.Context) error {
	z.logger.Info("Shutting down zcsazzurro receiver")
	if z.cancel != nil {
		z.cancel()
	}

	return nil
}
