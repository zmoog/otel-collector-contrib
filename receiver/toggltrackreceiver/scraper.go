package toggltrackreceiver

import (
	"time"

	toggl "github.com/jason0x43/go-toggl"
	"go.uber.org/zap"
)

func NewScraper(apiToken string, logger *zap.Logger) *Scraper {
	session := toggl.OpenSession(apiToken)

	return &Scraper{
		session: &session,
		logger:  logger,
	}
}

type Scraper struct {
	session *toggl.Session
	logger  *zap.Logger
}

func (s *Scraper) Scrape(referenceTime time.Time) ([]toggl.TimeEntry, error) {
	endDate := referenceTime                           // time.Parse("2006-01-02", "2024-10-05")
	startDate := endDate.Add(-1 * 24 * 30 * time.Hour) // time.Parse("2006-01-02", "2024-10-04")

	entries, err := s.session.GetTimeEntries(startDate, endDate)
	if err != nil {
		return nil, err
	}

	return entries, nil
}
