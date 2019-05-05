package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

func main() {
	auth := smtp.PlainAuth("", "958752538@qq.com", "xxsmlcsglyhdbege", "smtp.qq.com")
	to := []string{"958752538@qq.com"}
	nickname := "test"
	user := "958752538@qq.com"
	subject := "学长智障"
	content_type := "Content-Type: text/" + "html" + "; charset=UTF-8"
	body := "This is the email body.<img src = 'http://p1.music.126.net/QBQvSe84znaWIm2PrhQ_ng==/5930765720512200.jpg?param=130y130'/>"
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	for i := 0; i < 2; i++ {
		err := smtp.SendMail("smtp.qq.com:25", auth, user, to, msg)
		if err != nil {
			fmt.Printf("send mail error: %v", err)
		}
	}
}
