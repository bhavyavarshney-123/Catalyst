package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bhavyavarshney-123/catalyst/internal/services/gmail"
	"golang.org/x/oauth2"
)

func ConnectGmail(config *oauth2.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		AuthURL := gmail.GenerateAuthURL(config)
		http.Redirect(w, r, AuthURL, http.StatusTemporaryRedirect)
	}
}

func OAuthCallback(config *oauth2.Config, manager *gmail.GmailManager) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		code := r.URL.Query().Get("code")

		gmailService, err := gmail.NewGmailService(config, code)
		if err != nil {
			http.Error(w, fmt.Sprintf("Gmail Service error: %v", err), http.StatusBadRequest)
			return
		}

		fmt.Fprintln(w, "Gmail authenticated successfully!")

		manager.Service = gmailService
	}
}

func ListRecentEmails(manager *gmail.GmailManager) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query().Get("limit")
		limit, err := strconv.Atoi(param)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		message, err := manager.Service.ListRecentMessages(int64(limit))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Printf("%+v\n", message)
	}
}

func SearchEmails(manager *gmail.GmailManager) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		query := r.URL.Query().Get("q")

		message, err := manager.Service.SearchEmails(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Printf("%+v\n", message)
	}
}

func GetUnreadEmails(manager *gmail.GmailManager) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query().Get("limit")
		limit, err := strconv.Atoi(param)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		message, err := manager.Service.GetUnreadEmails(int64(limit))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Printf("%+v\n", message)
	}
}


func SyncEmails(...) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        emails := gmailService.ListRecentMessages(20)

        for _, email := range emails {

            // Ignore non-job emails

            // Convert Email -> Opportunity

            // Save Opportunity
        }

    }
}
