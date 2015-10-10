package powerplug

import (
	"appengine"
	"capitalone"
	// "craigslist"
	"encoding/json"
	"net/http"
	"postmates"

	"github.com/go-zoo/bone"
)

type MainRequest struct {
	Cl_url                  string `json:"cl_url"`
	Co_payer_id             string `json:"co_payer_id"`
	Co_payee_id             string `json:"co_payee_id"`
	Pm_pickup_name          string `json:"pm_pickup_name"`
	Pm_pickup_address       string `json:"pm_pickup_address"`
	Pm_pickup_phone_number  string `json:"pm_pickup_phone_number"`
	Pm_dropoff_name         string `json:"pm_dropoff_name"`
	Pm_dropoff_address      string `json:"pm_dropoff_address"`
	Pm_dropoff_phone_number string `json:"pm_dropoff_phone_number"`
}

func init() {
	mux := bone.New()

	// mux.Get, Post, etc ... takes http.Handler
	mux.Post("/get_my_stuff", http.HandlerFunc(RequestHandler))
	http.Handle("/", mux)
}

func RequestHandler(rw http.ResponseWriter, req *http.Request) {
	c := appengine.NewContext(req)
	decoder := json.NewDecoder(req.Body)
	var request MainRequest
	err := decoder.Decode(&request)
	if err != nil {
		panic(err)
	}
	// listing, err := craigslist.NewListing(c, request.Cl_url)
	// c.Debugf("%s", listing)
	// if err != nil {
	// 	panic(err)
	// }
	// co_resp, err := capitalone.CreateTransfer(c, request.Co_payer_id, request.Co_payee_id, listing.Price)
	co_resp, err := capitalone.CreateTransfer(c, request.Co_payer_id, request.Co_payee_id, 50)
	if err != nil {
		panic(err)
	}
	c.Debugf("%s", co_resp.Body)
	pm_resp, err := postmates.CreateDelivery(c, "Gum balls", request.Pm_pickup_name, request.Pm_pickup_address, request.Pm_pickup_phone_number, "Craigslist", "", request.Pm_dropoff_name, request.Pm_dropoff_address, request.Pm_dropoff_phone_number, "Craigslist", "")
	if err != nil {
		panic(err)
	}
	c.Debugf("%s", pm_resp.Body)
}
