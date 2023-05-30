package daraja

import "encoding/json"

func (s DarajaService) C2BRegisterURL(c2bRegisterURL RegisterC2BURL) (string, error) {
	body, err := json.Marshal(c2bRegisterURL)
	if err != nil {
		return "", err
	}

	token, err := s.getToken()

	if err != nil {
		return "", nil
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Bearer " + token
	headers["Cache-Control"] = "no-cache"

	params :=

	url := s.baseURL() + "mpesa/c2b/v1/registerurl"
	return s.newReq(url, body, headers)
}
