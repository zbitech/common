package mailer

import (
	"bytes"
	"context"
	"embed"
	"text/template"
	"time"

	"github.com/zbitech/common/pkg/model/object"
	"gopkg.in/mail.v2"
)

var templateFS embed.FS

type Mailer struct {
	dailer    *mail.Dialer
	sender    string
	mail_tmpl []object.FileTemplate
}

func New(host string, port int, username, password, sender string, template_files []string) (*Mailer, error) {
	dialer := mail.NewDialer(host, port, username, password)
	dialer.Timeout = 5 * time.Second
	mail_tmpl := make([]object.FileTemplate, len(template_files))

	for index, template_file := range template_files {
		tmpl, err := object.NewFileTemplate(template_file)
		if err != nil {
			return nil, err
		}

		mail_tmpl[index] = *tmpl
	}

	return &Mailer{dailer: dialer, sender: sender, mail_tmpl: mail_tmpl}, nil
}

func (m Mailer) Send(ctx context.Context, recipient, templateFile string, data interface{}) error {
	tmpl, err := template.New("email").ParseFS(templateFS, "templates/"+templateFile)
	if err != nil {
		return err
	}

	subject := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(subject, "subject", data); err != nil {
		return err
	}

	plainBody := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(plainBody, "plainBody", data); err != nil {
		return err
	}

	htmlBody := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(htmlBody, "htmlBody", data); err != nil {
		return err
	}

	msg := mail.NewMessage()
	msg.SetHeader("To", recipient)
	msg.SetHeader("From", m.sender)
	msg.SetHeader("Subject", subject.String())

	msg.SetBody("text/plain", plainBody.String())
	msg.AddAlternative("text/html", htmlBody.String())

	for i := 1; i <= 3; i++ {
		err = m.dailer.DialAndSend(msg)
		if err == nil {
			return nil
		}

		time.Sleep(time.Millisecond * 500)
	}

	return err
}
