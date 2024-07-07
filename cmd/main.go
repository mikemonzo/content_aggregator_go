package main

import (
	"content_aggregator_go/internal/database"
	"content_aggregator_go/internal/scraper"
	"fmt"
	"log"
)

func main() {
	urls := []string{
		"https://www.tecsidel.com",
		"https://eysaservicios.com",
	}

	scraper := scraper.NewScraper(urls)
	content, err := scraper.FetchContent()
	if err != nil {
		log.Fatalf("Error fetching content: %v", err)
	}

	db, err := database.NewDatabase("content.db")
	if err != nil {
		log.Fatalf("Error creating database: %v", err)
	}
	defer db.Close()

	for _, data := range content {
		if err := db.InsertContent(data); err != nil {
			log.Fatalf("Error inserting content: %v", err)
		}
	}

	allContent, err := db.FetchAllContent()
	if err != nil {
		log.Fatalf("Error fetching all content: %v", err)
	}

	for _, data := range allContent {
		fmt.Println(data)
	}
}
