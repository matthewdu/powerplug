package postmates

import (
	"appengine"
	"appengine/urlfetch"
	"fmt"
	"net/http"
)

type CreateDeliveryRequest struct {
	manifest              string
	pickup_name           string
	pickup_address        string
	pickup_phone_number   string
	pickup_business_name  string
	pickup_notes          string
	dropoff_name          string
	dropoff_address       string
	dropoff_phone_number  string
	dropoff_business_name string
	dropoff_notes         string
}

const BASE_URL string = "https://api.postmates.com/"

func CreateDelivery(c appengine.Context, manifest string, pickup_name string, pickup_address string, pickup_phone_number string, pickup_business_name string, pickup_notes string, dropoff_name string, dropoff_address string, dropoff_phone_number string, dropoff_business_name string, dropoff_notes string) (*http.Response, error) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	url := BASE_URL + "/accounts/" + payer_id + "/transfers"
	req := CreateDeliveryRequest{
		manifest:              manifest,
		pickup_name:           pickup_name,
		pickup_address:        pickup_address,
		pickup_phone_number:   pickup_phone_number,
		pickup_business_name:  pickup_business_name,
		pickup_notes:          pickup_notes,
		dropoff_name:          dropoff_name,
		dropoff_address:       dropoff_address,
		dropoff_phone_number:  dropoff_phone_number,
		dropoff_business_name: dropoff_business_name,
		dropoff_notes:         dropoff_notes,
	}
	encoded, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	return client.Post(url, "application/json", bytes.NewBuffer(encoded))
}
