package postmates

import (
	"appengine"
	"appengine/urlfetch"
	"net/http"
	"net/url"
	"strings"
)

type CreateDeliveryRequest struct {
	Manifest              string `json:"manifest"`
	Pickup_name           string `json:"pickup_name"`
	Pickup_address        string `json:"pickup_address"`
	Pickup_phone_number   string `json:"pickup_phone_number"`
	Pickup_business_name  string `json:"pickup_business_name"`
	Pickup_notes          string `json:"pickup_notes"`
	Dropoff_name          string `json:"dropoff_name"`
	Dropoff_address       string `json:"dropoff_address"`
	Dropoff_phone_number  string `json:"dropoff_phone_number"`
	Dropoff_business_name string `json:"dropoff_business_name"`
	Dropoff_notes         string `json:"dropoff_notes"`
}

const BASE_URL string = "https://api.postmates.com"
const CUSTOMER_ID string = "cus_KWZSpBpTC3PdsV"
const API_KEY string = "5d8dbe7d-1897-4239-b5b3-780c2a3965d5"

func CreateDelivery(c appengine.Context, manifest string, pickup_name string, pickup_address string, pickup_phone_number string, pickup_business_name string, pickup_notes string, dropoff_name string, dropoff_address string, dropoff_phone_number string, dropoff_business_name string, dropoff_notes string) (*http.Response, error) {
	client := urlfetch.Client(c)
	postUrl := BASE_URL + "/v1/customers/" + CUSTOMER_ID + "/deliveries"

	form := url.Values{}
	form.Add("manifest", manifest)
	form.Add("pickup_name", pickup_name)
	form.Add("pickup_address", pickup_address)
	form.Add("pickup_phone_number", pickup_phone_number)
	form.Add("pickup_business_name", pickup_business_name)
	form.Add("pickup_notes", pickup_notes)
	form.Add("dropoff_name", dropoff_name)
	form.Add("dropoff_address", dropoff_address)
	form.Add("dropoff_phone_number", dropoff_phone_number)
	form.Add("dropoff_business_name", dropoff_business_name)
	form.Add("dropoff_notes", dropoff_notes)

	req, err := http.NewRequest("POST", postUrl, strings.NewReader(form.Encode()))
	if err != nil {
		c.Errorf("%s", err)
	}
	req.PostForm = form
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(API_KEY, "")
	return client.Do(req)
}
