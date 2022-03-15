package pkg

import (
	"errors"
	"net/mail"
	"net/smtp"
	"strconv"
)

type smtpClient struct {
	user     string
	password string
	host     string
	port     int
	identity string
	Email
}

type Email struct {
	From    string
	To      []string
	Cc      []string
	Bcc     []string
	Subject string
	Body    string
}

func NewSMTPClient(user, password, host string,port int) *smtpClient {
	return &smtpClient{
		user:     user,
		password: password,
		host:     host,
		port:     port,
	}
}

func (s *smtpClient) SetIdentity(identity string) {
	s.identity = identity
}

func (s *smtpClient) Send() error {
	contentType := "Content-Type: text/plain; charset=UTF-8"
	if err := s.send(contentType); err != nil {
		return err
	}
	return nil
}

func (s *smtpClient) SendHTML() error {
	contentType := "Content-Type: text/html; charset=UTF-8"
	if err := s.send(contentType); err != nil {
		return err
	}
	return nil
}

func (s *smtpClient) TestSend() error {
	return nil
}

func (s *smtpClient) send(contentType string) error {
	if s.From == "" {
		s.From = s.user
	}
	from, err := mail.ParseAddress(s.user)
	if err != nil {
		return err
	}

	to := make([]string, 0, len(s.To)+len(s.Cc)+len(s.Bcc))
	to = append(append(append(to, s.To...), s.Cc...), s.Bcc...)
	for i := 0; i < len(to); i++ {
		addr, err := mail.ParseAddress(to[i])
		if err != nil {
			return err
		}
		to[i] = addr.Address
	}
	if len(to) == 0 {
		return errors.New("至少要有一个接收人")
	}

	raw := s.Bytes(contentType)

	addr := s.host + ":" + strconv.Itoa(s.port)
	auth := smtp.PlainAuth(s.identity, s.user, s.password, s.host)
	return smtp.SendMail(addr, auth, from.Address, to, raw)
}

func (s *smtpClient) Bytes(contentType string) []byte {
	msg := "To: "
	to := make([]string, 0, len(s.To)+len(s.Cc)+len(s.Bcc))
	to = append(append(append(to, s.To...), s.Cc...), s.Bcc...)
	for i := 0; i < len(to); i++ {
		msg += to[i] + ","
	}
	msg += "\r\nFrom: " + s.From + "<" + s.user + ">" + "\r\nSubject: " + s.Subject + "\r\n" + contentType + "\r\n\r\n" + s.Body
	return []byte(msg)
}
