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

// AccountBalanceRequestBody for checking the account balance -.
type AccountBalanceRequestBody struct {
	Initiator          string `json:"Initiator"`
	SecurityCredential string `json:"SecurityCredential"`
	CommandID          string `json:"CommandID"`
	PartyA             int    `json:"PartyA"`
	IdentifierType     int    `json:"IdentifierType"`
	Remarks            string `json:"Remarks"`
	QueueTimeOutURL    string `json:"QueueTimeOutURL"`
	ResultURL          string `json:"ResultURL"`
}

// AccountBalanceResponseBody acknowledgement response from mpesa -.
type AccountBalanceResponseBody struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ConversationID           string `json:"ConversationID"`
	ResponseCode             string `json:"ResponseCode"`
	ResponseDescription      string `json:"ResponseDescription"`
}

//{
//"Initiator": "testapi",
//"SecurityCredential": "J0Pk2QHyUJ0b+iuPDDukHSMdRiLHp916TG81WvyBKr2GQIpUW5blHqt6LRVNiD1VIIWgtjYwMqwUgDfc29wKKs3DiTrcjDF/KdhwcSCYzmxiZg2HSixQ0sS+UCbQwv/KxJ+Ugd+9hCMULeW7NdfZ5ZK04jmUiUw/e2i1hjNnSAKpo9SaPuVLM8OCs9tHbfZxM8PQplb+/r3uFzxzryd1yf2WjrFecrOLOnp7UFCbZhzhdoL/um+1UxvbFYyGdfpC+PaOPr1P9IT0zchVJvCB78ovkLplT4vZadZJ6dY7EaqTX4Bl+MaNBRaEVr+sMIZhUPbRpgzlgaOSyQe13P57NQ==",
//"CommandID": "AccountBalance",
//"PartyA": 174379,
//"IdentifierType": 4,
//"Remarks": "Checking balance",
//"QueueTimeOutURL": "https://webhook.site/7da5ccfd-3a90-4038-b822-273887b3de7f",
//"ResultURL": "https://webhook.site/7da5ccfd-3a90-4038-b822-273887b3de7f",
//}
