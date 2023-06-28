package main

import (
	"fmt"
	"github.com/harmannkibue/golang-mpesa-sdk/pkg/daraja"
	"log"
	"os"
)

var (
	// Set environment variables for daraja before testing -.
	mpesaApiKey         = os.Getenv("MPESA_KEY")
	mpesaConsumerSecret = os.Getenv("MPESA_SECRET")
	mpesaPassKey        = "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919"
)

func main() {
	// Initialize daraja service. PassKey is provided on your portal for sandbox,
	// while it is shared on your go live email from apisupport@safaricom.co.ke
	darajaService, err := daraja.New(mpesaApiKey, mpesaConsumerSecret, mpesaPassKey, daraja.SANDBOX)

	if err != nil {
		log.Println("failed initializing safaricom daraja client ", err)
	}

	// Implements getting token from daraja, if not available in the memory cache -.
	token, err := darajaToken(darajaService)

	if err != nil {
		log.Println("Error fetching token ", err.Error())
	}

	log.Println("Daraja Token ", token)

	// Implements stk push service -.
	stkRes, err := initiateStkPush(darajaService)

	if err != nil {
		log.Println("Error in stk push initiation ", err.Error())
	}
	log.Printf("STKPUSH response is %+v \n", stkRes)

	// Implements registering a confirmation and validation url.If response code is zero then it passed -.
	confirmationResponseCode, err := registerConfirmationUrl(darajaService)

	if err != nil {
		log.Println("Error registering a URL ", err.Error())
	}

	log.Println("Register URL response code ", confirmationResponseCode)

	// Simulate C2B transaction -.
	simulateResponse, err := simulateC2BPayment(darajaService)

	if err != nil {
		log.Println("Error simulating C2B request: ", err.Error())
	}

	fmt.Printf("C2B Response: %+v \n ", simulateResponse)

}

// Simulate a c2b payment which normally occur when you do a paybill payment -.
func simulateC2BPayment(darajaService *daraja.DarajaService) (*daraja.C2BSimulateResponse, error) {
	simulateResponse, err := darajaService.C2BSimulate(daraja.C2BSimulateRequestBody{
		ShortCode:     600982,
		CommandID:     "CustomerPayBillOnline",
		Amount:        1,
		Msisdn:        254708374149,
		BillRefNumber: "Unique reference",
	})

	if err != nil {
		return nil, err
	}

	return simulateResponse, nil
}

// Calls the service to initiate stk push -.
func initiateStkPush(darajaService *daraja.DarajaService) (*daraja.StkPushResponse, error) {
	// "CustomerPayBillOnline" for PayBill Numbers and "CustomerBuyGoodsOnline" for Till Numbers.
	stkRes, err := darajaService.InitiateStkPush(daraja.STKPushBody{
		BusinessShortCode: "174379",
		TransactionType:   "CustomerPayBillOnline",
		Amount:            "1",
		PartyA:            "254728922269",
		PartyB:            "174379",
		PhoneNumber:       "254728922269",
		CallBackURL:       "https://webhook.site/c882c5f6-4209-4f12-911b-85f13a69eb65",
		AccountReference:  "Custom unique identifier",
		TransactionDesc:   "Daraja sdk testing STK push",
	})

	if err != nil {
		return nil, err
	}

	return stkRes, nil
}

// Calls service to register urls -.
func registerConfirmationUrl(darajaService *daraja.DarajaService) (string, error) {
	regUrl, err := darajaService.C2BRegisterURL(daraja.RegisterC2BURLBody{
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

// Calls reusable service to request token from daraja -.
func darajaToken(darajaService *daraja.DarajaService) (string, error) {
	token, err := darajaService.GetToken()

	if err != nil {
		return "", err
	}

	return token, nil
}
