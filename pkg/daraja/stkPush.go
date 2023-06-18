package daraja

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/harmannkibue/golang-mpesa-sdk/internal/utils/httprequest"
	"time"
)

type completeStkRequestBody struct {
	BusinessShortCode string `json:"BusinessShortCode"`
	Password          string `json:"Password"`
	Timestamp         string `json:"Timestamp"`
	TransactionType   string `json:"TransactionType"`
	Amount            string `json:"Amount"`
	PartyA            string `json:"PartyA"`
	PartyB            string `json:"PartyB"`
	PhoneNumber       string `json:"PhoneNumber"`
	CallBackURL       string `json:"CallBackURL"`
	AccountReference  string `json:"AccountReference"`
	TransactionDesc   string `json:"TransactionDesc"`
}

func (s DarajaService) InitiateStkPush(stkRequest STKPushBody) (*StkPushResponse, error) {
	timestamp := time.Now().Format("20060102150405")

	passwordMessage := stkRequest.BusinessShortCode + s.ApiSecret + timestamp
	password := base64.StdEncoding.EncodeToString([]byte(passwordMessage))

	body, err := json.Marshal(completeStkRequestBody{
		BusinessShortCode: stkRequest.BusinessShortCode,
		Password:          password,
		Timestamp:         timestamp,
		TransactionType:   stkRequest.TransactionType,
		Amount:            stkRequest.Amount,
		PartyA:            stkRequest.PartyA,
		PartyB:            stkRequest.PartyB,
		PhoneNumber:       stkRequest.PhoneNumber,
		CallBackURL:       stkRequest.CallBackURL,
		AccountReference:  stkRequest.AccountReference,
		TransactionDesc:   stkRequest.TransactionDesc,
	})

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

	url := s.baseURL() + "mpesa/stkpush/v1/processrequest"
	//url := "https://webhook.site/c882c5f6-4209-4f12-911b-85f13a69eb65"

	response, err := s.HttpRequest.PerformPost(httprequest.RequestDataParams{
		Endpoint: url,
		Data:     body,
		Params:   make(map[string]string),
	}, BackOffStrategy,
		headers)

	if err != nil {
		return nil, err
	}

	fmt.Printf("THE STK RESPONSE %+v", response)

	var stkResponse StkPushResponse
	err = json.NewDecoder(response.Body).Decode(&stkResponse)

	if err != nil {
		return nil, err
	}

	return &stkResponse, nil
}
