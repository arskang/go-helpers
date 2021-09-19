package goHelpers

import "gopkg.in/gomail.v2"

type RequestSendMails struct {
	ErrorEmail   *[]ErrorMail `json:"errorEmail,omitempty"`
	SuccessEmail *[]string    `json:"successEmail,omitempty"`
	Data         interface{}  `json:"data,omitempty"`
}

type ErrorMail struct {
	Email string `json:"email,omitempty"`
	Error string `json:"error,omitempty"`
}

type Email struct {
	From    string        `json:"from,omitempty"`
	To      []string      `json:"to,omitempty"`
	CC      *[]CCMail     `json:"cc,omitempty"`
	Subject string        `json:"subject,omitempty"`
	Body    BodyMail      `json:"body,omitempty"`
	Attach  *[]AttachMail `json:"attach,omitempty"`
	Config  *EmailConfig  `json:"config,omitempty"`
}

type CCMail struct {
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

type BodyMail struct {
	Type string `json:"type,omitempty"`
	Info string `json:"info,omitempty"`
}

type AttachMail struct {
	Filename string             `json:"filename,omitempty"`
	Settings gomail.FileSetting `json:"settings,omitempty"`
}

type EmailConfig struct {
	SMTP string `json:"smtp,omitempty"`
	PORT int    `json:"port,omitempty"`
	USER string `json:"user,omitempty"`
	PASS string `json:"password,omitempty"`
}

type EmailTypes struct {
	HTML *string
}

func (email *Email) SendMail() error {

	m := gomail.NewMessage()
	m.SetHeader("From", email.From)
	m.SetHeader("To", email.To...)

	if email.CC != nil {
		for _, cc := range *email.CC {
			m.SetAddressHeader("Cc", cc.Email, cc.Name)
		}
	}

	m.SetHeader("Subject", email.Subject)
	m.SetBody(email.Body.Type, email.Body.Info)

	if email.Attach != nil {
		for _, attach := range *email.Attach {
			m.Attach(attach.Filename, attach.Settings)
		}
	}

	var d *gomail.Dialer
	if email.Config != nil {
		d = gomail.NewDialer(
			email.Config.SMTP,
			email.Config.PORT,
			email.Config.USER,
			email.Config.PASS,
		)
	} else {
		smtp, err := GetENV("EMAIL_SMTP")
		if err != nil {
			return err
		}
		user, err := GetENV("EMAIL_USER")
		if err != nil {
			return err
		}
		pass, err := GetENV("EMAIL_PASS")
		if err != nil {
			return err
		}
		d = gomail.NewDialer(
			*smtp,
			587,
			*user,
			*pass,
		)
	}

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil

}
