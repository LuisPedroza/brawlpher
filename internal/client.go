package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/LuisPedroza/brawlpher/api"
)

const (
	apiURLFormat              = "%s://%s%s"
	baseURL                   = "bsproxy.royaleapi.dev"
	scheme                    = "https"
	apiTokenHeaderKey         = "Authorization"
	apiTokenHeaderValueFormat = "Bearer %s"
)

type Client struct {
	APIKey string
	Client Requester
}

func NewClient(apiKey string, client Requester) *Client {
	return &Client{
		APIKey: apiKey,
		Client: client,
	}
}

func (c *Client) NewRequest(method, endpoint string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, fmt.Sprintf(apiURLFormat, scheme, baseURL, endpoint), body)
	if err != nil {
		return nil, err
	}

	request.Header.Add(apiTokenHeaderKey, fmt.Sprintf(apiTokenHeaderValueFormat, c.APIKey))
	request.Header.Add("Accept", "application/json")

	return request, nil

}

func (c *Client) DoRequest(method, endpoint string, body io.Reader) (*http.Response, error) {
	request, err := c.NewRequest(method, endpoint, body)
	if err != nil {
		return nil, err
	}

	response, err := c.Client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == http.StatusServiceUnavailable {
		time.Sleep(time.Second)
		response, err = c.Client.Do(request)
		if err != nil {
			return nil, err
		}
	}
	if response.StatusCode < 200 || response.StatusCode > 299 {
		err, ok := api.StatusToError[response.StatusCode]
		if !ok {
			err = api.Error{Message: "unknown error", StatusCode: response.StatusCode}
		}
		return nil, err
	}

	return response, nil
}

func (c *Client) Get(endpoint string) (*http.Response, error) {
	return c.DoRequest("GET", endpoint, nil)
}

func (c *Client) GetInto(endpoint string, target interface{}) error {
	response, err := c.Get(endpoint)
	if err != nil {
		return err
	}
	if err := json.NewDecoder(response.Body).Decode(target); err != nil {
		return err
	}
	return nil
}
