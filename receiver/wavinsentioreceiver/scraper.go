package wavinsentioreceiver

import (
	"go.uber.org/zap"

	"github.com/zmoog/ws/ws"
	"github.com/zmoog/ws/ws/identity"
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

// Scrape scrapes the data from the wavinsentio API.
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

// ScrapeResult is the result of a scrape.
// It contains a location and all rooms in that location.
type ScrapeResult struct {
	Location ws.Location
	Rooms    []ws.Room
}
