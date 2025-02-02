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
	session        *toggl.Session
	logger         *zap.Logger
	lastScrapeTime time.Time
}

func (s *Scraper) Scrape(referenceTime time.Time, lookback time.Duration) ([]toggl.TimeEntry, error) {
	var endDate = referenceTime
	var startDate = endDate.Add(-lookback)

	// Get the time entries started between startDate and endDate.
	entries, err := s.session.GetTimeEntries(startDate, endDate)
	if err != nil {
		return nil, err
	}

	// We only want to send the entries
	// we haven't processed before.
	var newEntries []toggl.TimeEntry

	for _, entry := range entries {
		if entry.IsRunning() {
			// we only want to
			// consider completed
			// entries.
			continue
		}

		if entry.Stop.After(s.lastScrapeTime) {
			// add the entry to the new entries
			newEntries = append(newEntries, entry)
		}
	}

	s.lastScrapeTime = endDate

	return newEntries, nil
}
