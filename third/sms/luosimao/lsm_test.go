package luosimao

import (
	"fmt"
	"testing"
	"time"

	"myproject/types"
)

var (
	// secretKey = "e7709b5dfc2ccd28359eff07f08eba99"
	secretKey = "e7709b5dfc2ccd28359eff07f08eba82"

	apiKeys = map[string]string{
		"api": "key-" + secretKey,
	}
	format = "json"
)

func newLSM() *LSM {
	urls := map[string]string{
		LsmSendURLKey:      "http://sms-api.luosimao.com/v1/send.json",
		LsmSendBatchURLKey: "http://sms-api.luosimao.com/v1/send_batch.json",
		LsmStatusURLKey:    "http://sms-api.luosimao.com/v1/status.json",
	}
	options := &types.SMSOptions{
		Format: "json",
		Key:    "api",
		Pwd:    secretKey,
	}
	return New(urls, options)
}

func Test_NewUrlsPanic(t *testing.T) {
	options := &types.SMSOptions{}
	defer func() {
		if err := recover(); err != nil {
			t.Logf("test error new: %v", err)
		}
	}()
	lsm := New(nil, options)
	t.Logf("lsm: %v", lsm)
}

func Test_NewOptionsPanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Logf("test error new: %v", err)
		}
	}()
	lsm := New(nil, nil)
	t.Logf("lsm: %v", lsm)
}
func Test_New(t *testing.T) {
	lsm := newLSM()
	t.Logf("lsm: %v", lsm)
	options := &types.SMSOptions{
		Format: "json",
		Key:    "api",
		Pwd:    "secretKey",
	}
	lsm = New(nil, options)
	t.Logf("lsm: %v", lsm)

	lsm = New(nil, options)
	t.Logf("lsm: %v", lsm)

	options.Format = "xlms"
	lsm = New(nil, options)
	t.Logf("lsm: %v", lsm)

}

func Test_Send(t *testing.T) {
	lsm := newLSM()
	t.Logf("lsm: %v", lsm)
	msg := fmt.Sprintf("您的验证码是 %v，请尽快完成验证。【通客时代】", time.Now().Format("20060102150405"))
	res, err := lsm.Send("18310422105", msg)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("send: %v, receive: %#v", msg, res)
}

func Test_SendBatch(t *testing.T) {
	lsm := newLSM()
	t.Logf("lsm: %v", lsm)
	msg := fmt.Sprintf("您的验证码是 %v，请尽快完成验证。【通客时代】", time.Now().Format("20060102150405"))
	mobileList := "18310422105,18310422105"
	res, err := lsm.SendBatch(mobileList, msg, "")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("sendBatch: mobileList=%v, msg=%v, receive: %#v", mobileList, msg, res)
}

func Test_Status(t *testing.T) {
	lsm := newLSM()
	t.Logf("lsm: %v", lsm)
	res, err := lsm.Status()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("status: receive: %#v", res)
}
