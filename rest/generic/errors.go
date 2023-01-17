package generic

import (
	"fmt"
)

// APIError return the api error
type APIError struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	ErrorId string         `json:"id"`
	Code    int            `json:"code"`
	Type    string         `json:"error"`
	Details map[string]any `json:"details"`
}

// Error return the error message
func (e APIError) Error() string {
	return fmt.Sprintf("APIError: status=%d, message=%s errorId=%s", e.Status, e.Message, e.ErrorId)
}