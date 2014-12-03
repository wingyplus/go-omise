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

type NotFoundError struct {
	Location string `json:"location"`
	Code string `json:"code"`
	Message string `json:"message"`
}

func (err *NotFoundError) Error() string {
	return ""
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

func (ts *TokensService) Get(key string) (*Token, error) {
	req, _ := http.NewRequest("GET", ts.URL+"/tokens/"+key, nil)
	req.SetBasicAuth(ts.Key, "")
	c := &http.Client{}
	resp, _ := c.Do(req)
	b, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var (
		t Token
		e NotFoundError
	)

	if resp.StatusCode == http.StatusNotFound {
		json.Unmarshal(b, &e)
		return nil, &e
	}
	err := json.Unmarshal(b, &t)
	return &t, err
}
