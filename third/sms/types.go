package sms

// SendResponse response for send or sendBatch
type SendResponse struct {
	Error   int    `json:"error"`              // required error code
	Msg     string `json:"msg"`                // required error msg
	BatchID string `json:"batch_id,omitempty"` // only in send
	Hit     string `json:"hit,omitempty"`      // only in send sensitive word hit
}

// StatusResponse response for status
type StatusResponse struct {
	Error   int    `json:"error"`   // required error code
	Deposit string `json:"deposit"` // required balance
}
