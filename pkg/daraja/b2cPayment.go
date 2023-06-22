package daraja

import (
	"encoding/json"
	"github.com/harmannkibue/golang-mpesa-sdk/internal/utils/httprequest"
)

// B2CPayment Supported CommandID SalaryPayment - This supports sending money to both registered and unregistered M-Pesa customers
// BusinessPayment - This is a normal business to customer payment,  supports only M-Pesa registered customers
// PromotionPayment - This is a promotional payment to customers. The M-Pesa notification message is a congratulatory message. Supports only M-Pesa registered customers.
func (s DarajaService) B2CPayment(b2cBody B2CRequestBody) (*B2CResponseBody, error) {
	body, err := json.Marshal(b2cBody)

	if err != nil {
		return nil, err
	}

	token, err := s.GetToken()

	if err != nil {
		return nil, err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Bearer " + token
	headers["Cache-Control"] = "no-cache"

	url := s.baseURL() + "mpesa/b2c/v1/paymentrequest"

	response, err := s.HttpRequest.PerformPost(httprequest.RequestDataParams{
		Endpoint: url,
		Data:     body,
		Params:   make(map[string]string),
	}, BackOffStrategy,
		headers)

	if err != nil {
		return nil, err
	}

	// Unmarshal the response body into the B2CResponseBody struct
	var b2cResponse B2CResponseBody
	err = json.NewDecoder(response.Body).Decode(&b2cResponse)

	if err != nil {
		return nil, err
	}

	return &b2cResponse, nil
}
