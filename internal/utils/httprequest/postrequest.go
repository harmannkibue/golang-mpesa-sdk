package httprequest

import (
	"bytes"
	"log"
	"net/http"
	"time"
)

var BackOffStrategy = []time.Duration{
	1 * time.Second,
	3 * time.Second,
	4 * time.Second,
}

// PerformPost used to perform raw post request -.
func (h HttpRequest) PerformPost(args RequestDataParams) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, args.Endpoint, bytes.NewBuffer(args.Data))

	if err != nil {
		return nil, err
	}

	// Setting the necessary headers even for authentication token to formance -.
	req.Header.Set("Content-Type", args.ContentType)

	data, err := h.RetryDo(req, 3, time.Second*10, BackOffStrategy)

	if err != nil {
		log.Println("ERROR EXECUTING POST REQUEST CLIENT ", err.Error())
		return nil, err
	}

	return data, nil
}
