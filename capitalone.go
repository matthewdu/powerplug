package capitalone

import (
	"fmt"
	"net/http"

	"appengine"
	"appengine/urlfetch"
)

type CreateTransferRequest struct {
	medium   string
	payee_id string
	amount   int
}

const BASE_URL string = "http://api.reimaginebanking.com/"

func CreateTransfer(c appengine.Context, payer_id string, payee_id string, amount int) (*http.Response, error) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	url := BASE_URL + "/accounts/" + payer_id + "/transfers"
	return client.Post(url, "application/json", bytes.NewBuffer(encoded))
}
