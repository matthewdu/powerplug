package capitalone

import (
	"appengine"
	"appengine/urlfetch"
	"bytes"
	"encoding/json"
	"net/http"
)

type CreateTransferRequest struct {
	Medium   string `json:"medium"`
	Payee_id string `json:"payee_id"`
	Amount   int    `json:"amount"`
}

const BASE_URL string = "http://api.reimaginebanking.com/"
const API_KEY string = "156d8519f2dde3a7bc87dbe66a2c11f8"

func CreateTransfer(c appengine.Context, payer_id string, payee_id string, amount int) (*http.Response, error) {
	client := urlfetch.Client(c)
	url := BASE_URL + "/accounts/" + payer_id + "/transfers?key=" + API_KEY
	req := CreateTransferRequest{
		Medium:   "balance",
		Payee_id: payee_id,
		Amount:   amount,
	}
	encoded, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	return client.Post(url, "application/json", bytes.NewBuffer(encoded))
}
