package gmail

import (
	"context"
	"encoding/base64"
	"fmt"

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

func (g *GmailService) FetchEmail(id string) (Email, error) {

	userid := "me"
	call := g.service.Users.Messages.Get(userid, id)

	message, err := call.Do()
	if err != nil {
		return Email{}, err
	}

	payload := message.Payload
	if payload == nil {
		fmt.Errorf("Payload is nil")
		return Email{}, nil
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
			return Email{}, err
		}

		body = string(decoded)
	} else {

		for _, part := range payload.Parts {
			        switch part.MimeType {
      case "text/plain", "text/html":
    decoded, err := base64.RawURLEncoding.DecodeString(part.Body.Data)
    if err != nil {
        return Email{}, err
    }

    if part.MimeType == "text/plain" {
        body = string(decoded)
    } else if body == "" {
        body = string(decoded)
    }
	}
	return Email{
		ID:      message.Id,
		Subject: subject,
		From:    from,
		Body:    body,
		Date:    date,
	}, nil

}
func (g *GmailService) ListRecentMessages(limit int64) ([]Email, error) {

	var Email []Email

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
