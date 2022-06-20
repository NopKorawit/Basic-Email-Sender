package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"s_email/config"
)

func main() {

	configModel := config.EmailConfig{}
	configModel.LoadConfig()

	// Sender data.
	from := configModel.Email        // <------------- (1) แก้ไขอีเมลที่ใช้ส่ง
	password := configModel.Password // <------- (2) แก้ไขรหัสผ่านของอีเมลที่ใช้ส่ง

	// smtp server configuration.
	smtpHost := configModel.SmtpHost //<--------------(3) Host ที่ใช้
	smtpPort := configModel.SmtpPort // <----------------- (4) port ใช้ tls คือ 465,587

	// Receiver email address.
	to := "example001@gmail.com" // <-------------- (5) แก้ไขอีเมลของผู้รับ
	servername := smtpHost + ":" + smtpPort
	// Message.
	subj := "คุณถูกล็อตเตอรี่รางวัลที่ 2 หากโอนมา 500 ในตอนนี้!"               // <----------------- (6) หัวเรื่อง
	body := "รับเงินไปเลย 3 ล้านบาท รีบโอนมาที่ 0856727284 พร้อมเพย์ กรวิชญ์." // <----------------- (7) เนื้อความ

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from
	headers["To"] = to
	headers["Subject"] = subj

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpHost,
	}

	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	c, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = c.Mail(from); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(to); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	// Sending email.
	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	c.Quit()

	fmt.Println("Email Sent Successfully!")
}
