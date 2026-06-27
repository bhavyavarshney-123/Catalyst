package gmail

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

func loadCredentials() (*oauth2.Config, error) {
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		return nil, err
	}

	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func authenticate() (*http.Client, error) {
	config, err := loadCredentials()
	if err != nil {
		return nil, err
	}

	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	fmt.Println("Please retrieve authentication code by logging into this URL:", authURL)

	var code string
	_, err = fmt.Scan(&code)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()

	token, err := config.Exchange(ctx, code)
	if err != nil {
		return nil, err

	}

	httpClient := config.Client(ctx, token)

	return httpClient, nil
}
