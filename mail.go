package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
)

var to string = os.Args[1]
var subject string = os.Args[2]
var body string = os.Args[3]

var smtpServer string
var From string

var username string
var password string

var logFile string = "mail.log"

func initLog() {
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	return
}

func init() {
	initLog()
}

func main() {
	a := smtp.PlainAuth("", username, password, smtpServer)
	toSlice := strings.Split(to, ";")
	msg := []byte(body)
	fmt.Println(body)
	if err := smtp.SendMail(smtpServer, a, From, toSlice, msg); err != nil {
		log.Fatalf("SendMail:%v\n", err)
	}
}
