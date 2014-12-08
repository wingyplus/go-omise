package omise

import (
	"time"
)

type Account struct {
	ID      string    `json:"id"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
}

type AccountService struct {
	Key    string
	client *client
}

func (as *AccountService) Retrieve() (*Account, error) {
	resp, err := as.client.doGet(as.Key, "/account")
	if err != nil {
		return nil, err
	}

	var acc Account
	err = resp.decode(&acc)

	return &acc, err
}
