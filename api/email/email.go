package email

import (
	"auth-service/config"
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"

	rdb "auth-service/storage/redis"

	"github.com/pkg/errors"
)

func SendEmail(email string) (string, error) {
	source := rand.NewSource(time.Now().UnixNano())
	myRand := rand.New(source)

	randomNumber := myRand.Intn(900000) + 100000

	code := strconv.Itoa(randomNumber)

	err := rdb.StoreCode(context.Background(), email, code)
	if err != nil {
		return "", err
	}

	_, err = rdb.GetCode(context.Background(), email)
	if err != nil {
		return "", err
	}

	err = sendCode(email, code)
	if err != nil {
		return "", errors.Wrap(err, "failed to send code")
	}

	return code, nil

}

func sendCode(email string, code string) error {
	from := config.Load().Server.EMAIL
	password := config.Load().Server.APP_KEY

	to := []string{
		email,
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, err := template.ParseFiles("api/email/template.html")
	if err != nil {
		return errors.Wrap(err, "failed to parse template")
	}

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Your verification code \n%s\n\n", mimeHeaders)))
	t.Execute(&body, struct {
		Passwd string
	}{

		Passwd: code,
	})

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		return errors.Wrap(err, "failed to send email")
	}

	log.Println("Email sended to:", email)
	return nil
}
