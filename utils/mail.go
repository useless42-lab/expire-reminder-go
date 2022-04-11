package utils

import (
	"fmt"
	"net/smtp"
	"net/textproto"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/jordan-wright/email"
)

func SendMail(toUser string, goodsName string, message string) {
	e := &email.Email{
		To:      []string{toUser},
		From:    os.Getenv("MAIL_NAME") + "<" + os.Getenv("MAIL_USERNAME") + ">",
		Subject: goodsName,
		// Text:    []byte("Text Body is, of course, supported!"),
		HTML:    []byte(message),
		Headers: textproto.MIMEHeader{},
	}
	err := e.Send(os.Getenv("MAIL_HOST")+":"+os.Getenv("MAIL_PORT"), smtp.PlainAuth("", os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"), os.Getenv("MAIL_HOST")))
	if err != nil {
		fmt.Println(err)
	}
}

func SendResetPasswordMail(targetEmail string, token string) {
	e := &email.Email{
		To:      []string{targetEmail},
		From:    os.Getenv("MAIL_NAME") + "<" + os.Getenv("MAIL_USERNAME") + ">",
		Subject: "监控提醒",
		// Text:    []byte("Text Body is, of course, supported!"),
		HTML:    []byte("https://www.pingsilo.com/auth/reset?t=" + token + "<br/>" + "五分钟内有效"),
		Headers: textproto.MIMEHeader{},
	}
	err := e.Send(os.Getenv("MAIL_HOST")+":"+os.Getenv("MAIL_PORT"), smtp.PlainAuth("", os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"), os.Getenv("MAIL_HOST")))
	if err != nil {
		fmt.Println(err)
	}
}
