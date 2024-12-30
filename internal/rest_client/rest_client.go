package rest

import (
	"net/http"
)

type Client struct {
	client *http.Client
	host   string
}

func NewClient(host string) *Client {
	return &Client{
		http.DefaultClient,
		host,
	}
}

func (c *Client) Get(urlPath string) (resource *string, err error) {
	resp, err := c.client.Get(c.host + urlPath)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, NotFound{urlPath}
	}

	return
}

// func DecodeResponse[T interface{}](resp *http.Response, resource string) (*T, error) {
// 	defer resp.Body.Close()

// 	var buf T
// 	return &buf, json.NewDecoder(resp.Body).Decode(&buf)
// }
