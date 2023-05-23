package main

import (
	"fmt"
	"github.com/harmannkibue/golang-mpesa-sdk/internal/utils/httprequest"
)

func main() {
	fmt.Println("Heere in the main function!")

	httpReq := httprequest.HttpRequest{}

	resp, err := httpReq.PerformPost(httprequest.RequestDataParams{
		Endpoint:    "https://webhook.site/93e45dd0-d2ce-42e9-8f90-a55767a4e982",
		ContentType: "application/json",
		Data:        []byte(`{"key": "value"}`),
	})

	if err != nil {
		return
	}

	fmt.Println("THE RESPONSE ISS ", resp.StatusCode)
}
