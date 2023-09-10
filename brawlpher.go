package brawlpher

import (
	"net/http"

	"github.com/LuisPedroza/brawlpher/brawl/bs"
	"github.com/LuisPedroza/brawlpher/internal"
)

type Client struct {
	client internal.Requester
	apiKey string
	Brawl  *bs.Client
}

type Option func(*Client)

func WithClient(c internal.Requester) Option {
	return func(client *Client) {
		client.client = c
	}
}

func NewClient(apiKey string) *Client {
	c := &Client{
		client: http.DefaultClient,
		apiKey: apiKey,
	}

	baseClient := internal.NewClient(c.apiKey, c.client)
	c.Brawl = bs.NewClient(baseClient)
	return c
}
