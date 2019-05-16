package omdbFetcher

import (
	"fmt"
	"github.com/mitchellh/go-wordwrap"
	"log"
	"strconv"
	"strings"
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
		or.Director, or.Writer, or.rottenTomatoes().Value, wordwrap.WrapString(or.Plot, 80))
}

func (or OmdbResponse) PipeableOutput() string {
	return fmt.Sprintf("%s:%f", or.Title, rottenTomatoesFormatter(or.rottenTomatoes().Value))
}

func rottenTomatoesFormatter(score string) float64 {
	numeric := strings.Replace(score, "%", "", -1)
	value, _err := strconv.ParseFloat(numeric, 8)
	if _err != nil {
		log.Fatalln(fmt.Sprintf("Score is not a valid number: %s", numeric))
	}

	if value < 0 || value > 100 {
		log.Fatalln(fmt.Sprintf("Score is not a valid percentage: %d", numeric))
	}
	return value / 100
}

func (or OmdbResponse) rottenTomatoes() *SV {
	for _, v := range or.Ratings {
		if v.Source == "Rotten Tomatoes" {
			return &v
		}
	}
	return nil
}
