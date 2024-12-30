package api

import (
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
	resp, err := c.restClient.R().SetResult(&model.Status{}).Get(c.host + "/status")
	if err != nil {
		return nil, err
	}

	return resp.Result().(*model.Status), nil
}
