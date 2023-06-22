package daraja

import (
	"encoding/json"
	"github.com/harmannkibue/golang-mpesa-sdk/internal/utils/httprequest"
)

// C2BTransactionReversal reverses C2B payment -.
func (s DarajaService) C2BTransactionReversal(reversalBody TransactionReversalRequestBody) (*TransactionReversalResponseBody, error) {
	body, err := json.Marshal(reversalBody)

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

	url := s.baseURL() + "mpesa/reversal/v1/request"

	response, err := s.HttpRequest.PerformPost(httprequest.RequestDataParams{
		Endpoint: url,
		Data:     body,
		Params:   make(map[string]string),
	}, BackOffStrategy,
		headers)

	if err != nil {
		return nil, err
	}

	// Unmarshal the response body into the TransactionReversalResponseBody struct
	var reversalResponse TransactionReversalResponseBody
	err = json.NewDecoder(response.Body).Decode(&reversalResponse)

	if err != nil {
		return nil, err
	}

	return &reversalResponse, nil
}
