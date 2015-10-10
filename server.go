package powerplug

import (
	"capitalone"
	"craigslist"
	"fmt"
	"go-zoo/bone"
	"net/http"
	"postmates"
)

type MainRequest struct {
	cl_url                  string
	co_payer_id             string
	co_payee_id             string
	pm_pickup_name          string
	pm_pickup_address       string
	pm_pickup_phone_number  string
	pm_dropoff_name         string
	pm_dropoff_address      string
	pm_dropoff_phone_number string
}

func init() {
	mux := bone.New()

	// mux.Get, Post, etc ... takes http.Handler
	mux.Post("/get_my_stuff", http.HandlerFunc(RequestHandler))

	http.ListenAndServe(":8080", mux)
}

func RequestHandler(rw http.ResponseWriter, req *http.Request) {
	c := appengine.NewContext(req)
	decoder := json.NewDecoder(req.Body)
	var request MainRequest
	err := decoder.Decode(&request)
	if err != nil {
		panic()
	}
	listing, err := craigslist.NewListing(c, request.cl_url)
	if err != nil {
		panic()
	}
	co_resp, err := capitalone.CreateTransfer(request.co_payer_id, request.co_payee_id, amount)
	if err != nil {
		panic()
	}
	pm_resp, err := postmates.CreateDelivery(c, title, request.pm_pickup_name, request.pm_pickup_address, request.pm_pickup_phone_number, "Craigslist", "", request.pm_dropoff_name, request.pm_dropoff_address, request.pm_dropoff_phone_number, "Craigslist", "")
	if err != nil {
		panic()
	}
}
