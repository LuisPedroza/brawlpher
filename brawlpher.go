package brawlpher

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/LuisPedroza/brawlpher/brawl/bs"
	"github.com/LuisPedroza/brawlpher/internal"
)

type Client struct {
	client internal.Requester
	apiKey string
	logger slog.Logger
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
		logger: *slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}

	baseClient := internal.NewClient(c.apiKey, c.client, c.logger)
	c.Brawl = bs.NewClient(baseClient)

	c.logger.Info("Client created", slog.String("package", "brawlpher"), slog.String("function", "NewClient"))

	return c
}
