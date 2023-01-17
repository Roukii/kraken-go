package rest

import (
	"fmt"

	"github.com/Roukii/kraken-go/openapi"
)

func (c *Client) GetTradableAssetPairs(params *openapi.GetTradableAssetPairsParams) (*map[string]openapi.Pairs, error) {
	request, err := openapi.NewGetTradableAssetPairsRequest(ENDPOINT+"/"+VERSION+"/", params)
	response, err := c.queryPublic(*request)
	if err != nil {
		fmt.Println("err : " + err.Error())
		return nil, err
	}
	fmt.Println("response " + response.Status)
	assets, err := openapi.ParseGetTradableAssetPairsResponse(response)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println(string(assets.Body))
	return assets.JSON200.Result, nil
}

func (c *Client) GetTickerInformation(params *openapi.GetTickerInformationParams) (*map[string]openapi.Ticker, error) {
	request, err := openapi.NewGetTickerInformationRequest(ENDPOINT+"/"+VERSION+"/", params)
	response, err := c.queryPublic(*request)
	if err != nil {
		fmt.Println("err : " + err.Error())
		return nil, err
	}
	fmt.Println("response " + response.Status)
	assets, err := openapi.ParseGetTickerInformationResponse(response)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return assets.JSON200.Result, nil
}