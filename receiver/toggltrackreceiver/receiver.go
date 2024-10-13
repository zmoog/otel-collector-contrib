package toggltrackreceiver

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.uber.org/zap"
)

type togglTrackReceiver struct {
	cancel   context.CancelFunc
	logger   *zap.Logger
	consumer consumer.Logs
	config   *Config
}

func (t *togglTrackReceiver) Start(ctx context.Context, host component.Host) error {
	t.logger.Info("Starting toggltrack receiver")

	_ctx, cancel := context.WithCancel(ctx)
	t.cancel = cancel

	interval, _ := time.ParseDuration(t.config.Interval)
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-_ctx.Done():
				return
			case <-ticker.C:
				// Do something
				t.logger.Info("Doing something")
			}
		}
	}()

	return nil
}

func (t *togglTrackReceiver) Shutdown(ctx context.Context) error {
	t.logger.Info("Shutting down toggltrack receiver")
	if t.cancel != nil {
		t.cancel()
	}

	return nil
}
