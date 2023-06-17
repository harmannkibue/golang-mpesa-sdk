package main

import (
	"fmt"
	"github.com/harmannkibue/golang-mpesa-sdk/pkg/daraja"
	"log"
	"os"
)

var (
	mpesaApiKey         = os.Getenv("MPESA_KEY")
	mpesaConsumerSecret = os.Getenv("MPESA_SECRET")
)

func main() {

	darajaService, err := daraja.New(os.Getenv("MPESA_KEY"), os.Getenv("MPESA_SECRET"), daraja.SANDBOX)

	if err != nil {
		log.Println("failed initializing safaricom daraja client ", err)
	}

	//TODO: Rearrange the examples correctly 2
	regUrl, err := darajaService.C2BRegisterURL(daraja.RegisterC2BURL{
		ShortCode:       "600989",
		ResponseType:    "Completed",
		ConfirmationURL: "https://webhook.site/c882c5f6-4209-4f12-911b-85f13a69eb65",
		ValidationURL:   "https://webhook.site/c882c5f6-4209-4f12-911b-85f13a69eb65",
	})

	if err != nil {
		fmt.Errorf("THE ERROR IN C2B ", err)
	}

	fmt.Printf("THE RES SS %+v ", regUrl)

	// TODO: Rearrange the examples correctly 1
	//token, err := darajaService.getToken()
	//
	//if err != nil {
	//	log.Println("The token not found")
	//}
	//
	//log.Println("THE TOKEN ISS ", token)
}
