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

func LoadCredentials() (*oauth2.Config, error) {
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

func GenerateAuthURL(config *oauth2.Config) string {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Println(authURL)
	return authURL

}

func ExchangeCode(config *oauth2.Config, code string) (*http.Client, error) {

	ctx := context.Background()

	token, err := config.Exchange(ctx, code)
	if err != nil {
		return nil, err

	}

	httpClient := config.Client(ctx, token)

	return httpClient, nil
}
