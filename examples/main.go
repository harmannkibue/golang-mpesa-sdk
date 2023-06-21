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

	balance, err := accountBalance(darajaService)

	if err != nil {
		fmt.Println("THE BALANCE ERROR ", err.Error())
	}

	fmt.Printf("THE BALANCE RESPONSE %+v \n", balance)
}

// checking account balance for both B2C and C2B short codes -.
func accountBalance(darajaService *daraja.DarajaService) (*daraja.AccountBalanceResponseBody, error) {
	balance, err := darajaService.QueryAccountBalance(daraja.AccountBalanceRequestBody{
		Initiator:          "testapi",
		SecurityCredential: "RZiAQmdpGfWEw2HggGwFCOE/1mr2aEfIUVradvgi3vSWzJDrchOXXUQmgKf8Z21NR38AN4n523sXsl3Trcvl2JQx8uylLxTKE+lG5RgfCauzG3qIBkAJCuklo1hJLj5h1CNht8U9Ac011DaN3lpzy39CTjO/Y/BEdSisX72cogeb+GehAZa0Q0opj8XpfAy2nPQ5fosF79P6KDm/pDjyLFamzZIoUGSIBUaEVwCdWm/AkdaXooYV3LWk4pwYkqtDWeFELFW29a2FYRBT6BU2JZttJO8XFmZI/x+U/ZQuQyC5CltA9CdgjXTHrK9LWZ7PoKwCESWbytYH2eddQd3VsA==",
		CommandID:          "AccountBalance",
		PartyA:             600991,
		IdentifierType:     2,
		Remarks:            "Churpy Balance",
		QueueTimeOutURL:    "https://webhook.site/7da5ccfd-3a90-4038-b822-273887b3de7f",
		ResultURL:          "https://webhook.site/7da5ccfd-3a90-4038-b822-273887b3de7f",
	})

	if err != nil {
		return nil, err
	}

	return balance, nil
}

// Implements business to customer payment -.
func b2cPayment(darajaService *daraja.DarajaService) (*daraja.B2CResponseBody, error) {
	b2cPayment, err := darajaService.B2CPayment(daraja.B2CRequestBody{
		InitiatorName:      "testapi",
		SecurityCredential: "UKCrm4IVKWEoW640M3pUHS4hZ2ynDpz+LT6c+acBK28TOMULxVhMP0YM2FNCh2QXx+m6HR8iLNsR0bfbIB1kpvNhciKUrn7Glp4f7UNPF8mHXgNsa/09+i7X8+JUy7tQLEOoPE/xCWBOh2ofBq8N+lX77RUAxDp9HC8Nj6nN6kH07Ygmz7NnRd/dlayqcFKV4UNP/nQAV8lum2HSh9xRBnlexcziYipt/d293qrSSvXtAfz+lmgzzbzwML02zlCQxXS2YQjTluQWzRgxkl+9aCCs51a5BWppTE6iYd8qcMlX/+hMZvl2D9LjQKwisSKJsWP2MtxFxG86DRpwI41I4A==",
		CommandID:          "SalaryPayment",
		Amount:             1,
		PartyA:             600998,
		PartyB:             254728922269,
		Remarks:            "Payment from VA",
		QueueTimeOutURL:    "https://webhook.site/7da5ccfd-3a90-4038-b822-273887b3de7f",
		ResultURL:          "https://webhook.site/7da5ccfd-3a90-4038-b822-273887b3de7f",
		Occassion:          "VA Occasion",
	})

	if err != nil {
		return nil, err
	}

	return b2cPayment, nil
}

// Simulate customer to business payments -.
func simulateC2BPayment(darajaService *daraja.DarajaService) (*daraja.C2BSimulateResponse, error) {
	simulateResponse, err := darajaService.C2BSimulate(daraja.C2BSimulateRequestBody{
		ShortCode:     600982,
		CommandID:     "CustomerPayBillOnline",
		Amount:        1,
		Msisdn:        254708374149,
		BillRefNumber: "VIrtual Account",
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
		TransactionType:   "CustomerBuyGoodsOnline",
		Amount:            "1",
		PartyA:            "254728922269",
		PartyB:            "174379",
		PhoneNumber:       "254728922269",
		CallBackURL:       "https://webhook.site/7da5ccfd-3a90-4038-b822-273887b3de7f",
		AccountReference:  "999200200",
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
