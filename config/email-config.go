package config

import (
	"common-web-framework/common"
	"common-web-framework/helper"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

type EmailConfig struct {
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	Host     string `yaml:"host" json:"host"`
	Addr     string `yaml:"addr" json:"addr"`
}

func (ef EmailConfig) SendEmail(to string, subject string, isHTML bool, text string) {

	var e = email.NewEmail()

	e.From = fmt.Sprintf("%s <%s>", "", ef.Username)

	e.To = []string{to}

	e.Subject = subject

	if isHTML {
		e.HTML = []byte(text)
	} else {
		e.Text = []byte(text)
	}

	if err := e.Send(ef.Addr, smtp.PlainAuth("", ef.Username, ef.Password, ef.Host)); err != nil {
		helper.ErrorCommonF(common.AutoFail(common.SendEmailFailCode))
	}

}
