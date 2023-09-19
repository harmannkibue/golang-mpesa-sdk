package daraja

import (
	"encoding/json"
	"github.com/harmannkibue/golang-mpesa-sdk/internal/utils/httprequest"
)

// QueryAccountBalance checks account balance for both the B2C and C2B short codes -.
func (s DarajaService) QueryAccountBalance(accountBalance AccountBalanceRequestBody) (*AccountBalanceResponseBody, error) {
	body, err := json.Marshal(accountBalance)

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

	url := s.baseURL() + "mpesa/accountbalance/v1/query"

	response, err := s.HttpRequest.PerformPost(httprequest.RequestDataParams{
		Endpoint: url,
		Data:     body,
		Params:   make(map[string]string),
	}, BackOffStrategy,
		headers)

	if err != nil {
		return nil, err
	}

	// Unmarshal the response body into the RegisterC2BURLResponse struct -.
	var balanceResponse AccountBalanceResponseBody
	err = json.NewDecoder(response.Body).Decode(&balanceResponse)

	if err != nil {
		return nil, err
	}

	return &balanceResponse, nil
}
