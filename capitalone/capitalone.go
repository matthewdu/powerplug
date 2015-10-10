package capitalone

import (
	"appengine"
	"appengine/urlfetch"
	"bytes"
	"encoding/json"
	"net/http"
)

type CreateTransferRequest struct {
	medium   string
	payee_id string
	amount   int
}

const BASE_URL string = "http://api.reimaginebanking.com/"

func CreateTransfer(c appengine.Context, payer_id string, payee_id string, amount int) (*http.Response, error) {
	client := urlfetch.Client(c)
	url := BASE_URL + "/accounts/" + payer_id + "/transfers"
	req := CreateTransferRequest{
		medium:   "balance",
		payee_id: payee_id,
		amount:   amount,
	}
	encoded, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	return client.Post(url, "application/json", bytes.NewBuffer(encoded))
}
