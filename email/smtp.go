package Email

import (
	"net/smtp"
	"strings"
)

var SMTP *Smtp

type Smtp struct {
	Host      string
	SendUser  string
	SendPwd   string
	ReplyUser string
	UserName  string
}

func NewSMTP() *Smtp {
	return SMTP
}

func InitSMTP(Host, User, Pwd, Reply string) {
	SMTP = &Smtp{
		Host:      Host,
		SendUser:  User,
		SendPwd:   Pwd,
		ReplyUser: Reply,
	}
}

type SMTPSendData interface {
	GetTo() string
	GetSubject() string
	GetBody() string
	GetMailType() string
}

func (E *Smtp) Send(Data SMTPSendData) error {
	hp := strings.Split(E.Host, ":")
	auth := smtp.PlainAuth("", E.SendUser, E.SendPwd, hp[0])
	Type := Data.GetMailType()
	if Type != "html" {
		Type = "plain"
	}
	msg := []byte("To: " + Data.GetTo() + "\r\nFrom: " + E.ReplyUser + ">\r\nSubject: " + Data.GetSubject() + "\r\n" + "Content-Type: text/" + Type + "; charset=UTF-8" + "\r\n\r\n" + Data.GetBody())
	sendTo := strings.Split(Data.GetTo(), ";")
	err := smtp.SendMail(E.Host, auth, E.SendUser, sendTo, msg)
	return err
}
