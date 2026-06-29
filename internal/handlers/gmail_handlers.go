package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bhavyavarshney-123/catalyst/internal/services/gmail"
	"golang.org/x/oauth2"
)

func ConnectGmail(config *oauth2.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		AuthURL := gmail.GenerateAuthURL(config)
		http.Redirect(w, r, AuthURL, http.StatusTemporaryRedirect)
	}
}

func OAuthCallback(config *oauth2.Config) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		code := r.URL.Query().Get("code")

		gmailService, err := gmail.NewGmailService(config, code)
		if err != nil {
			http.Error(w, fmt.Sprintf("Gmail Service error: %v", err), http.StatusBadRequest)
		}

		message, err := gmailService.ListRecentMessages(5)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v\n", message)
	}
}
