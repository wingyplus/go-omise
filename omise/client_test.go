package omise

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type testAccount struct {
	ID    string
	Email string
}

func TestClient(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/account" {
			t.Error("expect path is /account")
		}
		if r.Header.Get("Authorization") != "Basic dG9rbl90ZXN0XzEyMzQ1Njc6" {
			t.Error("expect basic is set", "tokn_test_1234567:")
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write([]byte(`{
			"object": "account",
			"id": "acct_4xhe277ad8n8zc8t9q9",
			"email": "abc@mail.com"
		}`))
	}))
	defer ts.Close()

	c := &client{
		Client: &http.Client{},
		url:    ts.URL,
	}
	resp, err := c.do("GET", "tokn_test_1234567", "/account", nil)

	if err != nil {
		t.Error(err)
	}

	var acc testAccount
	err = resp.decode(&acc)

	if err != nil {
		t.Error(err)
	}
	if acc.ID != "acct_4xhe277ad8n8zc8t9q9" {
		t.Error("expect account must be decode")
	}
}

func TestErrorClient(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write([]byte(`{
			"object": "error",
			"location": "https://docs.omise.co/api/errors#not-found",
			"code": "not_found",
			"message": "token tokn_test_1234567 was not found"
		}`))
	}))
	defer ts.Close()

	c := &client{
		Client: &http.Client{},
		url:    ts.URL,
	}

	_, err := c.do("GET", "tokn_test_1234567", "/account", nil)

	if err == nil {
		t.Error("expect error not to be nil")
	}

	if e, ok := err.(*OmiseError); ok {
		testOmiseError(t, e)
	} else {
		t.Error("expect err must be cast to OmiseError", err)
	}

	if err.Error() != "[not_found] token tokn_test_1234567 was not found" {
		t.Error("unexpect message", err.Error())
	}
}

func testOmiseError(t *testing.T, e *OmiseError) {
	if e.Location != "https://docs.omise.co/api/errors#not-found" {
		t.Errorf(
			"expect location https://docs.omise.co/api/errors#not-found but got %s",
			e.Location,
		)
	}
	if e.Code != "not_found" {
		t.Errorf(
			"expect code is not_found but got %s",
			e.Code,
		)
	}
	if e.Message != "token tokn_test_1234567 was not found" {
		t.Errorf(
			"unexpect token is %s",
			e.Message,
		)
	}
}
