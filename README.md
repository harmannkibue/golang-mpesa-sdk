# Daraja Golang SDK 

The SDK provides fast, secure and easy integration for [Safaricom MPESA Daraja API](https://developer.safaricom.co.ke/apis-explorer) for golang applications.

## Installing
You can install the package by running:

```
go get github.com/harmannkibue/golang-mpesa-sdk
```

## Technology
1. Golang v 1.17 or higher
2. github.com/patrickmn/go-cache - For caching of authentication token

## Getting started
Obtain mpesaApiKey, mpesaConsumerSecret, mpesaPassKey from [Safaricom MPESA Daraja API](https://developer.safaricom.co.ke/apis-explorer). Also generate credentials either for sandbox or production.

```
const (
	// Sandbox or production credentials 
	mpesaApiKey         = "xzbnAPtuYxchAZ7fEQKLnTpWUQeeADIC"
	mpesaConsumerSecret = "Sjr7WnjMZvqoo2ta"
	mpesaPassKey        = "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919"
	securityCredential = "C7agVsde6iC4ku48FfcuOgc1iNtzIK9j1oX60pJeqouVo6XPlJMkR/Pqc22UXz3qu2T0CLkVDORsvJSAxBuZk0KAtro5vbdMyvn5aUYQ0uGN1RY=="
)
```

### MPESAExpress/Network Initiated (STKPush)
This api initiates payment prompt on customer sim tool kit.
```go
package main

import (
	"log"
	"github.com/harmannkibue/golang-mpesa-sdk/pkg/daraja"
)

const (
	mpesaApiKey         = "xzbnAPtuYxchAZ7fEQKLnTpWUQeeADIC"
	mpesaConsumerSecret = "Sjr7WnjMZvqoo2ta"
	mpesaPassKey        = "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919"
)

func main() {

	darajaService, err := daraja.New(mpesaApiKey, mpesaConsumerSecret, mpesaPassKey, daraja.SANDBOX)

	if err != nil {
		log.Println("failed initializing safaricom daraja client ", err)
	}

	stkRes, err := darajaService.InitiateStkPush(daraja.STKPushBody{
		BusinessShortCode: "174379",
		TransactionType:   "CustomerBuyGoodsOnline",
		Amount:            "1",
		PartyA:            "254728922469",
		PartyB:            "174379",
		PhoneNumber:       "254728922269",
		CallBackURL:       "https://webhook.site/7da5ccfd-3a90-4038-b822-273887b3de7f",
		AccountReference:  "999200200",
		TransactionDesc:   "Daraja sdk testing STK push",
	})

	if err != nil {
		log.Println("STK ERROR ", err)
	}
	log.Printf("STK push response %+v \n", stkRes)

}

```

### C2B URL register
Register confirmation and validation urls, however you need to request validation url setup from safaricom daraja.

```go
package main

import (
	"log"
	"github.com/harmannkibue/golang-mpesa-sdk/pkg/daraja"
)

const (
	mpesaApiKey         = "xzbnAPtuYxchAZ7fEQKLnTpWUQeeADIC"
	mpesaConsumerSecret = "Sjr7WnjMZvqoo2ta"
	mpesaPassKey        = "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919"
)

func main() {
	darajaService, err := daraja.New(mpesaApiKey, mpesaConsumerSecret, mpesaPassKey, daraja.SANDBOX)

	if err != nil {
		log.Println("failed initializing safaricom daraja client ", err)
	}
	
	regUrl, err := darajaService.C2BRegisterURL(daraja.RegisterC2BURLBody{
		ShortCode:       "600989",
		ResponseType:    "Completed",
		ConfirmationURL: "https://webhook.site/c882c5f6-4209-4f12-911b-85f13a69eb65",
		ValidationURL:   "https://webhook.site/c882c5f6-4209-4f12-911b-85f13a69eb65",
	})

	if err != nil {
		log.Println("Register URL error: ", err)
	}
	
	log.Printf("Register URL response %+v \n", regUrl)

}

```
### Simulate C2B payment for sandbox

```go
package main

import (
	"log"
	"github.com/harmannkibue/golang-mpesa-sdk/pkg/daraja"
)

const (
	mpesaApiKey         = "xzbnAPtuYxchAZ7fEQKLnTpWUQeeADIC"
	mpesaConsumerSecret = "Sjr7WnjMZvqoo2ta"
	mpesaPassKey        = "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919"
)

func main() {

	darajaService, err := daraja.New(mpesaApiKey, mpesaConsumerSecret, mpesaPassKey, daraja.SANDBOX)

	if err != nil {
		log.Println("failed initializing safaricom daraja client ", err)
	}

	simulateResponse, err := darajaService.C2BSimulate(daraja.C2BSimulateRequestBody{
		ShortCode:     600982,
		CommandID:     "CustomerPayBillOnline",
		Amount:        1,
		Msisdn:        254708374149,
		BillRefNumber: "VIrtual Account",
	})

	if err != nil {
		log.Println("Simulate C2B error: ", err)
	}
    log.Printf("Simulate response %+v \n", simulateResponse)
}
```

### B2C
This api allows you to do M-Pesa Transaction from company to client.

```go
package main

import (
	"log"
	"github.com/AndroidStudyOpenSource/mpesa-api-go"
)

const (
	appKey    = ""
	appSecret = ""
)

func main() {

	svc, err := mpesa.New(appKey, appSecret, mpesa.SANDBOX)
	if err != nil {
		panic(err)
	}

	res, err := svc.B2CRequest(mpesa.B2C{
		InitiatorName:      "",
		SecurityCredential: "",
		CommandID:          "",
		Amount:             "",
		PartyA:             "",
		PartyB:             "",
		Remarks:            "",
		QueueTimeOutURL:    "",
		ResultURL:          "",
		Occassion:          "",
	})

	if err != nil {
		log.Println(err)
	}
	log.Println(res)

}
```

### B2B
This api allows you to do M-Pesa Transaction from one company to another.

```go
package main

import (
	"log"
	"github.com/AndroidStudyOpenSource/mpesa-api-go"
)

const (
	appKey    = ""
	appSecret = ""
)

func main() {

	svc, err := mpesa.New(appKey, appSecret, mpesa.SANDBOX)
	if err != nil {
		panic(err)
	}

	res, err := svc.B2BRequest(mpesa.B2B{
		Initiator:              "",
		SecurityCredential:     "",
		CommandID:              "",
		SenderIdentifierType:   "",
		RecieverIdentifierType: "",
		Remarks:                "",
		Amount:                 "",
		PartyA:                 "",
		PartyB:                 "",
		AccountReference:       "",
		QueueTimeOutURL:        "",
		ResultURL:              "",
	})

	if err != nil {
		log.Println(err)
	}
	log.Println(res)

}

```

### Account Balance
This api allows you to do balance inquiry.

```go
package main

import (
	"log"
	"github.com/AndroidStudyOpenSource/mpesa-api-go"
)

const (
	appKey    = ""
	appSecret = ""
)

func main() {

	svc, err := mpesa.New(appKey, appSecret, mpesa.SANDBOX)
	if err != nil {
		panic(err)
	}

	res, err := svc.BalanceInquiry(mpesa.BalanceInquiry{
		Initiator:          "",
		SecurityCredential: "",
		CommandID:          "",
		PartyA:             "",
		IdentifierType:     "",
		Remarks:            "",
		QueueTimeOutURL:    "",
		ResultURL:          "",
	})

	if err != nil {
		log.Println(err)
	}
	log.Println(res)

}
```

### Transaction Status
This api allows you to check the status of transaction.

### Reversal
This api allows you to do a transaction reversal

```go
package main

import (
	"log"
	"github.com/AndroidStudyOpenSource/mpesa-api-go"
)

const (
	appKey    = ""
	appSecret = ""
)

func main() {

	svc, err := mpesa.New(appKey, appSecret, mpesa.SANDBOX)
	if err != nil {
		panic(err)
	}

	res, err := svc.Reversal(mpesa.Reversal{
		Initiator:              "",
		SecurityCredential:     "",
		CommandID:              "",
		TransactionID:          "",
		Amount:                 "",
		ReceiverParty:          "",
		ReceiverIdentifierType: "",
		QueueTimeOutURL:        "",
		ResultURL:              "",
		Remarks:                "",
		Occassion:              "",
	})

	if err != nil {
		log.Println(err)
	}
	log.Println(res)

}
```

### Making a Contribution

Before making significant changes, open an issue to discuss your proposal or bug fix. This allows the maintainers and community to provide feedback and ensure your work aligns with project goals.
Commit your changes using descriptive and concise commit messages.
Push your changes to your forked repository.
Submit a pull request (PR) to the main repository's master or relevant branch.

### License

```text
MIT License

Copyright (c) 2023 Harman Kibue

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NON INFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
