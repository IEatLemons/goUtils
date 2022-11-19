package Email_test

import (
	"fmt"
	"testing"

	Email "github.com/IEatLemons/goUtils/email"
	EmailModules "github.com/IEatLemons/goUtils/email/modules"
)

func Init() {
	Host := "smtp.163.com:25"
	User := "xxx@163.com"
	Pwd := "xxxxx"
	Reply := "xxx@163.com"

	Email.InitSMTP(Host, User, Pwd, Reply)
}

func TestSmtp_Send(t *testing.T) {
	Init()
	Data := &EmailModules.Test{
		User:    "xxx@163.com",
		Subject: "Test Sender",
		Name:    "Test Sender",
		Number:  "0",
	}
	err := Email.SMTP.Send(Data)
	fmt.Println("err", err)
}
