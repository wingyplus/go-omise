package omise

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type OmiseError struct {
	Location string
	Code     string
	Message  string
}

func (e *OmiseError) Error() string { return "" }

type client struct {
	*http.Client
	url string
}

type response struct {
	*http.Response
}

func newClient(url string) *client {
	return &client{&http.Client{}, url}
}

func (r *response) decode(v interface{}) error {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, v)
	return err
}

func (c *client) do(method, key, uri string, body io.Reader) (*response, error) {
	req, err := http.NewRequest(method, c.url+uri, body)
	req.SetBasicAuth(key, "")
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	var response = &response{resp}

	if resp.StatusCode != http.StatusOK {
		var omiseError OmiseError

		response.decode(&omiseError)
		return nil, &omiseError
	}
	return response, nil
}

func (c *client) doGet(key, uri string) (*response, error) {
	return c.do("GET", key, uri, nil)
}

func (c *client) doPost(key, uri string, body io.Reader) (*response, error) {
	return c.do("POST", key, uri, body)
}
