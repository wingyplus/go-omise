package omise

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAccount(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write([]byte(`{
			"object": "account",
			"id": "acct_4xhe277ad8n8zc8t9q9",
			"email": "abc@mail.com",
			"created": "2014-09-22T13:25:13Z"
		}`))
	}))
	defer ts.Close()

	var as = &AccountService{
		Key:    "skey_test_4xhd177bpqytcpk1w2a",
		client: newClient(ts.URL),
	}

	var (
		a   *Account
		err error
	)

	a, err = as.Retrieve()

	if err != nil {
		t.Error("expect error should be nil")
	}
	if a.Email != "abc@mail.com" {
		t.Error("expect abc@mail.com but got", a.Email)
	}
	if a.ID != "acct_4xhe277ad8n8zc8t9q9" {
		t.Error("expect acct_4xhe277ad8n8zc8t9q9 but got", a.ID)
	}
}
