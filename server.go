package powerplug

import (
	"bytes"
	"capitalone"
	"craigslist"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"postmates"

	"appengine"
	"appengine/datastore"
	"appengine/mail"

	"github.com/go-zoo/bone"
)

type MainRequest struct {
	Buyer_email             string `json:"buyer_email" datastore:"buyer_email"`
	Cl_url                  string `json:"cl_url" datastore:"craigslist_url"`
	Cl_title                string `json:"-" datastore:"craigslist_title"`
	Cl_price                int    `json:"-" datastore:"craigslist_price"`
	Offer                   int    `json:"offer,string" datastore:"offer"`
	Cl_email                string `json:"cl_email" datastore:"craigslist_email"`
	Co_payer_id             string `json:"co_payer_id" datastore:"capital_one_payer_id"`
	Co_payee_id             string `json:"co_payee_id" datastore:"capital_one_payee_id"`
	Pm_dropoff_name         string `json:"pm_dropoff_name" datastore:"postmates_dropoff_name"`
	Pm_dropoff_address      string `json:"pm_dropoff_address" datastore:"postmates_dropoff_address"`
	Pm_dropoff_phone_number string `json:"pm_dropoff_phone_number" datastore:"postmates_dropoff_phone_number"`
	Pm_pickup_name          string `json:"pm_pickup_name" datastore:"postmates_pickup_name"`
	Pm_pickup_address       string `json:"pm_pickup_address" datastore:"postmates_pickup_address"`
	Pm_pickup_phone_number  string `json:"pm_pickup_phone_number" datastore:"postmates_pickup_phone_number"`
	Pm_delivery_id          string `json:"-" datastore:"postmates_delivery_id"`
}

type ReqEmail struct {
	BuyerName    string
	ListingTitle string
	Price        int
	AcceptUrl    string
}

type StatusEmail struct {
	BuyerName    string
	SellerName   string
	Price        int
	ListingTitle string
	StatusUrl    string
}

func init() {
	mux := bone.New()

	mux.PostFunc("/buy_request", BuyRequestHandler)
	mux.GetFunc("/request/:key", RequestHandler)
	mux.GetFunc("/delivery_status/:key", DeliveryHandler)
	mux.PostFunc("/accept_request/:key", AcceptRequestHandler)
	mux.PostFunc("/get_cl", GetCL)
	http.Handle("/", mux)
}

func createConfirmationURL(c appengine.Context, key *datastore.Key) string {
	if appengine.IsDevAppServer() {
		return "http://localhost:8080/request/" + key.Encode()
	}
	return "http://" + appengine.AppID(c) + ".appspot.com/request/" + key.Encode()
}

func createStatusURL(c appengine.Context, key *datastore.Key) string {
	if appengine.IsDevAppServer() {
		return "http://localhost:8080/status/" + key.Encode()
	}
	return "http://" + appengine.AppID(c) + ".appspot.com/status/" + key.Encode()
}

func GetCL(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.Errorf("%s", err)
	}
	bstr := string(b)
	if bstr == "" {
		return
	}
	listing, err := craigslist.NewListing(c, bstr)
	if err != nil {
		c.Errorf("%s", err)
	}

	err = json.NewEncoder(w).Encode(&listing)
	if err != nil {
		c.Errorf("%s", err)
	}
}

func BuyRequestHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	decoder := json.NewDecoder(r.Body)
	var request MainRequest
	err := decoder.Decode(&request)
	if err != nil {
		c.Errorf("Error decoding main request json")
	}
	listing, err := craigslist.NewListing(c, request.Cl_url)
	if err != nil {
		c.Errorf("%s", err)
	}
	request.Cl_title = listing.Title
	request.Cl_price = listing.Price

	key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "request", nil), &request)
	if err != nil {
		c.Errorf("Error putting purchase request into database: %s", err)
	}
	emailStruct := ReqEmail{
		BuyerName:    request.Pm_dropoff_name,
		ListingTitle: listing.Title,
		Price:        request.Offer,
		AcceptUrl:    createConfirmationURL(c, key),
	}
	c.Infof("Email url: %s", emailStruct.AcceptUrl)
	emailTemplate, err := template.ParseFiles("new_offer_email.html")
	if err != nil {
		c.Errorf("Error parsing template %s", err)
	}
	b := new(bytes.Buffer)
	emailTemplate.ExecuteTemplate(b, "email", emailStruct)
	msg := &mail.Message{
		Sender:   "craig-o-mation <craigomation@appspot.gserviceaccount.com>",
		To:       []string{request.Cl_email},
		Subject:  "Purchase Request for \"" + listing.Title + "\"",
		HTMLBody: string(b.Bytes()),
	}
	c.Debugf("Email body: %s", msg.Body)
	if err := mail.Send(c, msg); err != nil {
		c.Errorf("Couldn't send email: %v", err)
	}
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	keyString := bone.GetValue(r, "key")
	decodedKey, err := datastore.DecodeKey(keyString)
	if err != nil {
		c.Errorf("Error decoding key %s", err)
	}
	var purchase_request MainRequest
	if err = datastore.Get(c, decodedKey, &purchase_request); err != nil {
		c.Errorf("Error retrieving request key from database", err)
	}

	// render the template
	request_template, err := template.ParseFiles("request.html")
	if err != nil {
		c.Errorf("Error parsing template %s", err)
	}
	request_template.ExecuteTemplate(w, "request", purchase_request)
}

