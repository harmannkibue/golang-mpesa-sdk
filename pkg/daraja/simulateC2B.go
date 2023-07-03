package daraja

import (
	"encoding/json"
	"github.com/harmannkibue/golang-mpesa-sdk/internal/utils/httprequest"
)

// C2BSimulate Simulate C2B Transaction, mainly for Sandbox -.
func (s DarajaService) C2BSimulate(c2bSimulateBody C2BSimulateRequestBody) (*C2BSimulateResponse, error) {
	body, err := json.Marshal(c2bSimulateBody)

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

	url := s.baseURL() + "mpesa/c2b/v1/simulate"

	response, err := s.HttpRequest.PerformPost(httprequest.RequestDataParams{
		Endpoint: url,
		Data:     body,
		Params:   make(map[string]string),
	}, BackOffStrategy,
		headers)

	if err != nil {
		return nil, err
	}

	// Unmarshal the response body into the C2BSimulateResponse struct
	var c2bSimulateResponse C2BSimulateResponse
	err = json.NewDecoder(response.Body).Decode(&c2bSimulateResponse)

	if err != nil {
		return nil, err
	}

	return &c2bSimulateResponse, nil
}
