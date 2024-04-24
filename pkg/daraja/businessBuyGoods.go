package daraja

import (
	"encoding/json"
	"github.com/harmannkibue/golang-mpesa-sdk/internal/utils/httprequest"
)

// BusinessBuyGoods sends money from your business account to till, merchant HO or merchant store number -.
func (s DarajaService) BusinessBuyGoods(b2cBody BusinessBuyGoodsRequestBody) (*BusinessBuyGoodsResponse, error) {
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
	//url := "https://webhook.site/996ed649-244a-4285-85c4-ab0bac869920/" + "mpesa/b2b/v1/paymentrequest"

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
	var buyGoodsResponse BusinessBuyGoodsResponse
	err = json.NewDecoder(response.Body).Decode(&buyGoodsResponse)

	if err != nil {
		return nil, err
	}

	return &buyGoodsResponse, nil
}