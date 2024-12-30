package api

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
	"github.com/milesjpool/kahd/internal/model"
)

type client struct {
	host       string
	restClient *resty.Client
}

func NewClient(host string) *client {
	rest_client := resty.New()
	return &client{host: host, restClient: rest_client}
}

func (c *client) Status() (*model.Status, error) {
	resp, err := c.restClient.R().Get(c.host + "/status")
	if err != nil {
		return nil, err
	}

	status := &model.Status{}
	if err := json.Unmarshal(resp.Body(), status); err != nil {
		return nil, err
	}
	return status, nil
}
