package omise

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateToken(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write([]byte(`{
			"object": "token",
			"id": "tokn_test_4y96o5lnx6m7fw8wpg9",
			"livemode": false,
			"location": "/tokens/tokn_test_4y96o5lnx6m7fw8wpg9",
			"used": false,
			"card": {
				"object": "card",
				"id": "card_test_4y96o5lmaos8ulnea6y",
				"livemode": false,
				"country": "us",
				"city": "Bangkok",
				"postal_code": "10320",
				"financing": "",
				"last_digits": "4242",
				"brand": "Visa",
				"expiration_month": 10,
				"expiration_year": 2018,
				"fingerprint": "dmCDUHPNUyfWPtkas7mm/IMBA7oYMEJ3B9SK3kMDzQQ=",
				"name": "Somchai Prasert",
				"security_code_check": true,
				"created": "2014-12-02T16:39:55Z"
			},
			"created": "2014-12-02T16:39:55Z"
		}`))
	}))

	var tks = TokensService{
		Key: "pkey_test_4xhd177bnqcnz8lqp7c",
		URL: ts.URL,
	}

	var (
		token *Token
		err   error
	)

	token, err = tks.Create(&CardInfo{
		Name:            "Somchai Prasert",
		Number:          "4242424242424242",
		ExpirationMonth: 10,
		ExpirationYear:  2018,
		City:            "Bangkok",
		PostalCode:      "10320",
		SecurityCode:    123,
	})

	if err != nil {
		t.Error("expect error should be nil")
	}
	testToken(t, token)
}

func testToken(t *testing.T, token *Token) {
	if token.ID != "tokn_test_4y96o5lnx6m7fw8wpg9" {
		t.Errorf("expect tokn_test_4y96o5lnx6m7fw8wpg9 but got %s", token.ID)
	}
	if token.LiveMode != false {
		t.Error("expect live mode to be false")
	}
	if token.Used != false {
		t.Error("expect used to be false")
	}
	if token.Location != "/tokens/tokn_test_4y96o5lnx6m7fw8wpg9" {
		t.Errorf("expect location %s but got %s", "/tokens/tokn_test_4y96o5lnx6m7fw8wpg9", token.Location)
	}

	testCard(t, token.Card)
}

func testCard(t *testing.T, c *Card) {
	if c == nil {
		t.Error("expect card not to be nil")
	}
	/*
		"id": "card_test_4y96o5lmaos8ulnea6y",
		"livemode": false,
		"country": "us",
		"city": "Bangkok",
		"postal_code": "10320",
		"financing": "",
		"last_digits": "4242",
		"brand": "Visa",
		"expiration_month": 10,
		"expiration_year": 2018,
		"fingerprint": "dmCDUHPNUyfWPtkas7mm/IMBA7oYMEJ3B9SK3kMDzQQ=",
		"name": "Somchai Prasert",
		"security_code_check": true,
	*/
	if c.ID != "card_test_4y96o5lmaos8ulnea6y" {
		t.Errorf("expect %s but got %s", "card_test_4y96o5lmaos8ulnea6y", c.ID)
	}
	if c.LiveMode != false {
		t.Error("expect livemode to be false")
	}
	if c.Country != "us" {
		t.Errorf("expect country is us but got %s", c.Country)
	}
}
