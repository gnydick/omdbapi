package omdbFetcher

import (
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"log"
)

var endpoint = "http://www.omdbapi.com/?apikey=%s&type=movie&t=%s&plot=short"

type OmdbFetcher struct {
	apiKey *string
	title  *string
}

func NewFetcher(apiKey *string, title *string) *OmdbFetcher {
	o := new(OmdbFetcher)
	o.apiKey = apiKey
	o.title = title
	return o
}

func (o OmdbFetcher) Fetch() *OmdbResponse {
	url := fmt.Sprintf(endpoint, *o.apiKey, *o.title)
	resp, _err := grequests.Get(url, nil)
	if _err != nil {
		log.Fatalln(fmt.Sprintf("Unable to make request: %s", _err.Error()))
	}
	or := OmdbResponse{}
	_err = json.Unmarshal(resp.Bytes(), &or)
	if _err != nil {
		log.Fatalln("Couldn't unpack response: %s", _err.Error())
	}

	return &or
}
