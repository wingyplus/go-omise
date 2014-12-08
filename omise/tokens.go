package omise

import (
	"net/url"
	"strconv"
	"strings"
)

type CardInfo struct {
	Name            string
	Number          string
	ExpirationMonth int
	ExpirationYear  int
	City            string
	PostalCode      string
	SecurityCode    int
}

type Card struct {
	ID                string `json:"id"`
	LiveMode          bool   `json:"livemode"`
	Country           string `json:"country"`
	City              string `json:"city"`
	PostalCode        string `json:"postal_code"`
	Financing         string `json:"financing"`
	LastDigits        string `json:"last_digits"`
	Brand             string `json:"brand"`
	ExpirationMonth   int    `json:"expiration_month"`
	ExpirationYear    int    `json:"expiration_year"`
	Fingerprint       string `json:"fingerprint"`
	Name              string `json:"name"`
	SecurityCodeCheck bool   `json:"security_code_check"`
}

type Token struct {
	ID       string `json:"id"`
	LiveMode bool   `json:"livemode"`
	Used     bool   `json:"used"`
	Location string `json:"location"`
	Card     *Card  `json:"card"`
}

type TokensService struct {
	Key    string
	client *client
}

func (ts *TokensService) Create(ci *CardInfo) (*Token, error) {
	var data = url.Values{}
	data.Set("card[name]", ci.Name)
	data.Set("card[number]", ci.Number)
	data.Set("card[expiration_month]", strconv.Itoa(ci.ExpirationMonth))
	data.Set("card[expiration_year]", strconv.Itoa(ci.ExpirationYear))
	data.Set("card[city]", ci.City)
	data.Set("card[postal_code]", ci.PostalCode)
	data.Set("card[security_code]", strconv.Itoa(ci.SecurityCode))

	resp, err := ts.client.doPost(ts.Key, "/tokens", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	var t Token
	err = resp.decode(&t)

	return &t, err
}

func (ts *TokensService) Get(key string) (*Token, error) {
	resp, err := ts.client.doGet(ts.Key, "/tokens/"+key)
	if err != nil {
		return nil, err
	}
	var t Token
	err = resp.decode(&t)
	return &t, err
}
