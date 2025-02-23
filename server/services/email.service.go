package services

import (
	"fmt"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailService struct {
	client *sendgrid.Client
}

func NewEmailService(apiKey string) *EmailService {
	emailClient := sendgrid.NewSendClient(apiKey)
	return &EmailService{
		client: emailClient,
	}
}

type EmailSendOptions struct {
	FromEmail   string
	ToEmail     string
	Subject     string
	HTMLContent string
}

func (s *EmailService) SendEmail(options *EmailSendOptions) bool {
	from := mail.NewEmail("", options.FromEmail)
	subject := options.Subject
	to := mail.NewEmail("", options.ToEmail)
	plainTextContent := "MyRankedList does not support plain text emails. Please enable HTML."
	htmlContent := options.HTMLContent
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	response, err := s.client.Send(message)
	if err != nil {
		log.Println(err)
		return false
	} else {
		if response.StatusCode != 202 {
			fmt.Println(response.StatusCode)
			fmt.Println(response.Body)
			fmt.Println(response.Headers)
			return false
		}
		return true
	}
}
