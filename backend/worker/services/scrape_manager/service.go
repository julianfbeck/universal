package scrapemanager

import (
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"

	shareddomain "github.com/julianfbeck/universal/backend/services/stock/shared/shared_domain"
)

type ScrapeManager struct {
	collector *colly.Collector
}

func NewScrapeManager() *ScrapeManager {
	collector := colly.NewCollector(
		//Only allow whitelisted domains to be visited
		colly.Async(true),
	)
	collector.Limit(&colly.LimitRule{
		RandomDelay: 2 * time.Second,
	})
	extensions.RandomUserAgent(collector)

	return &ScrapeManager{
		collector: collector,
	}
}

func (sm *ScrapeManager) HandleScrapeTask(task shareddomain.ScraperTask) error {
	return nil
}
