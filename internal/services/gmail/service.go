package gmail

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/bhavyavarshney-123/catalyst/internal/models"
	"golang.org/x/oauth2"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type GmailService struct {
	service *gmail.Service
}

func NewGmailService(config *oauth2.Config, code string) (*GmailService, error) {

	httpClient, err := ExchangeCode(config, code)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		return nil, err

	}

	return &GmailService{
		service: srv,
	}, nil

}

func (g *GmailService) FetchEmail(id string) (models.Email, error) {

	userid := "me"
	call := g.service.Users.Messages.Get(userid, id)

	message, err := call.Do()
	if err != nil {
		return models.Email{}, err
	}

	payload := message.Payload
	if payload == nil {
		fmt.Errorf("Payload is nil")
		return models.Email{}, nil
	}

	var subject, from, date string

	for _, h := range payload.Headers {
		switch h.Name {
		case "Subject":
			subject = h.Value
		case "From":
			from = h.Value
		case "Date":
			date = h.Value
		}
	}

	var body string

	if payload.Body.Data != "" {
		decoded, err := base64.URLEncoding.DecodeString(payload.Body.Data)
		if err != nil {
			return models.Email{}, err
		}

		body = string(decoded)
	} else {

		for _, part := range payload.Parts {

			if part.MimeType == "text/plain" {
				decoded, err := base64.URLEncoding.DecodeString(part.Body.Data)
				if err != nil {
					return models.Email{}, err
				}

				body = string(decoded)
				break
			}
		}
	}
	return models.Email{
		ID:      message.Id,
		Subject: subject,
		From:    from,
		Body:    body,
		Date:    date,
	}, nil

}
func (g *GmailService) ListRecentMessages(limit int64) ([]models.Email, error) {

	var Email []models.Email

	resp, err := g.service.Users.Messages.List("me").MaxResults(limit).Do()
	if err != nil {
		return nil, err
	}

	for _, msg := range resp.Messages {
		email, err := g.FetchEmail(msg.Id)
		if err != nil {
			return nil, err
		}
		Email = append(Email, email)
	}

	return Email, nil
}

func (g *GmailService) SearchEmails(query string) ([]models.Email, error) {

	var Email []models.Email

	resp, err := g.service.Users.Messages.List("me").Q(query).Do()
	if err != nil {
		return nil, err
	}

	for _, msg := range resp.Messages {
		email, err := g.FetchEmail(msg.Id)
		if err != nil {
			return nil, err
		}
		Email = append(Email, email)
	}

	return Email, nil

}

func (g *GmailService) GetUnreadEmails(limit int64) ([]models.Email, error) {

	var Email []models.Email

	resp, err := g.service.Users.Messages.List("me").MaxResults(limit).Q("is:unread").Do()
	if err != nil {
		return nil, err
	}

	for _, msg := range resp.Messages {
		email, err := g.FetchEmail(msg.Id)
		if err != nil {
			return nil, err
		}
		Email = append(Email, email)
	}

	return Email, nil

}
