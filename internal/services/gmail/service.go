package gmail

import (
	"context"
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

func (g *GmailService) FetchEmail(id string) (*Email, error) {

	userid := "me"
	call := g.service.Users.Messages.Get(userid, id)

	message, err := call.Do()
	if err != nil {
		return nil, err
	}

	payload := message.Payload
	if payload == nil {
		fmt.Errorf("Payload is nil")
		return nil, nil
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

	body := payload.Body

	fmt.Println(body)
	return &Email{
		ID:      message.Id,
		Subject: subject,
		From:    from,
		Date:    date,
	}, nil

}

func (g *GmailService) ListRecentMessages() (*gmail.Profile, error) {

	resp, err := g.service.Users.GetProfile("me").Do()

	return resp, err
}
