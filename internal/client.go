package internal

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/LuisPedroza/brawlpher/api"
)

const (
	apiURLFormat              = "%s://%s%s"
	baseURL                   = "bsproxy.royaleapi.dev/v1"
	scheme                    = "https"
	apiTokenHeaderKey         = "Authorization"
	apiTokenHeaderValueFormat = "Bearer %s"
)

type Client struct {
	APIKey string
	Client Requester
	Logger slog.Logger
}

func NewClient(apiKey string, client Requester, logger slog.Logger) *Client {
	return &Client{
		APIKey: apiKey,
		Client: client,
		Logger: logger,
	}
}

func (c *Client) NewRequest(method, endpoint string, pagination *api.PaginationQueryItems) (*http.Request, error) {
	request, err := http.NewRequest(method, fmt.Sprintf(apiURLFormat, scheme, baseURL, endpoint), nil)
	if err != nil {
		c.Logger.Debug("NewRequest",
			slog.String("package", "internal"),
			slog.String("type", "Client"),
			slog.String("endpoint", endpoint),
			slog.String("error", err.Error()))
		return nil, err
	}

	request.Header.Add(apiTokenHeaderKey, fmt.Sprintf(apiTokenHeaderValueFormat, c.APIKey))
	request.Header.Add("Accept", "application/json")

	if pagination != nil {
		q := request.URL.Query()
		if pagination.After != "" {
			q.Add("after", pagination.After)
		}
		if pagination.Before != "" {
			q.Add("before", pagination.Before)
		}
		if pagination.Limit != 0 {
			q.Add("limit", fmt.Sprintf("%d", pagination.Limit))
		}
		request.URL.RawQuery = q.Encode()
	}

	return request, nil

}

func (c *Client) DoRequest(method, endpoint string, pagination *api.PaginationQueryItems) (*http.Response, error) {
	request, err := c.NewRequest(method, endpoint, pagination)
	if err != nil {
		c.Logger.Debug("DoRequest",
			slog.String("package", "internal"),
			slog.String("type", "Client"),
			slog.String("endpoint", endpoint),
			slog.String("error", err.Error()))
		return nil, err
	}

	response, err := c.Client.Do(request)
	if err != nil {
		c.Logger.Debug("DoRequest",
			slog.String("package", "internal"),
			slog.String("type", "Client"),
			slog.String("endpoint", endpoint),
			slog.String("error", err.Error()))
		return nil, err
	}

	if response.StatusCode == http.StatusServiceUnavailable {
		time.Sleep(time.Second)
		response, err = c.Client.Do(request)
		if err != nil {
			c.Logger.Debug("DoRequest",
				slog.String("package", "internal"),
				slog.String("type", "Client"),
				slog.String("endpoint", endpoint),
				slog.String("error", err.Error()))
			return nil, err
		}
	}
	if response.StatusCode < 200 || response.StatusCode > 299 {
		err, ok := api.StatusToError[response.StatusCode]
		if !ok {
			err = api.Error{Message: "unknown error", StatusCode: response.StatusCode}
		}
		c.Logger.Debug("DoRequest",
			slog.String("package", "internal"),
			slog.String("type", "Client"),
			slog.String("endpoint", endpoint),
			slog.String("error", err.Error()))
		return nil, err
	}

	c.Logger.Info("DoRequest",
		slog.String("package", "internal"),
		slog.String("type", "Client"),
		slog.String("endpoint", endpoint),
		slog.Int64("contentLength", response.ContentLength))

	return response, nil
}

func (c *Client) Get(endpoint string, pagination *api.PaginationQueryItems) (*http.Response, error) {
	return c.DoRequest("GET", endpoint, pagination)
}

func (c *Client) GetInto(endpoint string, target interface{}, paginationSetters ...api.PaginationOption) error {
	pagination := &api.PaginationQueryItems{
		Before: "",
		After:  "",
		Limit:  0,
	}

	for _, setter := range paginationSetters {
		setter(pagination)
	}

	c.Logger.Info("GetInto",
		slog.String("package", "internal"),
		slog.String("type", "Client"),
		slog.String("endpoint", endpoint),
		slog.Group("queryItems",
			slog.String("before", pagination.Before),
			slog.String("after", pagination.After),
			slog.Int("limit", pagination.Limit)))

	response, err := c.Get(endpoint, pagination)
	if err != nil {
		c.Logger.Debug("GetInto",
			slog.String("package", "internal"),
			slog.String("type", "Client"),
			slog.String("endpoint", endpoint),
			slog.String("error", err.Error()))
		return err
	}
	if err := json.NewDecoder(response.Body).Decode(target); err != nil {
		c.Logger.Debug("GetInto",
			slog.String("package", "internal"),
			slog.String("type", "Client"),
			slog.String("endpoint", endpoint),
			slog.String("error", err.Error()))
		return err
	}
	return nil
}
