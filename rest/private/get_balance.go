package private

import (
	"net/http"
)

type RequestForGetAccountBalance struct {
	Uuid string
}

type ResponseForGetAccountBalance struct {
	// Account Account `json:"account"`
}

func (req *RequestForGetAccountBalance) Path() string {
	return "/api/v3/brokerage/accounts/" + req.Uuid
}

func (req *RequestForGetAccountBalance) Method() string {
	return http.MethodGet
}

func (req *RequestForGetAccountBalance) Query() string {
	return ""
}

func (req *RequestForGetAccountBalance) Payload() []byte {
	return nil
}