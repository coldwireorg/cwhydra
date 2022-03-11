package cwhydra

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Api interface {
	Get(endpoint string, response interface{}) error
	Post(endpoint string, body interface{}, response interface{}) error
	Put(endpoint string, body interface{}, response interface{}) error
	Del(endpoint string, body interface{}, response interface{}) error
	Patch(endpoint string, body interface{}, response interface{}) error
	request(method string, endpoint string, body interface{}, response interface{}) error
}

type Error struct {
	Error            string
	ErrorDebug       string
	ErrorDescription string
	StatusCode       string
}

type Config struct {
	Url  string
	Http http.Client
}

func New(c Config) Api {
	return c
}

func (c Config) Get(endpoint string, response interface{}) error {
	return c.request("GET", endpoint, nil, response)
}

func (c Config) Post(endpoint string, body interface{}, response interface{}) error {
	return c.request("POST", endpoint, body, response)
}

func (c Config) Put(endpoint string, body interface{}, response interface{}) error {
	return c.request("PUT", endpoint, body, response)
}

func (c Config) Del(endpoint string, body interface{}, response interface{}) error {
	return c.request("DELETE", endpoint, body, response)
}

func (c Config) Patch(endpoint string, body interface{}, response interface{}) error {
	return c.request("PATCH", endpoint, body, response)
}

// Send request
func (c Config) request(method string, endpoint string, body interface{}, response interface{}) error {
	var bdy []byte

	// parse url from config
	uri, err := url.Parse(c.Url)
	if err != nil {
		return err
	}

	// if there is a body convert is to json
	if body != nil {
		bdy, err = json.Marshal(body)
		if err != nil {
			return err
		}
	}

	// create request
	req, err := http.NewRequest(method, uri.String()+endpoint, bytes.NewBuffer(bdy))
	if err != nil {
		return err
	}

	// set header
	if body != nil {
		req.Header["Content-Type"] = []string{"application/json"}
	}

	// send request
	res, err := c.Http.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	resBodyEncoded, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// if response code is OK
	if res.StatusCode == 201 || res.StatusCode == 200 || res.StatusCode == 204 {
		if response != nil {
			err = json.Unmarshal(resBodyEncoded, &response)
			if err != nil {
				return err
			}
		}
	} else {
		var e Error
		err = json.Unmarshal(resBodyEncoded, &e)
		if err != nil {
			return err
		}

		if e.Error != "" {
			return errors.New(e.Error)
		}
	}

	return nil
}
