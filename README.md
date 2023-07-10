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
	// Sandbox or production credentials. Note store sucurely as environment vartiables
	mpesaApiKey         = "xzbnAPtuYxchAZ7fEQKLnTpWUQeeADIC"
	mpesaConsumerSecret = "Sjr7WnjMZvqoo2ta"
	mpesaPassKey        = "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919"
	securityCredential = "C7agVsde6iC4ku48FfcuOgc1iNtzIK9j1oX60pJeqouVo6XPlJMkR/Pqc22UXz3qu2T0CLkVDORsvJSAxBuZk0KAtro5vbdMyvn5aUYQ0uGN1RY=="
)
```

Then declare service for either production or sandbox
```
darajaService, err := daraja.New(mpesaApiKey, mpesaConsumerSecret, mpesaPassKey, daraja.SANDBOX)

	if err != nil {
		log.Println("failed initializing safaricom daraja client ", err)
	}
```

### Network Initiated/Mpesa Express (STKPush)
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

### C2B register URL
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
		BillRefNumber: "Account number",
	})

	if err != nil {
		log.Println("Simulate C2B error: ", err)
	}
    log.Printf("Simulate response %+v \n", simulateResponse)
}
```

### Transaction Status
This api allows you to track transaction status for payments.

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

	statusResponse, err := darajaService.TransactionStatus(daraja.TransactionStatusRequestBody{
		Initiator:          "testapi",
		SecurityCredential: "UKCrm4IVKWEoW640M3pUHS4hZ2ynDpz+LT6c+acBK28TOMULxVhMP0YM2FNCh2QXx+m6HR8iLNsR0bfbIB1kpvNhciKUrn7Glp4f7UNPF8mHXgNsa/09+i7X8+JUy7tQLEOoPE/xCWBOh2ofBq8N+lX77RUAxDp9HC8Nj6nN6kH07Ygmz7NnRd/dlayqcFKV4UNP/nQAV8lum2HSh9xRBnlexcziYipt/d293qrSSvXtAfz+lmgzzbzwML02zlCQxXS2YQjTluQWzRgxkl+9aCCs51a5BWppTE6iYd8qcMlX/+hMZvl2D9LjQKwisSKJsWP2MtxFxG86DRpwI41I4A==",
		CommandID:          "TransactionStatusQuery",
		TransactionID:      "RFL5LEUJ4H",
		PartyA:             600989,
		// 1 for MSISDN 2 FOR TILL NUMBER 4 FOR ORGANISATION SHORT CODE -.
		IdentifierType:  2,
		ResultURL:       "https://webhook.site/bbca16b1-fc3b-4a9f-9a91-14c08972657e",
		QueueTimeOutURL: "https://webhook.site/bbca16b1-fc3b-4a9f-9a91-14c08972657e",
		Remarks:         "TRANSACTION STATUS REMARKS",
		Occassion:       "TRANSACTION STATUS  OCCASION",
	})

	if err != nil {
		log.Println("Status error: ", err)
	}
	log.Printf("Status response %+v \n", statusResponse)

}
```

### B2C payment
This api allows you to do M-Pesa Transaction from company to client.

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

	b2cResponse, err := darajaService.B2CPayment(daraja.B2CRequestBody{
		InitiatorName:      "testapi",
		SecurityCredential: "UKCrm4IVKWEoW640M3pUHS4hZ2ynDpz+LT6c+acBK28TOMULxVhMP0YM2FNCh2QXx+m6HR8iLNsR0bfbIB1kpvNhciKUrn7Glp4f7UNPF8mHXgNsa/09+i7X8+JUy7tQLEOoPE/xCWBOh2ofBq8N+lX77RUAxDp9HC8Nj6nN6kH07Ygmz7NnRd/dlayqcFKV4UNP/nQAV8lum2HSh9xRBnlexcziYipt/d293qrSSvXtAfz+lmgzzbzwML02zlCQxXS2YQjTluQWzRgxkl+9aCCs51a5BWppTE6iYd8qcMlX/+hMZvl2D9LjQKwisSKJsWP2MtxFxG86DRpwI41I4A==",
		CommandID:          "SalaryPayment",
		Amount:             1,
		PartyA:             600998,
		PartyB:             254728920369,
		Remarks:            "Payment from Business",
		QueueTimeOutURL:    "https://webhook.site/7da5ccfd-3a90-4038-b822-273887b3de7f",
		ResultURL:          "https://webhook.site/7da5ccfd-3a90-4038-b822-273887b3de7f",
		Occassion:          "Occasion",
	})

	if err != nil {
		log.Println("B2C error: ", err)
	}
	log.Printf("B2C response %+v \n", b2cResponse)
}
```

### Account Balance
This api allows you to do balance inquiry.

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

	balanceResponse, err := darajaService.QueryAccountBalance(daraja.AccountBalanceRequestBody{
		Initiator:          "testapi",
		SecurityCredential: "UKCrm4IVKWEoW640M3pUHS4hZ2ynDpz+LT6c+acBK28TOMULxVhMP0YM2FNCh2QXx+m6HR8iLNsR0bfbIB1kpvNhciKUrn7Glp4f7UNPF8mHXgNsa/09+i7X8+JUy7tQLEOoPE/xCWBOh2ofBq8N+lX77RUAxDp9HC8Nj6nN6kH07Ygmz7NnRd/dlayqcFKV4UNP/nQAV8lum2HSh9xRBnlexcziYipt/d293qrSSvXtAfz+lmgzzbzwML02zlCQxXS2YQjTluQWzRgxkl+9aCCs51a5BWppTE6iYd8qcMlX/+hMZvl2D9LjQKwisSKJsWP2MtxFxG86DRpwI41I4A==",
		CommandID:          "AccountBalance",
		PartyA:             600991,
		// 1 for MSISDN 2 FOR TILL NUMBER 4 FOR ORGANISATION SHORT CODE -.
		IdentifierType:  2,
		Remarks:         "Churpy Balance",
		QueueTimeOutURL: "https://webhook.site/bbca16b1-fc3b-4a9f-9a91-14c08972657e",
		ResultURL:       "https://webhook.site/bbca16b1-fc3b-4a9f-9a91-14c08972657e",
	})

	if err != nil {
		log.Println("Balance error : ", err)
	}
	
	log.Printf("Balance response %+v \n", balanceResponse)

}
```


### Reversal
This api allows you to do a transaction reversal for C2B payment

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

	reversalResponse, err := darajaService.C2BTransactionReversal(daraja.TransactionReversalRequestBody{
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
    log.Printf("Reversal response %+v \n ", reversalResponse)
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
