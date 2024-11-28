package api

import (
	"net/http"
)

type Client struct {
	URL string
}

func (c *Client) Get(resource string) (*http.Response, error) {
	httpClient := &http.Client{}
	resp, err := httpClient.Get(c.URL + "/" + resource)
	if err != nil {
		// Handle error
	}
	// defer resp.Body.Close()
	return resp, err
}
