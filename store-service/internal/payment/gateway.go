package payment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type BankGateway struct {
	BankEndpoint string
}

type BankGatewayResponse struct {
	Status          string  `json:"status"`
	Fee             float64 `json:"fee"`
	AvailableBlance float64 `json:"available_balance"`
	Authorized      string  `json:"authorized"`
	TransactionID   string  `json:"transaction_id"`
}

func (gateway BankGateway) Payment(paymentDetail PaymentDetail) (string, error) {
	data, _ := json.Marshal(paymentDetail)
	resp, err := http.Post(gateway.BankEndpoint, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("response is not ok but it's %d", resp.StatusCode)
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var BankGatewayResponse BankGatewayResponse
	err = json.Unmarshal(responseData, &BankGatewayResponse)
	if err != nil {
		return "", err
	}

	return BankGatewayResponse.TransactionID, nil
}
