package bs

import "github.com/LuisPedroza/brawlpher/internal"

type Client struct {
	Brawler *BrawlerClient
}

func NewClient(base *internal.Client) *Client {
	return &Client{
		Brawler: &BrawlerClient{c: base},
	}
}
