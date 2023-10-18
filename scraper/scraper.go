package scraper

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

// Struct to hold scraped data
type WikiPage struct {
	Url string `json:"Url"`
	Title string `json:"Title"`
	Text string `json:"Text"`
	Tags []string `json:"Tags"`
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

// Read json created from scraper
var pages []WikiPage
func ReadWikiJson(filepath string) {
	jsonFile, _ := os.Open(filepath)
	defer jsonFile.Close()
	dec := json.NewDecoder(jsonFile)
	dec.Decode(&pages)

	for _, page := range pages {
		fmt.Println(page.Text)
	}
}

func WriteWikiJson(filepath string, wikijson []WikiPage) {
	file, _ := os.Create(filepath)
	defer file.Close()
	enc := json.NewEncoder(file)
	// Using enc.Encode(wikijson) would write to single line, tihs separates records with newline chars
	for _, record := range wikijson {
		enc.Encode(record)
	}
}

// Pointer to WikiPage object, used to work around passing arguments to callbacks
var currentPage *WikiPage
func ScrapeWikiUrls(urls []string) []WikiPage {
	// Colly
	wikiPageCollector := colly.NewCollector()
	wikiPageCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	wikiPageCollector.OnHTML("#firstHeading", scrapeElement)
	// "#mw-content-text" is taken from Python script "//div[@id="mw-content-text"]"
	wikiPageCollector.OnHTML("#mw-content-text", scrapeElement)

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

func scrapeElement(e *colly.HTMLElement) {
	currentPage.Title = fmt.Sprintf("%s", e.Text)
}
