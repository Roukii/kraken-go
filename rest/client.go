package rest

import (
	"net/http"
	"time"

	"github.com/Roukii/kraken-go/auth"
	"github.com/Roukii/kraken-go/openapi"
)

const ENDPOINT = "https://api.kraken.com"
const VERSION = "0"
const CONTENT_TYPE = "application/x-www-form-urlencoded; charset=utf-8"

type Client struct {
	Auth         *auth.Config
	KrakenClient *openapi.Client
	HTTPTimeout  time.Duration
}

func New(auth *auth.Config, httpClient *http.Client, timeout time.Duration) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	c, err := openapi.NewClient(ENDPOINT+"/"+VERSION, openapi.WithHTTPClient(httpClient))
	if err != nil {
		return nil
	}

	return &Client{
		Auth:        auth,
		KrakenClient: c,
		HTTPTimeout: timeout,
	}
}
