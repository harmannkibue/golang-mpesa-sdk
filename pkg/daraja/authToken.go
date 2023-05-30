package daraja

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type tokenResponse struct {
	AccessToken string `json:"access_token"`
}

// Fetching the token from safaricom daraja API -.
func (s DarajaService) getToken() (string, error) {
	darajaToken, err := s.Cache.GetCacheValue()

	fmt.Println("CACHE GOTTEN FROM STORE ", darajaToken)

	if err != nil {
		//The token is not found.
		tokenValue, err := s.fetchTokenFromMpesa()

		if err != nil {
			return "", err
		}

		// Fetch it from mpesa and set it in the cache

		err = s.Cache.SetCacheValue(tokenValue)

		if err != nil {
			return "", err
		}

		return tokenValue, nil

	}

	return darajaToken, nil
}

func (s DarajaService) fetchTokenFromMpesa() (string, error) {
	b := []byte(s.ApiKey + ":" + s.ApiSecret)
	encoded := base64.StdEncoding.EncodeToString(b)

	url := s.baseURL() + "oauth/v1/generate?grant_type=client_credentials"
	req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(encoded))
	if err != nil {
		return "", err
	}
	req.Header.Add("authorization", "Basic "+encoded)
	req.Header.Add("cache-control", "no-cache")

	response, err := s.HttpRequest.RetryDo(req, 3, time.Second*10, BackOffStrategy)

	if err != nil {
		return "", fmt.Errorf("failed to fetch token %v ", err)
	}

	var tokenResp tokenResponse
	err = json.NewDecoder(response.Body).Decode(&tokenResp)
	if err != nil {
		return "", fmt.Errorf("could not decode auth response: %v", err)
	}

	accessToken := tokenResp.AccessToken

	return accessToken, nil
}
