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
