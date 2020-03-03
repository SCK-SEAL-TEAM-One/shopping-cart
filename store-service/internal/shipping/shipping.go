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
	resp, err := http.Post(gateway.KerryEndpoint, "application/json", bytes.NewBuffer(data))
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

	var shippingGatewayResponse ShippingGatewayResponse
	err = json.Unmarshal(responseData, &shippingGatewayResponse)
	if err != nil {
		return "0", err
	}

	return shippingGatewayResponse.TrackingID, nil
}
