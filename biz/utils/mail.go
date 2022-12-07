package utils

import (
	"github.com/lutasam/check_in_sys/biz/common"
	"gopkg.in/gomail.v2"
)

func SendMail(mailTo, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(GetConfigString("mail.address"), GetConfigString("mail.username")))
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(GetConfigString("mail.host"), 25, GetConfigString("mail.username"), GetConfigString("mail.auth"))
	err := d.DialAndSend(m)
	if err != nil {
		return common.EMAILSYSTEMERROR
	}
	return nil
}
