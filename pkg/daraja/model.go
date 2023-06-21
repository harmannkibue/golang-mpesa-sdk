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

// TransactionStatusRequestBody  checking the status of a transaction based of the M-Pesa receipt number -.
type TransactionStatusRequestBody struct {
	Initiator          string `json:"Initiator"`
	SecurityCredential string `json:"SecurityCredential"`
	CommandID          string `json:"CommandID"`
	TransactionID      string `json:"TransactionID"`
	PartyA             int    `json:"PartyA"`
	IdentifierType     int    `json:"IdentifierType"`
	ResultURL          string `json:"ResultURL"`
	QueueTimeOutURL    string `json:"QueueTimeOutURL"`
	Remarks            string `json:"Remarks"`
	Occassion          string `json:"Occassion"`
}

// TransactionStatusResponseBody response for initiating transaction status check -.
type TransactionStatusResponseBody struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ConversationID           string `json:"ConversationID"`
	ResponseCode             string `json:"ResponseCode"`
	ResponseDescription      string `json:"ResponseDescription"`
}

// TransactionReversalRequestBody request body for reversing C2B payments -.
type TransactionReversalRequestBody struct {
	Initiator              string `json:"Initiator"`
	SecurityCredential     string `json:"SecurityCredential"`
	CommandID              string `json:"CommandID"`
	TransactionID          string `json:"TransactionID"`
	Amount                 int    `json:"Amount"`
	ReceiverParty          int    `json:"ReceiverParty"`
	ReceiverIdentifierType int    `json:"ReceiverIdentifierType"`
	ResultURL              string `json:"ResultURL"`
	QueueTimeOutURL        string `json:"QueueTimeOutURL"`
	Remarks                string `json:"Remarks"`
	Occassion              string `json:"Occassion"`
}

// TransactionReversalResponseBody response after initiating C2B transaction -.
type TransactionReversalResponseBody struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ConversationID           string `json:"ConversationID"`
	ResponseCode             string `json:"ResponseCode"`
	ResponseDescription      string `json:"ResponseDescription"`
}
