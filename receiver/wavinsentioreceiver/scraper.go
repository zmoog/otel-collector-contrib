package wavinsentioreceiver

import (
	"github.com/zmoog/ws/ws"
	"github.com/zmoog/ws/ws/identity"
	"go.uber.org/zap"
)

func newScraper(username, password string, logger *zap.Logger) *Scraper {
	identityManager := identity.NewManager(username, password)
	client := ws.NewClient(identityManager, "https://wavin-api.jablotron.cloud")

	return &Scraper{
		logger: logger,
		client: client,
	}
}

type Scraper struct {
	logger *zap.Logger
	client *ws.Client
}

func (s *Scraper) Scrape() ([]ScrapeResult, error) {
	results := []ScrapeResult{}

	locations, err := s.client.ListLocations()
	if err != nil {
		return nil, err
	}

	for _, location := range locations {
		rooms, err := s.client.ListRooms(location.Ulc)
		if err != nil {
			return nil, err
		}

		results = append(results, ScrapeResult{
			Location: location,
			Rooms:    rooms,
		})
	}

	return results, nil
}

type ScrapeResult struct {
	Location ws.Location
	Rooms    []ws.Room
}
