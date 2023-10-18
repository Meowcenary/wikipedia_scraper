package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

// Struct to hold scraped data
type WikiPage struct {
	Url string
	Title string
	Text string
	Tags []string
}

// Constructor
func NewWikiPage(url string) *WikiPage {
	return &WikiPage{
		Url: url,
		Title: "",
		Text: "",
		Tags: []string{},
	}
}

func scrapeTitle(e *colly.HTMLElement) {
	currentPage.Title = fmt.Sprintf("%s", e.Text)
}

func scrapeText(e *colly.HTMLElement) {
	currentPage.Text = fmt.Sprintf("%s", e.Text)
}

func scrapeWikiUrls(urls []string, wikiPageCollector *colly.Collector) []WikiPage {
	json := []WikiPage{}

	for _, url := range urls {
		// Assign currentPage pointer to a new WikiPage struct
		currentPage = NewWikiPage(url)
		// Visit url and trigger scraping functions
		wikiPageCollector.Visit(url)
		// Dereference and append WikiPage struct
		json = append(json, *currentPage)
	}

	return json
}

// Pointer to WikiPage object, used to work around passing arguments to callbacks
var currentPage *WikiPage
func main() {
	// Urls to scrape
	urls := []string{
		"https://en.wikipedia.org/wiki/Robotics",
	}

	// Colly
	wikiPageCollector := colly.NewCollector()
	wikiPageCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	wikiPageCollector.OnHTML("#firstHeading", scrapeTitle)
	// #mw-content-text is taken from Python script "//div[@id="mw-content-text"]"
	wikiPageCollector.OnHTML("#mw-content-text", scrapeText)
	// wikiPageCollector is passed by reference to prevent instantiating it for each url
	wikijson := scrapeWikiUrls(urls, wikiPageCollector)

	// Encode to json and write to file
	filepath := "wikipages.json"
	file, _ := os.Create(filepath)
	defer file.Close()
	enc := json.NewEncoder(file)
	enc.Encode(wikijson)
}
