package wallet

import (
	"net/http"
)

type RequestForDepositHistories struct {
}

type ResponseForDepositHistories []History

type History struct {
	Coin   string `json:"coin"`
	Status string `json:"status"`
	Txid   string `json:"txid"`

	Size float64 `json:"size"`
	Fee  float64 `json:"fee"`

	Confirmations int `json:"confirmations"`
	ID            int `json:"id"`

	ConfirmedTime string `json:"confirmedTime"`
	SentTime      string `json:"sentTime"`
	Time          string `json:"time"`
}

func (req *RequestForDepositHistories) Path() string {
	return "/wallet/deposits"
}

func (req *RequestForDepositHistories) Method() string {
	return http.MethodGet
}

func (req *RequestForDepositHistories) Query() string {
	return ""
}

func (req *RequestForDepositHistories) Payload() []byte {
	return nil
}
