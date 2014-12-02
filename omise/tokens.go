package omise

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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

type Token struct {
	ID       string `json:"id"`
	LiveMode bool   `json:"livemode"`
}

type TokensService struct {
	Key string
	URL string
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

	req, _ := http.NewRequest("POST", ts.URL+"/tokens", strings.NewReader(data.Encode()))
	req.SetBasicAuth(ts.Key, "")

	c := &http.Client{}

	resp, _ := c.Do(req)

	b, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var t Token
	err := json.Unmarshal(b, &t)

	return &t, err
}
