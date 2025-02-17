package zcsazzurroreceiver

import (
	"go.uber.org/zap"

	"github.com/zmoog/zcs/azzurro"
)

func NewScraper(clientID, authKey, thingKey string, logger *zap.Logger) *Scraper {
	client := azzurro.NewClient(authKey, clientID)
	return &Scraper{
		client: client,
		logger: logger,
	}
}

type Scraper struct {
	logger *zap.Logger
	client *azzurro.Client
}

func (s *Scraper) Scrape(thingKey string) (azzurro.RealtimeDataResponse, error) {
	response, err := s.client.FetchRealtimeData(thingKey)
	if err != nil {
		return azzurro.RealtimeDataResponse{}, err
	}

	return response, nil
}
