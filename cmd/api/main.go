package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bhavyavarshney-123/catalyst/cmd/cli"
	"github.com/bhavyavarshney-123/catalyst/internal/database"
	"github.com/bhavyavarshney-123/catalyst/internal/extractor"
	"github.com/bhavyavarshney-123/catalyst/internal/handlers"
	"github.com/bhavyavarshney-123/catalyst/internal/repository"
	"github.com/bhavyavarshney-123/catalyst/internal/services/ai"
	embeddings "github.com/bhavyavarshney-123/catalyst/internal/services/embeddings"
	"github.com/bhavyavarshney-123/catalyst/internal/services/gmail"
	"github.com/bhavyavarshney-123/catalyst/internal/services/rag"
	"github.com/bhavyavarshney-123/catalyst/internal/services/sync"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {

	r := chi.NewRouter()

	err := godotenv.Load()

	if err != nil {

		fmt.Println("Error Loading the .env")
	}

	db, err := database.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	repo := repository.NewOpportunityRepository(db)
	config, err := gmail.LoadCredentials()
	if err != nil {
		panic(err)

	}

	r.Post("/opportunities", handlers.PostOpportunity(repo))
	r.Get("/opportunities", handlers.GetOpportunity(repo))
	r.Get("/opportunities/{id}", handlers.GetOpportunitybyID(repo))
	r.Put("/opportunities/{id}", handlers.UpdateOpportunity(repo))
	r.Delete("/opportunities/{id}", handlers.DeleteOpportunity(repo))

	manager := &gmail.GmailManager{}

	r.Get("/gmail/connect", handlers.ConnectGmail(config))
	r.Get("/oauth/callback", handlers.OAuthCallback(config, manager))
	r.Get("/gmail/messages", handlers.ListRecentEmails(manager))
	r.Get("/gmail/messages/Search", handlers.SearchEmails(manager))
	r.Get("/gmail/UnreadEmails", handlers.GetUnreadEmails(manager))

	OpenAI_Key := os.Getenv("OPENAI_API_KEY")
	aiService := ai.NewOpenAIService(OpenAI_Key)
	embeddingservice := embeddings.NewOpenAIEmbeddingService(OpenAI_Key)

	extractor := extractor.NewExtractor(aiService)
	syncService := sync.NewSyncService(manager, extractor, repo, embeddingservice)

	r.Get("/sync", handlers.Sync(syncService))

	go func() {
		fmt.Println("Server started on :8080")
		if err := http.ListenAndServe(":8080", r); err != nil {
			fmt.Println(err)
		}
	}()

	RAGService := rag.NewRaGService(aiService, repo, embeddingservice)
	cli.CLI(RAGService)
	r.Post("/query", handlers.RAG(RAGService))
}
