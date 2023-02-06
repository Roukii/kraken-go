package rest

import (
	"errors"
	"fmt"

	"github.com/Roukii/kraken-go/openapi"
)

func (c *Client) GetTradableAssetPairs(params *openapi.GetTradableAssetPairsParams) (*map[string]openapi.Pairs, error) {
	request, err := openapi.NewGetTradableAssetPairsRequest(ENDPOINT+"/"+VERSION+"/", params)
	response, err := c.queryPublic(*request)
	if err != nil {
		return nil, err
	}
	assets, err := openapi.ParseGetTradableAssetPairsResponse(response)
	if err != nil {
		return nil, err
	}
	if assets.JSON200.Error != nil && len(*assets.JSON200.Error) != 0 {
		return nil, errors.New((*assets.JSON200.Error)[0])
	}
	fmt.Println(string(assets.Body))
	return assets.JSON200.Result, nil
}

func (c *Client) GetTickerInformation(params *openapi.GetTickerInformationParams) (*map[string]openapi.Ticker, error) {
	request, err := openapi.NewGetTickerInformationRequest(ENDPOINT+"/"+VERSION+"/", params)
	response, err := c.queryPublic(*request)
	if err != nil {
		return nil, err
	}
	assets, err := openapi.ParseGetTickerInformationResponse(response)
	if err != nil {
		return nil, err
	}
	if assets.JSON200.Error != nil && len(*assets.JSON200.Error) != 0 {
		return nil, errors.New((*assets.JSON200.Error)[0])
	}
	fmt.Println(string(assets.Body))

	return assets.JSON200.Result, nil
}

func (c *Client) GetAssets() (*map[string]openapi.Info, error) {
	request, err := openapi.NewGetAssetInfoRequest(ENDPOINT+"/"+VERSION+"/", &openapi.GetAssetInfoParams{})
	response, err := c.queryPublic(*request)
	if err != nil {
		return nil, err
	}
	assets, err := openapi.ParseGetAssetInfoResponse(response)
	if err != nil {
		return nil, err
	}
	if assets.JSON200.Error != nil && len(*assets.JSON200.Error) != 0 {
		return nil, errors.New((*assets.JSON200.Error)[0])
	}
	fmt.Println(string(assets.Body))

	return assets.JSON200.Result, nil
}
