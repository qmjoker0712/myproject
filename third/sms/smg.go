package sms

// SMS sms interface
type SMS interface {
	// Send send message POST
	Send(mobile, message string) (*SendResponse, error)

	// SendBatch send message batch POST
	SendBatch(mobileList, message, time string) (*SendResponse, error)

	// Status get the sms accounnt status GET
	Status() (*StatusResponse, error) // balance
}
