package main

import (
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

	token, err := darajaService.getToken()

	if err != nil {
		log.Println("The token not found")
	}

	log.Println("THE TOKEN ISS ", token)
}
