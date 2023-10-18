# wikipedia_scraper
Wikipedia scraper written in Go

### Installation
To install the stats_cli package as an executable use:
    `go install github.com/Meowcenary/wikipedia_scraper@latest`
After installing commands can be run with:
    `wikipedia_scraper`

To clone the stats_cli package use:
    `git clone https://github.com/Meowcenary/wikipedia_scraper.git`
Once cloned commands can be run from the root directory:
    `go run main.go`

The file that will be written and the format that it will be written with is set
in main. By default it writes to "wikipages.jl". The flag "newlineDelim" on the
scraper function WriteWikiJson will format the written json such that each line
is a record structure. If set to false it will write the json as a compact
single line.

### Running the tests
To run all the tests use `go test ./...` from the root directory of the project.
Alternatively use `go test ./scraper`.

### Helpful extras
- [Colly documentation](https://go-colly.org/)
- [Colly Github](https://github.com/gocolly/colly/tree/master)
- [Colly API reference](https://pkg.go.dev/github.com/gocolly/colly?utm_source=godoc)
- Validate json output with [JSON lint](https://jsonlint.com/)
