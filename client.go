package crypt0cloud_sdk_go

import (
	"bytes"
	"github.com/onlyangel/apihandlers"
	"io/ioutil"
	"net/http"
)

type Crypt0Client struct {
	Endpoint string
	Client   *http.Client
}

func GetClient(endpoint_url string) *Crypt0Client {
	return &Crypt0Client{
		Endpoint: endpoint_url,
		Client:   &http.Client{},
	}
}

func (c Crypt0Client) _get(url string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	apihandlers.PanicIfNotNil(err)

	resp, err := c.Client.Do(req)
	apihandlers.PanicIfNotNil(err)

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	return body
}

func (c Crypt0Client) _post(url string, data []byte) []byte {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	apihandlers.PanicIfNotNil(err)

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	apihandlers.PanicIfNotNil(err)

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	return body
}
