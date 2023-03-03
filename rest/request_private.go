package rest

import (
	"errors"
	"fmt"

	"github.com/Roukii/kraken-go/openapi"
)

func (c *Client) GetAccountBalance() (*openapi.Balance, error) {
	request, err := openapi.NewGetAccountBalanceRequestWithFormdataBody(ENDPOINT+"/"+VERSION+"/", openapi.GetAccountBalanceFormdataRequestBody{})
	response, err := c.queryPrivate(*request)
	if err != nil {
		return nil, err
	}
	assets, err := openapi.ParseGetAccountBalanceResponse(response)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(assets.Body))

	return assets.JSON200.Result, nil
}

func (c *Client) CreateOrder(params openapi.Add) (string, error) {
	request, err := openapi.NewAddOrderRequestWithFormdataBody(ENDPOINT+"/"+VERSION+"/", params)
	if err != nil {
		return "", err
	}
	response, err := c.queryPrivate(*request)
	if err != nil {
		return "", err
	}
	assets, err := openapi.ParseAddOrderResponse(response)
	if err != nil {
		return "", err
	}
	fmt.Println(string(assets.Body))
	if (*assets.JSON200.Error) != nil && len(*assets.JSON200.Error) > 0 {
		return "", errors.New((*assets.JSON200.Error)[0])
	}
	return (*assets.JSON200.Result.Txid)[0], nil
}

func (c *Client) GetOrderStatus(orderId string) (*openapi.Closed, error) {
	request, err := openapi.NewGetOrdersInfoRequestWithFormdataBody(ENDPOINT+"/"+VERSION+"/", openapi.Query{
		Txid: orderId,
	})
	if err != nil {
		return nil, err
	}
	response, err := c.queryPrivate(*request)
	if err != nil {
		return nil, err
	}
	assets, err := openapi.ParseGetOrdersInfoResponse(response)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(assets.Body))
	closed, err := (*assets.JSON200.Result)[orderId].AsClosed()
	if err != nil {
		return nil, err
	}
	return &closed, nil
}

func (c *Client) CancelOrder(orderId string) (bool, error) {
	var txid openapi.Cancel_Txid
	err := txid.UnmarshalJSON([]byte(orderId))
	if err != nil {
		return false, err
	}
	request, err := openapi.NewCancelOrderRequestWithFormdataBody(ENDPOINT+"/"+VERSION+"/", openapi.Cancel{
		Nonce: 0,
		Txid:  txid,
	})
	response, err := c.queryPrivate(*request)
	if err != nil {
		return false, err
	}
	assets, err := openapi.ParseCancelOrderResponse(response)
	if err != nil {
		return false, err
	}
	if assets.JSON200.Error != nil && len(*assets.JSON200.Error) != 0 {
		return false, errors.New((*assets.JSON200.Error)[0][0])
	}
	fmt.Println(string(assets.Body))
	return (*assets.JSON200.Result.Count) > 0, nil
}

func (c *Client) GetWithdrawHistories() (*openapi.Balance, error) {
	request, err := openapi.NewGetStatusRecentWithdrawalsRequestWithFormdataBody(ENDPOINT+"/"+VERSION+"/", openapi.GetStatusRecentWithdrawalsFormdataRequestBody{
		Asset: "ZEUR",
		Nonce: 0,
	})
	response, err := c.queryPrivate(*request)
	if err != nil {
		return nil, err
	}
	assets, err := openapi.ParseGetStatusRecentWithdrawalsResponse(response)
	if err != nil {
		return nil, err
	}
	if assets.JSON200.Error != nil && len(*assets.JSON200.Error) != 0 {
		return nil, errors.New((*assets.JSON200.Error)[0])
	}
	fmt.Println(string(assets.Body))

	return nil, nil
}
