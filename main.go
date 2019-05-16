package main

import (
	"flag"
	"fmt"
	. "github.com/gnydick/omdbapi/omdbFetcher"
	"os"
)

func main() {
	fetcher := setup()
	or := fetcher.Fetch()
	fmt.Println(or.PipeableOutput())
}

func setup() *OmdbFetcher {
	title := flag.String("title", "", "Movie title.")
	flag.Parse()
	apiKey := os.Getenv("API_KEY")
	of := NewFetcher(&apiKey, title)
	return of
}
