package main

import (
	"fmt"
)

type MessSender interface {
	Send(message string)
}

type EmailSender struct{}

func (s *EmailSender) Send(message string) {
	fmt.Printf("Email: %s \n", message)
}

type SmsSender struct{}

func (s *SmsSender) Send(message string) {
	fmt.Printf("SMS: %s \n", message)
}

type TGSender struct{}

func (s *TGSender) Send(message string) {
	fmt.Printf("Telegram: %s \n", message)
}

type NotificationServ struct {
	sender MessSender
}

func (n *NotificationServ) changeSender(newSender MessSender) {
	n.sender = newSender
}

func (n *NotificationServ) SendNotification(msg string) {
	n.sender.Send(msg)
}

func main() {
	email := &EmailSender{}
	sms := &SmsSender{}
	tg := &TGSender{}

	serv := &NotificationServ{sender: email}
	serv.SendNotification("hellow to email")

	serv.changeSender(sms)
	serv.SendNotification("hellow to sms")

	serv.changeSender(tg)
	serv.SendNotification("hellow to tg")
}
