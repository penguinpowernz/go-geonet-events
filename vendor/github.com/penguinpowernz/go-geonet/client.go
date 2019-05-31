package geonet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const defaultVersion = "2"
const baseURI = "https://api.geonet.org.nz"

var sfmt = fmt.Sprintf

type Client struct {
	*http.Client
	Version string
}

func NewClient() *Client {
	return &Client{&http.Client{}, defaultVersion}
}

func (c *Client) Request(method, path string) (*http.Request, error) {
	url := fmt.Sprintf("%s/%s", baseURI, strings.TrimLeft(path, "/"))
	req, err := http.NewRequest(method, url, nil)
	req.Header.Set("Accept", "application/vnd.geo+json;version="+c.Version)
	return req, err
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	log.Printf("requesting %s", req.URL)
	return c.Client.Do(req)
}

func (c *Client) Get(path string, v interface{}) error {
	req, err := c.Request("get", path)
	if err != nil {
		return err
	}

	res, err := c.Do(req)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(res.Body)
	// log.Println(string(data))
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}
