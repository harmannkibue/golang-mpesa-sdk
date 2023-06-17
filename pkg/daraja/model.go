package daraja

// RegisterC2BURL is a model
type RegisterC2BURL struct {
	ShortCode       string
	ResponseType    string
	ConfirmationURL string
	ValidationURL   string
}

type RegisterC2BURLResponse struct {
	OriginatorCoversationID string `json:"OriginatorCoversationID"`
	ResponseCode            string `json:"ResponseCode"`
	ResponseDescription     string `json:"ResponseDescription"`
}
