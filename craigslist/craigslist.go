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
	Url      string `json:"url"`
	Title    string `json:"title"`
	Price    int    `json:"price"`
	ImageUrl string `json:"imageUrl"`
}

func NewListing(ctx appengine.Context, url string) (*Listing, error) {
	client := urlfetch.Client(ctx)
	resp, err := client.Get("http://167.88.16.61:2138/" + url)
	if err != nil {
		ctx.Errorf("%s", err)
	}
	ctx.Debugf("Craigslist request came back with status: %s", resp.Status)
	if err != nil {
		ctx.Errorf("%s", err)
		return nil, errors.New("Get listing failed")
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		ctx.Errorf("%s", "Parsing Error")
		return nil, errors.New("Parse body failed")
	}

	title, ok := scrape.Find(root, scrape.ByTag(atom.Title))
	if !ok {
		ctx.Errorf("%s", "Error getting title")
		return nil, errors.New("Get title failed")
	}
	price, ok := scrape.Find(root, scrape.ByClass("price"))
	if !ok {
		ctx.Errorf("%s", "Error getting price")
		return nil, errors.New("Get price failed")
	}
	intPrice, err := strconv.Atoi(scrape.Text(price)[1:])
	if err != nil {
		ctx.Errorf("Error casting price: %s", scrape.Text(price))
		return nil, err
	}
	images := scrape.FindAll(root, scrape.ByTag(atom.Img))
	imageUrl := ""
	for _, image := range images {
		if scrape.Attr(image, "title") == "image 1" {
			imageUrl = scrape.Attr(image, "src")
		}
	}

	ctx.Debugf("Craigslist returned listing.Price: %d, listing.Title: %s", intPrice, scrape.Text(title))

	return &Listing{
		Url:      url,
		Title:    scrape.Text(title),
		Price:    intPrice,
		ImageUrl: imageUrl,
	}, nil
}
