package postmates

import (
	"appengine"
	"appengine/urlfetch"
	"encoding/json"
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

type Status struct {
	ID       string `json:"id"`
	Status   string `json:"status"`
	Complete bool   `json:"complete"`
	Courier  struct {
		Location struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
		ImgLink string `json:"img_href"`
	} `json:"courier"`
	Pickup struct {
		Location struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
	} `json:"pickup"`
	Dropoff struct {
		Location struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
	} `json:"dropoff"`
	EndAddress string `json:"endAddress"`
}

type Rescue struct {
	Data        []Status `json:"data"`
	Total_count int      `json:"total_count"`
}

const BASE_URL string = "https://api.postmates.com"
const CUSTOMER_ID string = "cus_KWZSpBpTC3PdsV"
const API_KEY string = "5d8dbe7d-1897-4239-b5b3-780c2a3965d5"

func CreateDelivery(c appengine.Context, manifest string, pickup_name string, pickup_address string, pickup_phone_number string, pickup_business_name string, pickup_notes string, dropoff_name string, dropoff_address string, dropoff_phone_number string, dropoff_business_name string, dropoff_notes string) (*Status, error) {
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
	resp, err := client.Do(req)
	decoder := json.NewDecoder(resp.Body)
	var status Status
	err = decoder.Decode(&status)
	if status.ID == "" {
		c.Errorf("Creating a delivery was unsuccessful", nil)
		c.Errorf("%s", resp.Body)
	}
	if err != nil {
		return nil, err
	}
	return &status, nil
}

func GetStatus(c appengine.Context, delivery_id string) (*Status, error) {
	client := urlfetch.Client(c)
	url := BASE_URL + "/v1/customers/" + CUSTOMER_ID + "/deliveries/" + delivery_id
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(API_KEY, "")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(resp.Body)
	var status Status
	err = decoder.Decode(&status)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

func RescueDelivery(c appengine.Context) (*Status, error) {
	c.Errorf("Attempting a rescue", nil)
	client := urlfetch.Client(c)
	url := BASE_URL + "/v1/customers/" + CUSTOMER_ID + "/deliveries?filter=ongoing"
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(API_KEY, "")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(resp.Body)
	var rescue Rescue
	err = decoder.Decode(&rescue)
	if err != nil {
		return nil, err
	}
	if rescue.Total_count == 1 {
		c.Errorf("Rescue successful!", nil)
		return &rescue.Data[0], nil
	}
	return nil, nil
}
