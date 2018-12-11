package emails

import (
	"fmt"
	"testing"
	"time"

	"github.com/jordan-wright/email"

	"myproject/common"
	"myproject/types"
)

var (
	identity            = ""
	username            = "wangxin@zsbatech.com"
	password            = "Wangxin123"
	host                = "smtp.exmail.qq.com"
	serverAddr          = "smtp.exmail.qq.com:25"
	defaultEmailOptions = &types.EmailOptions{
		Address: serverAddr,
		Count:   4,
		AuthConfig: types.EmailAuthConfig{
			Identity: identity,
			Username: username,
			Password: password,
			Host:     host,
		},
	}

	from    = fmt.Sprintf("%s <%s>", "wangxin", username)
	toAddr  = []string{username}
	cc      = []string{}
	bcc     = []string{"wx_cs_db_88@163.com"}
	subject = "Email Send Test"
	text    = "Text Body is, of course, supported!"
	html    = "<h1>You win 100,000,000,000,000 $ in powerball! <a href=\"http://www.zsbatech.com\">welcome visit</a></h1>"
	sender  = fmt.Sprintf("%s <%s>", "tester", username)
)

func Test_GetEmailPool(t *testing.T) {
	if err := InitEmailPool(defaultEmailOptions); err != nil {
		t.Error(err)
	}
	emailPool, err := GetEmailPool()
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("GetEmailPool: %#v", emailPool)
	}

	EmailPool = nil
	emailPool, err = GetEmailPool()
	if err != nil {
		t.Logf("GetEmailPool error %v", err)
	} else {
		t.Errorf("GetEmailPool: %#v", emailPool)
	}

}

func Test_InitEmailPool(t *testing.T) {
	if err := InitEmailPool(defaultEmailOptions); err != nil {
		t.Error(err)
	}
	defaultEmailOptions.Timeout = 15
	if err := InitEmailPool(defaultEmailOptions); err != nil {
		t.Error(err)
	}
}

func Test_InitEmailPoolError(t *testing.T) {
	var test_defaultEmailOptions *types.EmailOptions = nil
	if err := InitEmailPool(test_defaultEmailOptions); err == nil {
		t.Error("Test_InitEmailPoolError should err")
	}

}

func getDefaultEmail() *email.Email {
	return GetDefaultEmail(from, toAddr, subject, text, html, cc, bcc, sender)
}

func getSimpleEmail() *email.Email {
	return GetSimpleEmail(from, toAddr, subject, text, html)
}

func Test_GetSimpleEmail(t *testing.T) {
	email := getSimpleEmail()
	if email == nil {
		t.Errorf("GetSimpleEmail get nil email")
	}
}

func Test_SendEmail(t *testing.T) {
	email := getDefaultEmail()
	if err := SendEmail(email); err != nil {
		t.Logf("SendEmail, err %v", err)
	}
}

func Test_SendEmailWithTimeout(t *testing.T) {
	email := getDefaultEmail()
	timeout := time.Second * 5
	if err := SendEmailWithTimeout(email, timeout); err != nil {
		t.Logf("SendEmailWithTimeout with timeout %v, err %v", timeout, err)
	}
}

func Test_AppendEmailCc(t *testing.T) {
	email := getDefaultEmail()
	cc1 := email.Cc
	t.Log(cc1)
	AppendEmailCc(email, []string{"test@123.com"})
	cc2 := email.Cc
	t.Log(cc2)
	if len(cc1) == len(cc2) {
		t.Error("AppendEmailCc error")
	}
}

func Test_SetEmailCc(t *testing.T) {
	email := getDefaultEmail()
	cc1 := email.Cc
	t.Log(cc1)
	SetEmailCc(email, []string{"test@123.com"})
	cc2 := email.Cc
	if ok := common.StringSliceEqual(cc1, cc2); ok {
		t.Errorf("SetEmailCc from %v to %v", cc1, cc2)
	}
}

func Test_AppendEmailBcc(t *testing.T) {
	email := getDefaultEmail()
	bcc1 := email.Bcc
	t.Log(bcc1)
	AppendEmailBcc(email, []string{"test@123.com"})
	bcc2 := email.Bcc
	t.Log(bcc2)
}

func Test_SetEmailBcc(t *testing.T) {
	email := getDefaultEmail()
	bcc1 := email.Bcc
	t.Log(bcc1)
	SetEmailBcc(email, []string{"test@123.com"})
	bcc2 := email.Bcc
	t.Log(bcc2)
	if ok := common.StringSliceEqual(bcc1, bcc2); ok {
		t.Errorf("SetEmailBcc from %v to %v", bcc1, bcc2)
	}
}

func Test_SetEmailSubject(t *testing.T) {
	email := getDefaultEmail()
	subject1 := email.Subject
	SetEmailSubject(email, "new email subject for test"+time.Now().String())
	subject2 := email.Subject
	if subject1 == subject2 {
		t.Errorf("SetEmailSubject from %v to %v", subject1, subject2)
	}
}

func Test_SetEmailText(t *testing.T) {
	email := getDefaultEmail()
	text1 := email.Text
	SetEmailText(email, "new email text for test"+time.Now().String())
	text2 := email.Text
	if string(text1) == string(text2) {
		t.Errorf("SetEmailText from %v to %v", text1, text2)
	}
}

func Test_SetEmailHTML(t *testing.T) {
	email := getDefaultEmail()
	html1 := email.HTML
	SetEmailHTML(email, "new email html for test"+time.Now().String())
	html2 := email.HTML
	if string(html1) == string(html2) {
		t.Errorf("SetEmailHTML from %v to %v", html1, html2)
	}
}
