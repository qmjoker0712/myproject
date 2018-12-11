package luosimao

import (
	"encoding/json"

	"myproject/common"
	"myproject/third/sms"
	"myproject/types"
)

// LSM sms const
const (
	LsmFormatJSON = "json"
	LsmFormatXML  = "xml"

	LsmSendURLKey      = "send"
	LsmSendBatchURLKey = "sendBatch"
	LsmStatusURLKey    = "status"
)

// LSM URLS
var (
	LsmSendURL      = "http://sms-api.luosimao.com/v1/send"
	LsmSendBatchURL = "http://sms-api.luosimao.com/v1/send_batch"
	LsmStatusURL    = "http://sms-api.luosimao.com/v1/status"
)

// LSM luosimao sms API
type LSM struct {
	format  string
	URLs    map[string]string
	APIKeys map[string]string
}

// New create LSM instance
func New(urls map[string]string, options *types.SMSOptions) *LSM {
	if options == nil {
		panic("error options for sms")
	}
	if options.Key == "" || options.Pwd == "" {
		panic("error key or pwd for sms")
	}

	format := options.Format
	if format != LsmFormatJSON && format != LsmFormatXML {
		format = LsmFormatJSON
		options.Format = format
	}
	if urls == nil || len(urls) == 0 {
		urls = make(map[string]string)
		urls = map[string]string{
			LsmSendURLKey:      getURL(LsmSendURL, format),
			LsmSendBatchURLKey: getURL(LsmSendBatchURL, format),
			LsmStatusURLKey:    getURL(LsmSendURL, format),
		}
	}
	return &LSM{
		format: format,
		URLs:   urls,
		APIKeys: map[string]string{
			options.Key: options.Pwd,
		},
	}
}

func getURL(url, format string) string {
	return url + "." + format
}

// Send send message POST
func (l *LSM) Send(mobile, message string) (*sms.SendResponse, error) {
	url := l.URLs[LsmSendURLKey]
	params := map[string]interface{}{
		"mobile":  mobile,
		"message": message,
	}
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	res, err := common.HTTPPostRequestWithBasicAuth(url, params, headers, l.APIKeys)
	if err != nil {
		return nil, err
	}
	var response sms.SendResponse
	json.Unmarshal([]byte(res), &response)
	return &response, nil
}

// SendBatch send message POST
func (l *LSM) SendBatch(mobileList, message, time string) (*sms.SendResponse, error) {
	url := l.URLs[LsmSendBatchURLKey]
	params := map[string]interface{}{
		"mobile_list": mobileList,
		"message":     message,
		"time":        time,
	}
	if time == "" {
		delete(params, "time")
	}

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	res, err := common.HTTPPostRequestWithBasicAuth(url, params, headers, l.APIKeys)
	if err != nil {
		return nil, err
	}
	var response sms.SendResponse
	json.Unmarshal([]byte(res), &response)
	return &response, nil
}

// Status get the account status
func (l *LSM) Status() (*sms.StatusResponse, error) {
	url := l.URLs[LsmStatusURLKey]
	res, err := common.HTTPGetRequestWithBasicAuth(url, nil, nil, l.APIKeys)
	if err != nil {
		return nil, err
	}
	var response sms.StatusResponse
	json.Unmarshal([]byte(res), &response)
	return &response, nil
}
