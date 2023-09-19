package main

import (
	"github.com/harmannkibue/golang-mpesa-sdk/pkg/daraja"
	"log"
)

const (
	// Set environment variables for daraja before testing. Here I am using sandbox settings-.
	mpesaPassKey        = "0f2f587066b975699eac311b466c3cab733b4e73d9d20cacfde56be48d24bc6a"
	mpesaApiKey         = "xzbnAPtuYxchAZ7fEQKLnTpWUQeeADIC"
	mpesaConsumerSecret = "Sjr7WnjMZvqoo2ta"
)

func main() {
	// Initialize daraja service. PassKey is provided on your portal for sandbox,
	// while it is shared on your go live email from apisupport@safaricom.co.ke
	darajaService, err := daraja.New(mpesaApiKey, mpesaConsumerSecret, mpesaPassKey, daraja.PRODUCTION)

	if err != nil {
		log.Println("failed initializing safaricom daraja client ", err)
	}

	// Implements getting token from daraja, if not available in the memory cache -.
	//token, err := darajaToken(darajaService)
	//
	//if err != nil {
	//	log.Println("Error fetching token ", err.Error())
	//}
	//
	//log.Println("Daraja Token ", token)

	// Implements stk push service -.
	//stkRes, err := initiateStkPush(darajaService)
	//
	//if err != nil {
	//	log.Println("Error in stk push initiation ", err.Error())
	//}
	//log.Printf("STKPUSH response is %+v \n", stkRes)
	//
	//// Implements registering a confirmation and validation url.If response code is zero then it passed -.
	//confirmationResponseCode, err := registerConfirmationUrl(darajaService)
	//
	//if err != nil {
	//	log.Println("Error registering a URL ", err.Error())
	//}
	//
	//log.Println("Register URL response code ", confirmationResponseCode)
	//
	//// Simulate C2B transaction -.
	//simulateResponse, err := simulateC2BPayment(darajaService)
	//
	//if err != nil {
	//	log.Println("Error simulating C2B request: ", err.Error())
	//}
	//
	//fmt.Printf("C2B Response: %+v \n ", simulateResponse)
	//
	//balance, err := accountBalance(darajaService)
	//
	//if err != nil {
	//	fmt.Println("THE BALANCE ERROR ", err.Error())
	//}
	//
	//fmt.Printf("THE BALANCE RESPONSE %+v \n", balance)

	mpesaExpressStatus, err := queryMpesaExpressTransactionStatus(darajaService)

	if err != nil {
		log.Println("MPESA EXPRESS TRANSACTION STATUS RESPONSE ", err.Error())
	}

	log.Printf("MPESA EXPRESS TRANSACTION STATUS RESPONSE %+v \n", mpesaExpressStatus)

	//b2CStatus, err := queryTransactionStatus(darajaService)
	//
	//if err != nil {
	//	log.Println("B2C TRANSACTION STATUS RESPONSE ", err.Error())
	//}
	//
	//log.Printf("B2C TRANSACTION STATUS RESPONSE %+v \n", b2CStatus)

	//reversal, err := reverseC2BPayment(darajaService)
	//
	//if err != nil {
	//	log.Println("C2B REVERSAL ERROR ", err.Error())
	//}
	//
	//log.Printf("C2B REVERSAL RESPONSE %+v \n ", reversal)
	//
	//b2cRes, err := b2cPayment(darajaService)
	//
	//if err != nil {
	//	log.Println("B2C error: ", err)
	//}
	//
	//log.Printf("B2C response %+v \n", b2cRes)

}

// checking account balance for both B2C and C2B short codes -.
func accountBalance(darajaService *daraja.DarajaService) (*daraja.AccountBalanceResponseBody, error) {

	balance, err := darajaService.QueryAccountBalance(daraja.AccountBalanceRequestBody{
		Initiator:          "churpypayapi2",
		SecurityCredential: "m6FIrs4lHtTi8h+a347/G5/KWTWJObd+SuicTY43tQ+wPcinkuv+En82XvcYCITgT+RheQJYw2x1ICJviJhtjNM9no2eipC7Tag8C6fap7TjemT1FjPHzsgPWc3zWigBYcxijTJfVq9liUUf+hgfzlmglMmH9m7z4LaLRwIo/eAq+vHos8DC9bEPQvi2avnG2k8VsCb5R2PaQU//OTqSjfxrsRZrRdhHBjLrG8shlN+No4UhQAvBAvQl/22tUMfJjPG2SmbFcA+6o/Yr8XWV9+TRH8zimhyh86ibhERSVHst6kFYeoY7B+QIoDPvLZ2BKDyU7WP8g85HalM1slOttA==",
		CommandID:          "AccountBalance",
		PartyA:             3038361,
		// 1 for MSISDN 2 FOR TILL NUMBER 4 FOR ORGANISATION SHORT CODE -.
		IdentifierType:  4,
		Remarks:         "Churpy Balance",
		QueueTimeOutURL: "https://webhook.site/bbca16b1-fc3b-4a9f-9a91-14c08972657e",
		ResultURL:       "https://webhook.site/bbca16b1-fc3b-4a9f-9a91-14c08972657e",
	})

	if err != nil {
		return nil, err
	}

	return balance, nil
}

