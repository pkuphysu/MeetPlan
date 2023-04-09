package mail

import (
	gomail "github.com/wneessen/go-mail"
	"html/template"
)

type Mail struct {
	Subject          string
	To               []string
	BodyTemplateHTML *template.HTML
}

func SendSameMail(mail *Mail) (err error) {
	m := gomail.NewMsg()
	err = m.From("gwlxgbfwq@pku.edu.cn")
	if err != nil {
		return
	}
	err = m.To(mail.To...)
	if err != nil {
		return
	}
	m.Subject(mail.Subject)
	//m.SetBodyHTMLTemplate(mail.Body)
	//m.AddAlternativeHTMLTemplate()
	return nil
}
