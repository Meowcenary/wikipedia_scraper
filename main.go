package main

import (
	"log"

	"github.com/Meowcenary/wikipedia_scraper/scraper"
)

func main() {
	log.Println("Starting...")
	urls := []string{
		"https://en.wikipedia.org/wiki/Robotics",
		"https://en.wikipedia.org/wiki/Robot",
		"https://en.wikipedia.org/wiki/Reinforcement_learning",
    "https://en.wikipedia.org/wiki/Robot_Operating_System",
    "https://en.wikipedia.org/wiki/Intelligent_agent",
    "https://en.wikipedia.org/wiki/Software_agent",
    "https://en.wikipedia.org/wiki/Robotic_process_automation",
    "https://en.wikipedia.org/wiki/Chatbot",
    "https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
    "https://en.wikipedia.org/wiki/Android_(robot)",
	}
	wikijson := scraper.ScrapeWikiUrls(urls)
	filepath := "wikipages.jl"
	newlineDelim := true
	log.Printf("Done scraping, writing to file %s...\n", filepath)
	err := scraper.WriteWikiJson(filepath, wikijson, newlineDelim)

	if err != nil {
		log.Fatalf("Error writing to file: %s", err)
	}
	log.Println("Done writing to file!")
}
