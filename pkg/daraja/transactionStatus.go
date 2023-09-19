package daraja

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/harmannkibue/golang-mpesa-sdk/internal/utils/httprequest"
)

// MpesaExpressTransactionStatus Transaction status query api -.
func (s DarajaService) MpesaExpressTransactionStatus(statusBody MpesaExpressTransactionStatusQueryBodyComplete) (*MpesaExpressTransactionStatusQueryResponse, error) {
	encodingPasswordData := fmt.Sprintf("%s%s%s", statusBody.BusinessShortCode, statusBody.Password, statusBody.Timestamp)

	password := base64.StdEncoding.EncodeToString([]byte(encodingPasswordData))

	fmt.Println("THE PASSWORD ENCODED ISS ", encodingPasswordData)

	data := MpesaExpressTransactionStatusQueryBodyComplete{BusinessShortCode: statusBody.BusinessShortCode, Password: password, Timestamp: statusBody.Timestamp, CheckoutRequestID: statusBody.CheckoutRequestID}

	body, err := json.Marshal(data)

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

	url := s.baseURL() + "mpesa/stkpushquery/v1/query"

	response, err := s.HttpRequest.PerformPost(httprequest.RequestDataParams{
		Endpoint: url,
		Data:     body,
		Params:   make(map[string]string),
	}, BackOffStrategy,
		headers)

	if err != nil {
		return nil, err
	}

	// Unmarshal the response body into the TransactionStatusResponseBody struct
	var statusResponse MpesaExpressTransactionStatusQueryResponse
	err = json.NewDecoder(response.Body).Decode(&statusResponse)

	if err != nil {
		return nil, err
	}

	return &statusResponse, nil
}

// TransactionStatus checks the account balance for a given short code -.
func (s DarajaService) TransactionStatus(statusBody TransactionStatusRequestBody) (*TransactionStatusResponseBody, error) {
	body, err := json.Marshal(statusBody)

	if err != nil {
		return nil, err
	}

	token, err := s.GetToken()
	fmt.Println("THE TOEKN ISS ", token)

	if err != nil {
		return nil, err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Bearer " + token
	headers["Cache-Control"] = "no-cache"

	url := s.baseURL() + "mpesa/transactionstatus/v1/query"

	response, err := s.HttpRequest.PerformPost(httprequest.RequestDataParams{
		Endpoint: url,
		Data:     body,
		Params:   make(map[string]string),
	}, BackOffStrategy,
		headers)

	if err != nil {
		return nil, err
	}

	// Unmarshal the response body into the TransactionStatusResponseBody struct
	var statusResponse TransactionStatusResponseBody
	err = json.NewDecoder(response.Body).Decode(&statusResponse)

	if err != nil {
		return nil, err
	}

	return &statusResponse, nil
}
