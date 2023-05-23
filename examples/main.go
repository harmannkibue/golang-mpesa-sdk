package main

import (
	"fmt"
	"os"
)

var (
	mpesaApiKey         = os.Getenv("MPESA_KEY")
	mpesaConsumerSecret = os.Getenv("MPESA_SECRET")
)

func main() {
	fmt.Printf("Safaricom: %s AND %s ", os.Getenv("MPESA_KEY"), os.Getenv("MPESA_SECRET"))

	//httpReq := httprequest.HttpRequest{}
	//
	//resp, err := httpReq.PerformPost(httprequest.RequestDataParams{
	//	Endpoint:    "https://webhook.site/93e45dd0-d2ce-42e9-8f90-a55767a4e982",
	//	ContentType: "application/json",
	//	Data:        []byte(`{"key": "value"}`),
	//})
	//
	//if err != nil {
	//	return
	//}

	//fmt.Println("THE RESPONSE ISS ", resp.StatusCode)
}
