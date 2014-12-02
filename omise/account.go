package omise

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type Account struct {
	ID      string    `json:"id"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
}

type AccountService struct {
	Key string
	URL string
}

func (as *AccountService) Retrieve() (*Account, error) {
	c := &http.Client{}
	req, err := http.NewRequest("GET", as.URL+"/account", nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(as.Key, "")

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	var acc Account
	err = json.Unmarshal(b, &acc)

	return &acc, err
}
