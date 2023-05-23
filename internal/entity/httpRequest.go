package entity

type RequestDataParams struct {
	Endpoint    string            `json:"endpoint"`
	ContentType string            `json:"content_type"`
	Data        []byte            `json:"data"`
	Params      map[string]string `json:"params"`
}