func AcceptRequestHandler(w http.ResponseWriter, r *http.Request) {
	keyId := bone.GetValue(r, "key")
	c := appengine.NewContext(r)
	key, err := datastore.DecodeKey(keyId)
	var dbRequest MainRequest
	if err = datastore.Get(c, key, &dbRequest); err != nil {
		c.Errorf("%s", err)
	}

	decoder := json.NewDecoder(r.Body)
	var request MainRequest
	if err := decoder.Decode(&request); err != nil {
		c.Errorf("%s", err)
	}

	dbRequest.Co_payee_id = request.Co_payee_id
	dbRequest.Pm_pickup_name = request.Pm_pickup_name
	dbRequest.Pm_pickup_address = request.Pm_pickup_address
	dbRequest.Pm_pickup_phone_number = request.Pm_pickup_phone_number

	co_resp, err := capitalone.CreateTransfer(c, dbRequest.Co_payer_id, dbRequest.Co_payee_id, dbRequest.Offer)
	if err != nil {
		c.Errorf("%s", err)
	}
	c.Debugf("%s", co_resp.Body)

	status, err := postmates.CreateDelivery(c, dbRequest.Cl_title, dbRequest.Pm_pickup_name, dbRequest.Pm_pickup_address, dbRequest.Pm_pickup_phone_number, "Craigslist", "", dbRequest.Pm_dropoff_name, dbRequest.Pm_dropoff_address, dbRequest.Pm_dropoff_phone_number, "Craigslist", "")
	if err != nil {
		c.Errorf("%s", err)
	}
	dbRequest.Pm_delivery_id = status.ID
	if _, err := datastore.Put(c, key, &dbRequest); err != nil {
		c.Errorf("%s", err)
	}

	// email both parties
	emailStruct := StatusEmail{
		BuyerName:    dbRequest.Pm_dropoff_name,
		SellerName:   dbRequest.Pm_pickup_name,
		Price:        dbRequest.Offer,
		ListingTitle: dbRequest.Cl_title,
		StatusUrl:    createStatusURL(c, key),
	}
	emailTemplate, err := template.ParseFiles("finished_transaction_email.html")
	if err != nil {
		c.Errorf("Error parsing template %s", err)
	}
	b := new(bytes.Buffer)
	emailTemplate.ExecuteTemplate(b, "email", emailStruct)
	msg := &mail.Message{
		Sender:   "craig-o-mation <craigomation@appspot.gserviceaccount.com>",
		To:       []string{dbRequest.Cl_email, dbRequest.Buyer_email},
		Subject:  dbRequest.Cl_title + " has been sold!",
		HTMLBody: string(b.Bytes()),
	}
	c.Debugf("Email body: %s", msg.Body)
	if err := mail.Send(c, msg); err != nil {
		c.Errorf("Couldn't send email: %v", err)
	}

	// send back important info
	err = json.NewEncoder(w).Encode(&status)
	if err != nil {
		c.Errorf("%s", err)
	}
}

func DeliveryHandler(w http.ResponseWriter, r *http.Request) {
	keyId := bone.GetValue(r, "key")
	c := appengine.NewContext(r)
	key, err := datastore.DecodeKey(keyId)
	var dbRequest MainRequest
	if err = datastore.Get(c, key, &dbRequest); err != nil {
		c.Errorf("%s", err)
	}

	var status *postmates.Status

	if dbRequest.Pm_delivery_id == "" {
		status, err = postmates.RescueDelivery(c)
		if status != nil {
			dbRequest.Pm_delivery_id = status.ID
			if _, err := datastore.Put(c, key, &dbRequest); err != nil {
				c.Errorf("%s", err)
			}
		}
	} else {
		status, err = postmates.GetStatus(c, dbRequest.Pm_delivery_id)
		if err != nil {
			c.Errorf("%s", err)
		}
	}

	// send back important info
	err = json.NewEncoder(w).Encode(&status)
	if err != nil {
		c.Errorf("%s", err)
	}
}
