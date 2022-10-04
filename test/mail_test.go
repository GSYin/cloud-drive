package test

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestSendMail(t *testing.T) {
	e := email.NewEmail()
	e.From = "Ryan Gee <wencungsy@126.com>"
	e.To = []string{"wencungsy@163.com"}
	e.Subject = "test code"
	e.HTML = []byte("<p>【Cloud Drive】 验证码:</p>  <h2>823881</h2> 您正在进行邮箱登录，切勿将验证码泄露于他人，验证码10分钟内有效。")
	err := e.SendWithTLS("smtp.126.com:465", smtp.PlainAuth("", "wencungsy@126.com", "your mail password or token", "smtp.126.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtp.126.com"})
	if err != nil {
		t.Fatal(err)
	}
}
