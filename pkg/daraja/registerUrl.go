package daraja

import (
	"encoding/json"
	"github.com/harmannkibue/golang-mpesa-sdk/internal/utils/httprequest"
	"log"
)

func (s DarajaService) C2BRegisterURL(c2bRegisterURL RegisterC2BURL) (*RegisterC2BURLResponse, error) {
	body, err := json.Marshal(c2bRegisterURL)

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

	url := s.baseURL() + "mpesa/c2b/v1/registerurl"

	response, err := s.HttpRequest.PerformPost(httprequest.RequestDataParams{
		Endpoint: url,
		Data:     body,
		Params:   make(map[string]string),
	}, BackOffStrategy,
		headers)

	if err != nil {
		return nil, err
	}

	// Unmarshal the response body into the RegisterC2BURLResponse struct
	var registerResponse RegisterC2BURLResponse
	err = json.NewDecoder(response.Body).Decode(&registerResponse)
	if err != nil {
		log.Fatal(err)
	}

	return &registerResponse, nil
}
