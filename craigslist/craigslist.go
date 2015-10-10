package craigslist

import (
	"errors"
	"strconv"

	"appengine"
	"appengine/urlfetch"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Listing struct {
	Url   string
	Title string
	Price int
	Area  string
}

func NewListing(ctx appengine.Context, url string) (*Listing, error) {
	client := urlfetch.Client(ctx)
	resp, err := client.Get("http://sfbay.craigslist.org")
	ctx.Errorf("%s", resp.Status)
	if err != nil {
		ctx.Errorf("%s", err)
		return nil, errors.New("Get listing failed")
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		ctx.Errorf("%s", "Parsing Error")
		return nil, errors.New("Parse body failed")
	}

	title, ok := scrape.Find(root, scrape.ByClass("postingtitletext"))
	if !ok {
		ctx.Errorf("%s", "Error getting title")
		return nil, errors.New("Get title failed")
	}
	price, ok := scrape.Find(root, scrape.ByClass("price"))
	if !ok {
		ctx.Errorf("%s", "Error getting price")
		return nil, errors.New("Get price failed")
	}
	intPrice, err := strconv.Atoi(scrape.Text(price))
	if err != nil {
		return nil, err
	}

	area, ok := scrape.Find(title, scrape.ByTag(atom.Small))
	if !ok {
		ctx.Errorf("%s", "Error getting area")
		return nil, errors.New("Get area failed")
	}

	return &Listing{
		Url:   url,
		Title: scrape.Text(title),
		Price: intPrice,
		Area:  scrape.Text(area),
	}, nil
}
