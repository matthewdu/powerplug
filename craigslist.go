package craigslist

import (
	"errors"
	"fmt"
	"net/http"

	"appengine"
	"appengine/urlfetch"

	"github.com/x/net/html/atom"
	"github.com/yhat/scrape"
)

type Listing struct {
	Url   string
	Title string
	Price int
	Area  string
}

func NewListing(ctx appengine.Context, url string) (*Listing, error) {
	client := urlfetch.Client(ctx)
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf(w, "Get Error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, errors.New("Get listing failed")
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf(w, "Parsing Error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, errors.New("Parse body failed")
	}

	title, ok := scrape.Find(root, scrape.ByClass("postingtitletext"))
	if !ok {
		fmt.Println("Error getting title")
		return nil, errors.New("Get title failed")
	}
	price, ok := scrape.Find(root, scrape.ByClass("price"))
	if !ok {
		fmt.Println("Error getting price")
		return nil, errors.New("Get price failed")
	}
	area, ok := scrape.Find(title, scrape.ByTag(atom.Small))
	if !ok {
		fmt.Println("Error getting area")
		return nil, errors.New("Get area failed")
	}

	return &Listing{
		Url:   url,
		Title: title,
		Price: price,
		Area:  area,
	}, nil
}
