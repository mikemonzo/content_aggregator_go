package scraper

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Scraper struct {
	URLs []string
}

func NewScraper(urls []string) *Scraper {
	return &Scraper{URLs: urls}
}

func (s *Scraper) FetchContent() ([]string, error) {
	var content []string
	for _, url := range s.URLs {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			doc, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				return nil, err
			}
			text := doc.Text()
			content = append(content, text)
		} else {
			fmt.Printf("Error: %s returned status code %d \n", url, resp.StatusCode)
		}
	}
	return content, nil
}