// Implements business to customer payment -.
func b2cPayment(darajaService *daraja.DarajaService) (*daraja.B2CResponseBody, error) {
	b2cPayment, err := darajaService.B2CPayment(daraja.B2CRequestBody{
		InitiatorName:      "churpypayapi2",
		SecurityCredential: "m6FIrs4lHtTi8h+a347/G5/KWTWJObd+SuicTY43tQ+wPcinkuv+En82XvcYCITgT+RheQJYw2x1ICJviJhtjNM9no2eipC7Tag8C6fap7TjemT1FjPHzsgPWc3zWigBYcxijTJfVq9liUUf+hgfzlmglMmH9m7z4LaLRwIo/eAq+vHos8DC9bEPQvi2avnG2k8VsCb5R2PaQU//OTqSjfxrsRZrRdhHBjLrG8shlN+No4UhQAvBAvQl/22tUMfJjPG2SmbFcA+6o/Yr8XWV9+TRH8zimhyh86ibhERSVHst6kFYeoY7B+QIoDPvLZ2BKDyU7WP8g85HalM1slOttA==",
		CommandID:          "SalaryPayment",
		Amount:             10,
		PartyA:             3038361,
		PartyB:             254728922269,
		Remarks:            "Payment from Business",
		QueueTimeOutURL:    "https://webhook.site/bbca16b1-fc3b-4a9f-9a91-14c08972657e",
		//QueueTimeOutURL: "https://staging.virtualaccounts.agoodjoke.lol/va-core/api/v1/payout/daraja-b2c-callback/",
		ResultURL: "https://staging.virtualaccounts.agoodjoke.lol/va-core/api/v1/payout/daraja-b2c-callback/",
		Occassion: "Occasion",
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
		BusinessShortCode: "7675771",
		TransactionType:   "CustomerBuyGoodsOnline",
		Amount:            "1",
		PartyA:            "254728922269",
		PartyB:            "5706975",
		PhoneNumber:       "254728922269",
		CallBackURL:       "https://webhook.site/bbca16b1-fc3b-4a9f-9a91-14c08972657e",
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

func queryMpesaExpressTransactionStatus(darajaService *daraja.DarajaService) (*daraja.MpesaExpressTransactionStatusQueryResponse, error) {
	status, err := darajaService.MpesaExpressTransactionStatus(daraja.MpesaExpressTransactionStatusQueryBodyComplete{
		BusinessShortCode: "7675771",
		Password:          darajaService.ApiPassKey,
		Timestamp:         "20230619204330",
		CheckoutRequestID: "ws_CO_19092023131241018728922269",
	})

	if err != nil {
		return nil, err
	}

	return status, nil

}

// Query transaction status for a organisation shortcode, MSISDN or a till number -.
func queryTransactionStatus(darajaService *daraja.DarajaService) (*daraja.TransactionStatusResponseBody, error) {

	status, err := darajaService.TransactionStatus(daraja.TransactionStatusRequestBody{
		// For payout -.
		Initiator:                "churpypayapi2",
		SecurityCredential:       "iloJka53jHU9nDLzjiO9aysfQbRfpgWfSUue/Zy61T97jZzvxHFyiFIGcIa02pFXxKLyYqyc5RkNpmUC4IfhdPEVdMfBheayjJZY6h7qDEkNW6qql208XAjEbj5XdLB+z9+wE2J6YCGMTrwEBjYFIqgDNT8jHi7AW9ptXTzC/NTpFoQa7awrYbzwLdIeFBSOW6E1jMhxwUm7O7YE3B2l3bKVMvEkdvG0VHBAU7hAO8dRdfhTz/siEwe4LOeYt7pOyvIoMxgcXNrxle2jnXkuqPSI0gjtJnhfmuxkHSzhumNrZic+lPM0p28COXVlXG0jRtCe9cOv9cDdX3PWC6NWGw==",
		CommandID:                "TransactionStatusQuery",
		OriginatorConversationID: "26315-110726387-2",
		TransactionID:            "",
		PartyA:                   3038361,
		// 1 for MSISDN ,2 FOR TILL NUMBER and 4 FOR ORGANISATION SHORT CODE -.
		IdentifierType:  4,
		ResultURL:       "https://webhook.site/bbca16b1-fc3b-4a9f-9a91-14c08972657e",
		QueueTimeOutURL: "https://webhook.site/bbca16b1-fc3b-4a9f-9a91-14c08972657e",
		Remarks:         "OK",
		Occassion:       "Ok",
	})

	if err != nil {
		return nil, err
	}

	return status, nil
}

// reversing a C2B payment -.
func reverseC2BPayment(darajaService *daraja.DarajaService) (*daraja.TransactionReversalResponseBody, error) {
	reversal, err := darajaService.C2BTransactionReversal(daraja.TransactionReversalRequestBody{
		Initiator:              "testapi",
		SecurityCredential:     "oDx3GjKUpc3LJyPMdjiy2Qy64b+Smfyc8xyPTjYQfpGhVngg8OATaXYla0YazHGtM8rqqlRwGiW30NDTezm81YBpEwCvIWTaR1YN3RmiPPvN+kF03BgX8eCJXVzV/99758nSsEKmleudOMmkegHaTrMOlfjQlcVSiS94u2ZvJejS0X5xpp2dPkplITmpLBh/EpMsB0fJLh7fcrtc8v0V/NJG6Zd6W3d2uB3S6zfJPbc4Iby52iYhAWwFOAbJhrTMVDHKLLCzFXZUZufPpntWcElNAtgEb7AA1Os2FbNyJcrCwT22ATQaU/VMJTjMgMB3Cgdw7Xyw+gMilJ+er/kJzA==",
		CommandID:              "TransactionReversal",
		TransactionID:          "OEI2AK4Q16",
		Amount:                 1,
		ReceiverParty:          600978,
		ReceiverIdentifierType: 11,
		ResultURL:              "https://webhook.site/7da5ccfd-3a90-4038-b822-273887b3de7f",
		QueueTimeOutURL:        "https://webhook.site/7da5ccfd-3a90-4038-b822-273887b3de7f",
		Remarks:                "REVERSAL REMARK",
		Occassion:              "REVERSAL OCCASION",
	})

	if err != nil {
		return nil, err
	}

	return reversal, nil
}
