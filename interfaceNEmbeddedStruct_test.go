package gobasic

import (
	"fmt"
	"testing"
)

type Notifier interface {
	SendMessage(message string)
}

type BaseNotifier struct {
	AppName string
}

type EmailNotifier struct {
	Email string
	BaseNotifier
}

func (en *EmailNotifier) SendMessage(message string) {
	fmt.Printf(`
	APP %s
	to: %s
	Message: %s
	`, en.AppName, en.Email, message)
}

type SMSNotifier struct {
	PhoneNumber string
	BaseNotifier
}

func (sn *SMSNotifier) SendMessage(message string) {
	fmt.Printf(`
	APP %s
	Phone number: %s
	Message: %s
	`, sn.AppName, sn.PhoneNumber, message)
}

func TestSendNotification(t *testing.T) {
	e := EmailNotifier{
		Email:   "lim@example.com",
		AppName: "M-MAIL",
	}
	s := SMSNotifier{
		PhoneNumber: "099999428424",
		AppName:     "WA",
	}
	//Email Notifier
	SendNotification(&e, "Hello World")

	//SMS Notifier
	SendNotification(&s, "Hello World")
}

func SendNotification(notifier Notifier, message string) {
	notifier.SendMessage(message)
}
