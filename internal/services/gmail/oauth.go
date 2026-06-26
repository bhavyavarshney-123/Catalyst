package gmail

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

func loadcredentials() (*oauth2.Config, error) {
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

func authenticate(config *oauth2.Config) error(
config,err:=loadcredentials()
if err!=nil{
	return err
}

authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
if err!=nil{
	return err
}

return nil
)
