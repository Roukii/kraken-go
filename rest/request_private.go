package rest

import (
	"fmt"

	"github.com/Roukii/kraken-go/openapi"
)

func (c *Client) GetAccountBalance() (*openapi.Balance, error) {
	request, err := openapi.NewGetAccountBalanceRequestWithFormdataBody(ENDPOINT+"/"+VERSION+"/", openapi.GetAccountBalanceFormdataRequestBody{})
	response, err := c.queryPrivate(*request)
	if err != nil {
		fmt.Println("err : " + err.Error())
		return nil, err
	}
	fmt.Println("response " + response.Status)
	fmt.Println(response.ContentLength)
	assets, err := openapi.ParseGetAccountBalanceResponse(response)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println(string(assets.Body))

	return assets.JSON200.Result, nil
}


func (c *Client) CreateOrder(params openapi.Add) (*openapi.Balance, error) {
	request, err := openapi.NewAddOrderRequestWithFormdataBody(ENDPOINT+"/"+VERSION+"/", params)
	response, err := c.queryPrivate(*request)
	if err != nil {
		fmt.Println("err : " + err.Error())
		return nil, err
	}
	fmt.Println("response " + response.Status)
	fmt.Println(response.ContentLength)
	assets, err := openapi.ParseAddOrderResponse(response)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println(string(assets.Body))

	return nil, nil
}

func (c *Client) GetWithdrawHistories() (*openapi.Balance, error) {
	request, err := openapi.NewGetStatusRecentWithdrawalsRequestWithFormdataBody(ENDPOINT+"/"+VERSION+"/", openapi.GetStatusRecentWithdrawalsFormdataRequestBody{
		Asset: "ZEUR",
		Nonce: 0,
	})
	response, err := c.queryPrivate(*request)
	if err != nil {
		fmt.Println("err : " + err.Error())
		return nil, err
	}
	fmt.Println("response " + response.Status)
	fmt.Println(response.ContentLength)
	assets, err := openapi.ParseGetStatusRecentWithdrawalsResponse(response)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println(string(assets.Body))

	return nil, nil
}
