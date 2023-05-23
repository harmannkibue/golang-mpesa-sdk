package httprequest

// IhttpRequest Interface implementation for custom http request.This helps in mock testing.
type IhttpRequest interface {
}

type HttpRequest struct {
}

func NewHttpRequest() IhttpRequest {
	return HttpRequest{}
}

type RequestDataParams struct {
	Endpoint    string            `json:"endpoint"`
	ContentType string            `json:"content_type"`
	Data        []byte            `json:"data"`
	Params      map[string]string `json:"params"`
}
