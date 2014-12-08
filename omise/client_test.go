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
		if r.Header.Get("Authorization") != "Basic dG9rbl90ZXN0XzR5OTZvNWxueDZtN2Z3OHdwZzk6" {
			t.Error("expect basic is set", "tokn_test_4y96o5lnx6m7fw8wpg9:")
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
	resp, err := c.do("GET", "tokn_test_4y96o5lnx6m7fw8wpg9", "/account", nil)

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
