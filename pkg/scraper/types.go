package scraper

import (
	"sync"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

type Scraper struct {
	colly *colly.Collector
	q     *queue.Queue

	mutex       *sync.Mutex
	urlBackoffs map[string]int
}

type Product struct {
	URL        string
	Name       string
	Price      string
	Savings    string
	Percentage string
	Image      string
}
