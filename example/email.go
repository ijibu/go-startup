package main

import (
	"net/smtp"
	"fmt"
	"strings"
)

func SendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type :text/plain" + "; charset=UTF-8"
	}
	
	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
} 

func main() {
	user := "43546936@qq.com"
	password := "?1pvm2!"
	host := "smtp.qq.com:25"
	to := "308185321@qq.com;contact@ita001.com"
	subject := "Test send email by golang"
	body := `<html><body>
			<h3>"test send email by golang"</h3>
			</body></html>`
	fmt.Println("Send email")
	err := SendMail(user, password, host, to, subject, body, "html")
	
	if err != nil {
		fmt.Println("send email error!")
		fmt.Println(err)
	} else {
		fmt.Println("send email success!")
	}
}



