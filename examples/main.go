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

	// Implements getting token from daraja, if not available in the memory cache -.
	token, err := darajaToken(darajaService)

	if err != nil {
		log.Println("Error fetching token ", err.Error())
	}

	log.Println("Daraja Token ", token)

	// Implements registering a confirmation and validation url.If response code is zero then it passed -.
	confirmationResponseCode, err := registerConfirmationUrl(darajaService)

	if err != nil {
		log.Println("Error registering a URL ", err.Error())
	}

	log.Println("Register URL response code ", confirmationResponseCode)

}

func registerConfirmationUrl(darajaService *daraja.DarajaService) (string, error) {
	regUrl, err := darajaService.C2BRegisterURL(daraja.RegisterC2BURL{
		ShortCode:       "600989",
		ResponseType:    "Completed",
		ConfirmationURL: "https://webhook.site/c882c5f6-4209-4f12-911b-85f13a69eb65",
		ValidationURL:   "https://webhook.site/c882c5f6-4209-4f12-911b-85f13a69eb65",
	})

	if err != nil {
		return "", err
	}
	return regUrl.ResponseCode, nil
}

func darajaToken(darajaService *daraja.DarajaService) (string, error) {
	token, err := darajaService.GetToken()

	if err != nil {
		return "", err
	}

	return token, nil
}
