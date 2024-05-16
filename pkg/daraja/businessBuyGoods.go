package daraja

import (
	"encoding/json"
	"github.com/harmannkibue/golang-mpesa-sdk/internal/utils/httprequest"
)

// BusinessToBusinessPayment sends money from your business account to till, merchant HO or merchant store number -.
func (s DarajaService) BusinessToBusinessPayment(b2cBody BusinessToBusinessRequestBody) (*BusinessToBusinessResponse, error) {
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

	url := s.baseURL() + "mpesa/b2b/v1/paymentrequest"

	response, err := s.HttpRequest.PerformPost(httprequest.RequestDataParams{
		Endpoint: url,
		Data:     body,
		Params:   make(map[string]string),
	}, BackOffStrategy,
		headers)

	if err != nil {
		return nil, err
	}

	// Unmarshal the response body into the Business buy goods struct -.
	var buyGoodsResponse BusinessToBusinessResponse
	err = json.NewDecoder(response.Body).Decode(&buyGoodsResponse)

	if err != nil {
		return nil, err
	}

	return &buyGoodsResponse, nil
}
