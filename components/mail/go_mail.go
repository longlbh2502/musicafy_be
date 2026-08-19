package mail

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"

	"example.com/musicafy_be/common"
	"github.com/wneessen/go-mail"
)

type Mailer struct {
	sender   string
	password string
	address  string
	m        *mail.Msg
}

func NewMailer(sender string, password string, address string) Mailer {
	m := mail.NewMsg()
	return Mailer{sender: sender, password: password, address: address, m: m}
}

func (ml *Mailer) SendMail(to string, subject string, body string) error {
	m := ml.m
	if err := m.From(ml.address); err != nil {
		return err
	}

	if err := m.To(to); err != nil {
		return err
	}

	m.Subject(subject)
	m.SetBodyString(mail.TypeTextHTML, body)

	// Secondly the mail client
	c, err := mail.NewClient("smtp.gmail.com",
		mail.WithPort(587), mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(ml.address), mail.WithPassword(ml.password),
	)
	if err != nil {
		// log.Fatal().Msgf("failed to create mail client: %s", err)
		return err
	}

	// Finally let's send out the mail
	err = c.DialAndSend(m)
	if err != nil {
		// log.Fatal().Msgf("failed to send mail: %s", err)
		return err
	}
	return nil
}

func (ml *Mailer) SendOTPMail(to string, code string) error {
	path := filepath.Join("components", "mail", "msg.html")
	htmlContent, err := os.ReadFile(path)
	if err != nil {
		return common.NewCustomError(err, "send otp mail", "Lỗi đọc file html", "MAIL")
	}

	tmpl, err := template.New("email").Parse(string(htmlContent))
	if err != nil {
		return common.NewCustomError(err, "send otp mail", "Lỗi parse template", "MAIL")
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, map[string]interface{}{
		"OTP": code,
	}); err != nil {
		return common.NewCustomError(err, "send otp mail", "Lỗi execute template", "MAIL")
	}

	err = ml.SendMail(to, "Xác thực tài khoản", body.String())
	if err != nil {
		return common.NewCustomError(err, "send otp mail", "Lỗi gửi email", "MAIL")
	}

	return nil

}
