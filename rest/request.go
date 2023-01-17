package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/Roukii/kraken-go/rest/generic"
)

// queryPrivate executes a private method query
func (c *Client) queryPrivate(request http.Request) (*http.Response, error) {
	if c.Auth == nil {
		return c.KrakenClient.Client.Do(&request)
	}
	urlPath := request.URL.Path
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}
	values, err := url.ParseQuery(string(body))
	values.Set("nonce", fmt.Sprintf("%d", time.Now().UnixNano()))
	// Create signature
	signature, values := c.Auth.SetAuthValueAndCreateSignature(urlPath, values)
	new_body := values.Encode()
	request.Body = ioutil.NopCloser(strings.NewReader(new_body))
	request.ContentLength = int64(len(new_body))
	// Add Key and signature to request headers
	request.Header.Add("API-Key", c.Auth.Key)
	request.Header.Add("API-Sign", signature)
	fmt.Println("do request")
	return c.KrakenClient.Client.Do(&request)
}

// queryPublic executes a public method query
func (c *Client) queryPublic(request http.Request) (*http.Response, error) {
	return c.KrakenClient.Client.Do(&request)
}

func (c *Client) do(request http.Request, reqURL string, values url.Values, headers map[string]string) (*http.Response, error) {
	tmp := strings.NewReader(values.Encode())
	// Create request
	req, err := http.NewRequest("POST", reqURL, tmp)
	if err != nil {
		return nil, fmt.Errorf("Could not execute request! #1 (%s)", err.Error())
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	// Execute request
	res, err := c.KrakenClient.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		var r interface{}
		var data []byte
		res.Body.Read(data)
		if err := json.Unmarshal(data, &r); err != nil {
			return nil, &generic.APIError{
				Status:  res.StatusCode,
				Message: fmt.Sprintf("%+v\n", string(data)),
			}
		}
		if resp, ok := r.(generic.APIError); ok {
			return nil, resp
		}
		return nil, &generic.APIError{
			Status:  res.StatusCode,
			Message: fmt.Sprintf("%+v\n", string(data)),
		}

	}
	return res, nil
}
