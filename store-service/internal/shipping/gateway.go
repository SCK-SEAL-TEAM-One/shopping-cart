package shipping

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"store-service/internal/order"
)

type Shipping interface {
	ShipByKerry(shippingInfo order.ShippingInfo) (string, error)
}

type ShippingGateway struct {
	KerryEndpoint string
}

type ShippingGatewayResponse struct {
	TrackingID string `json:"tracking_id"`
}

func (gateway ShippingGateway) ShipByKerry(shippingInfo order.ShippingInfo) (string, error) {
	data, _ := json.Marshal(shippingInfo)
	endPoint := gateway.KerryEndpoint + "/shipping/kerry"
	response, err := http.Post(endPoint, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	if response.StatusCode != 200 {
		return "", fmt.Errorf("response is not ok but it's %d", response.StatusCode)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var shippingGatewayResponse ShippingGatewayResponse
	err = json.Unmarshal(responseData, &shippingGatewayResponse)
	if err != nil {
		return "0", err
	}

	return shippingGatewayResponse.TrackingID, nil
}
