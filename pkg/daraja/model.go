package daraja

// RegisterC2BURLBody is a model -.
type RegisterC2BURLBody struct {
	ShortCode       string
	ResponseType    string
	ConfirmationURL string
	ValidationURL   string
}

// RegisterC2BURLResponse URL response body -.
type RegisterC2BURLResponse struct {
	OriginatorCoversationID string `json:"OriginatorCoversationID"`
	ResponseCode            string `json:"ResponseCode"`
	ResponseDescription     string `json:"ResponseDescription"`
}

// STKPushBody request body -.
type STKPushBody struct {
	BusinessShortCode string `json:"BusinessShortCode"`
	TransactionType   string `json:"TransactionType"`
	Amount            string `json:"Amount"`
	PartyA            string `json:"PartyA"`
	PartyB            string `json:"PartyB"`
	PhoneNumber       string `json:"PhoneNumber"`
	CallBackURL       string `json:"CallBackURL"`
	AccountReference  string `json:"AccountReference"`
	TransactionDesc   string `json:"TransactionDesc"`
}

// StkPushResponse STKPush response body -.
type StkPushResponse struct {
	MerchantRequestID   string `json:"MerchantRequestID,omitempty"`
	CheckoutRequestID   string `json:"CheckoutRequestID,omitempty"`
	ResponseCode        string `json:"ResponseCode,omitempty"`
	ResponseDescription string `json:"ResponseDescription,omitempty"`
	CustomerMessage     string `json:"CustomerMessage,omitempty"`
	RequestId           string `json:"requestId,omitempty"`
	ErrorCode           string `json:"errorCode,omitempty"`
	ErrorMessage        string `json:"errorMessage,omitempty"`
}

// C2BSimulateRequestBody --> This is the request body for simulating a C2B payment -.
type C2BSimulateRequestBody struct {
	ShortCode     int    `json:"ShortCode"`
	CommandID     string `json:"CommandID"`
	Amount        int    `json:"Amount"`
	Msisdn        int64  `json:"Msisdn"`
	BillRefNumber string `json:"BillRefNumber"`
}

// C2BSimulateResponse -.
type C2BSimulateResponse struct {
	OriginatorCoversationID string `json:"OriginatorCoversationID"`
	ResponseCode            string `json:"ResponseCode"`
	ResponseDescription     string `json:"ResponseDescription"`
}

// B2CRequestBody -.
type B2CRequestBody struct {
	InitiatorName      string `json:"InitiatorName"`
	SecurityCredential string `json:"SecurityCredential"`
	CommandID          string `json:"CommandID"`
	Amount             int    `json:"Amount"`
	PartyA             int    `json:"PartyA"`
	PartyB             int64  `json:"PartyB"`
	Remarks            string `json:"Remarks"`
	QueueTimeOutURL    string `json:"QueueTimeOutURL"`
	ResultURL          string `json:"ResultURL"`
	Occassion          string `json:"Occassion"`
}

// B2CResponseBody -.
type B2CResponseBody struct {
	ConversationID           string `json:"ConversationID"`
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ResponseCode             string `json:"ResponseCode"`
	ResponseDescription      string `json:"ResponseDescription"`
}

//{
//"InitiatorName": "testapi",
//"SecurityCredential": "YL2cUjka1r1YSZsI/9ATatEtK8DruJwyx/leMHV0kr+bUTO6UciatgwkLuVYrHEWScIHYIlomBxXbm3813aVe4mpoUUU2/EAHJPcdbBQTRat1rcuyudv3TYkBR28mA+M+Y+NiyeLLGQkSxLUBPga5xk7z4tjgERoAPnsfmw/NijdS3ArZKZ+WZWPU+WMBvDU5ZouWJVPMzJ99/MCY0LY5TNUNKwYvlxxgjTgpUl4+5k1dQW8VEM61fxUl4aeJx8WKXOZt4bf+XmBLkdVwCaRyacyYoaK89sT53EadfNUU/whytJYcLr7rOHSQDxc5QqkbE4ekhklzjpuGk6ZRhkTTQ==",
//"CommandID": "BusinessPayment",
//"Amount": 1,
//"PartyA": 600426,
//"PartyB": 254708374149,
//"Remarks": "Testing B2C",
//"QueueTimeOutURL": "https://webhook.site/c882c5f6-4209-4f12-911b-85f13a69eb65",
//"ResultURL": "https://webhook.site/c882c5f6-4209-4f12-911b-85f13a69eb65",
//"Occassion": "Testing B2C"
//}
