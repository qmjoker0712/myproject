package emails

import (
	"errors"
	"net/smtp"
	"net/textproto"
	"sync"
	"time"

	"github.com/jordan-wright/email"

	"myproject/types"
)

// EmailPool global email pool
var EmailPool *email.Pool

// EmailSendTimeout email send timeout, default is 30, unit is second
var EmailSendTimeout = 30
var lock sync.RWMutex

// InitEmailPool if enable email need init it before send email
func InitEmailPool(emailsOptions *types.EmailOptions) (err error) {
	lock.RLock()
	defer lock.RUnlock()
	if emailsOptions == nil {
		return errors.New("emailsOptions config error")
	}

	// minum timeout is 5s
	if emailsOptions.Timeout > 5 {
		EmailSendTimeout = emailsOptions.Timeout
	}

	authConfig := emailsOptions.AuthConfig
	EmailPool, err = email.NewPool(emailsOptions.Address,
		emailsOptions.Count,
		smtp.PlainAuth(
			authConfig.Identity,
			authConfig.Username,
			authConfig.Password,
			authConfig.Host))

	return err
}

// InitSimpleEmailPool if enable email need init it before send email
func InitSimpleEmailPool(endpoint string, size int, auth smtp.Auth) (err error) {
	lock.RLock()
	defer lock.RUnlock()

	EmailPool, err = email.NewPool(endpoint, size, auth)
	return err
}

// GetEmailPool get email pool
func GetEmailPool() (pool *email.Pool, err error) {
	lock.RLock()
	defer lock.RUnlock()
	if EmailPool != nil {
		return EmailPool, nil
	}
	return nil, errors.New("email pool is nil")

}

// SendEmail send email use default or init timeout
func SendEmail(email *email.Email) error {
	return SendEmailWithTimeout(email, time.Second*time.Duration(EmailSendTimeout))
}

// SendEmailWithTimeout send email use input timeout
func SendEmailWithTimeout(email *email.Email, timeout time.Duration) error {
	if EmailPool == nil {
		return errors.New("Email pool has not been initialized yet")
	}

	return EmailPool.Send(email, timeout)
}

// GetDefaultEmail text and html choose one, if html and text set both, html will be send
// cc Carbon Copy, bcc Blind CarbonCopy
func GetDefaultEmail(from string, to []string, subject string, text string, html string, cc, bcc []string, sender string) *email.Email {
	return &email.Email{
		From:    from,
		To:      to,
		Cc:      cc,
		Bcc:     bcc,
		Subject: subject,
		Text:    []byte(text),
		HTML:    []byte(html),
		Sender:  sender,
		Headers: textproto.MIMEHeader{},
	}
}

// GetSimpleEmail get simple email
func GetSimpleEmail(from string, to []string, subject string, text string, html string) *email.Email {
	return &email.Email{
		From:    from,
		To:      to,
		Subject: subject,
		Text:    []byte(text),
		HTML:    []byte(html),
		Headers: textproto.MIMEHeader{},
	}
}

// AppendEmailCc append new cc emails to email
func AppendEmailCc(email *email.Email, cc []string) {
	email.Cc = append(email.Cc, cc...)
}

// SetEmailCc set new cc emails
func SetEmailCc(email *email.Email, cc []string) {
	email.Cc = []string{}
	email.Cc = append(email.Cc, cc...)
}

// AppendEmailBcc append new bcc emails to email
func AppendEmailBcc(email *email.Email, bcc []string) {
	email.Bcc = append(email.Bcc, bcc...)
}

// SetEmailBcc set new bcc emails
func SetEmailBcc(email *email.Email, bcc []string) {
	email.Bcc = []string{}
	email.Bcc = append(email.Bcc, bcc...)
}

// SetEmailSubject set email subject
func SetEmailSubject(email *email.Email, subject string) {
	email.Subject = subject
}

// SetEmailText set email text
func SetEmailText(email *email.Email, text string) {
	email.Text = []byte(text)
}

// SetEmailHTML set email html
func SetEmailHTML(email *email.Email, html string) {
	email.HTML = []byte(html)
}
