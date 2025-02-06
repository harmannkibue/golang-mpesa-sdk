package daraja

import (
	"encoding/json"
	"github.com/harmannkibue/golang-mpesa-sdk/internal/utils/httprequest"
)

type ReceiverPartyIdentifierType string
type TransactionType string

const (
	ReceiverPartyMerchantTill      ReceiverPartyIdentifierType = "2"
	ReceiverPartyBusinessShortCode ReceiverPartyIdentifierType = "4"
)

const (
	TransactionTypePayBill     TransactionType = "Standing Order Customer Pay Bill"
	TransactionTypePayMerchant TransactionType = "Standing Order Customer Pay Merchant"
)

type MpesaRatibaRequestBody struct {
	StandingOrderName           string                      `json:"StandingOrderName"`
	ReceiverPartyIdentifierType ReceiverPartyIdentifierType `json:"ReceiverPartyIdentifierType"`
	TransactionType             TransactionType             `json:"TransactionType"`
	BusinessShortCode           string                      `json:"BusinessShortCode"`
	PartyA                      string                      `json:"PartyA"`
	Amount                      string                      `json:"Amount"`
	StartDate                   string                      `json:"StartDate"`
	EndDate                     string                      `json:"EndDate"`
	Frequency                   string                      `json:"Frequency"`
	AccountReference            string                      `json:"AccountReference"`
	TransactionDesc             string                      `json:"TransactionDesc"`
	CallBackURL                 string                      `json:"CallBackURL"`
}

type MpesaRatibaRequestResponseBody struct {
	ResponseHeader struct {
		ResponseRefID       string `json:"responseRefID"`
		ResponseCode        string `json:"responseCode"`
		ResponseDescription string `json:"responseDescription"`
	} `json:"ResponseHeader"`
	ResponseBody struct {
		ResponseDescription string `json:"responseDescription"`
		ResponseCode        string `json:"responseCode"`
	} `json:"ResponseBody"`
}

// InitiateMpesaRatibaRequest Initiate an Mpesa Ratiba standing order to customer -.
func (s DarajaService) InitiateMpesaRatibaRequest(stkRequest MpesaRatibaRequestBody) (*MpesaRatibaRequestResponseBody, error) {

	body, err := json.Marshal(MpesaRatibaRequestBody{
		StandingOrderName:           stkRequest.StandingOrderName,
		ReceiverPartyIdentifierType: stkRequest.ReceiverPartyIdentifierType,
		BusinessShortCode:           stkRequest.BusinessShortCode,
		TransactionType:             stkRequest.TransactionType,
		PartyA:                      stkRequest.PartyA,
		Amount:                      stkRequest.Amount,
		StartDate:                   stkRequest.StartDate,
		EndDate:                     stkRequest.EndDate,
		Frequency:                   stkRequest.Frequency,
		AccountReference:            stkRequest.AccountReference,
		TransactionDesc:             stkRequest.TransactionDesc,
		CallBackURL:                 stkRequest.CallBackURL,
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

	url := s.baseURL() + "standingorder/v1/createStandingOrderExternal"

	response, err := s.HttpRequest.PerformPost(httprequest.RequestDataParams{
		Endpoint: url,
		Data:     body,
		Params:   make(map[string]string),
	}, BackOffStrategy,
		headers)

	if err != nil {
		return nil, err
	}

	var mRatibaResponse MpesaRatibaRequestResponseBody
	err = json.NewDecoder(response.Body).Decode(&mRatibaResponse)

	if err != nil {
		return nil, err
	}

	return &mRatibaResponse, nil
}
