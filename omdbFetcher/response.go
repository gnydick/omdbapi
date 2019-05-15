package omdbFetcher

import (
	"fmt"
	"github.com/mitchellh/go-wordwrap"
)

var prettyOutput = "Title: %s\n" +
	"Actors: %s\n" +
	"Director: %s\n" +
	"Writer: %s\n" +
	"Rotten Tomatoes Score: %s\n" +
	"Plot: %s"

type SV struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}

type OmdbResponse struct {
	Title     string `json:"Title"`
	Actors    string `json:"Actors"`
	Director  string `json:"Director"`
	Writer    string `json:"Writer"`
	Plot      string `json:"Plot"`
	Ratings   []SV
	leftovers map[string]interface{} `json:"-"`
}

func (or OmdbResponse) PrettyPrint() string {
	return fmt.Sprintf(prettyOutput, or.Title, or.Actors,
		or.Director, or.Writer, or.rottenTomatoes().Value, wordwrap.WrapString(or.Plot,80))
}

func (or OmdbResponse) rottenTomatoes() *SV {
	for _, v := range or.Ratings {
		if v.Source == "Rotten Tomatoes" {
			return &v
		}
	}
	return nil
}

