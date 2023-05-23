package httprequest

import (
	"bytes"
	"github.com/harmannkibue/golang-mpesa-sdk/internal/entity"
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
func PerformPost(args entity.RequestDataParams) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, args.Endpoint, bytes.NewBuffer(args.Data))

	if err != nil {
		return nil, err
	}

	// Setting the necessary headers even for authentication token to formance -.
	req.Header.Set("Content-Type", args.ContentType)

	data, err := RetryDo(req, 3, time.Second*10, BackOffStrategy)

	if err != nil {
		log.Println("ERROR EXECUTING POST REQUEST CLIENT ", err.Error())
		return nil, err
	}

	return data, nil
}
