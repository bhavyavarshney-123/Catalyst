package app

import (
	"database/sql"

	"github.com/bhavyavarshney-123/catalyst/internal/agents"
	"github.com/bhavyavarshney-123/catalyst/internal/config"
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
	"golang.org/x/oauth2"
)

type App struct {
	Router *chi.Mux
	DB     *sql.DB

	GmailManager *gmail.GmailManager
	GmailConfig  *oauth2.Config

	Repo *repository.OpportunityRepository

	AIService        *ai.OpenAIService
	EmbeddingService *embeddings.OpenAIEmbeddingService

	Extractor   *extractor.Extractor
	SyncService *sync.SyncService

	RAGService *rag.RAGService
	Agent      *agents.Agent
}

func (a *App) registerDependencies() error {

	// Load environment
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	// Database
	db, err := database.Connect()
	if err != nil {
		return err
	}
	a.DB = db

	// Repository
	a.Repo = repository.NewOpportunityRepository(db)

	// Gmail
	config, err := gmail.LoadCredentials()
	if err != nil {
		return err
	}
	a.GmailConfig = config

	a.GmailManager = &gmail.GmailManager{}

	// AI
	openAIKey := cfg.OpenAIKey

	a.AIService = ai.NewOpenAIService(openAIKey)
	a.EmbeddingService = embeddings.NewOpenAIEmbeddingService(openAIKey)

	// Extractor
	a.Extractor = extractor.NewExtractor(a.AIService)

	// Sync
	a.SyncService = sync.NewSyncService(
		a.GmailManager,
		a.Extractor,
		a.Repo,
		a.EmbeddingService,
	)

	// RAG
	a.RAGService = rag.NewRaGService(
		a.AIService,
		a.Repo,
		a.EmbeddingService,
	)

	// Agent
	a.Agent = agents.NewAgent(
		a.RAGService,
		a.SyncService,
		a.Repo,
	)

	return nil
}

func (a *App) registerRoutes() {

	a.Router.Post("/opportunities", handlers.PostOpportunity(a.Repo))
	a.Router.Get("/opportunities", handlers.GetOpportunity(a.Repo))
	a.Router.Get("/opportunities/{id}", handlers.GetOpportunitybyID(a.Repo))
	a.Router.Put("/opportunities/{id}", handlers.UpdateOpportunity(a.Repo))
	a.Router.Delete("/opportunities/{id}", handlers.DeleteOpportunity(a.Repo))

	a.Router.Get("/gmail/connect", handlers.ConnectGmail(a.GmailConfig))
	a.Router.Get("/oauth/callback", handlers.OAuthCallback(a.GmailConfig, a.GmailManager))
	a.Router.Get("/gmail/messages", handlers.ListRecentEmails(a.GmailManager))
	a.Router.Get("/gmail/messages/Search", handlers.SearchEmails(a.GmailManager))
	a.Router.Get("/gmail/UnreadEmails", handlers.GetUnreadEmails(a.GmailManager))

	a.Router.Get("/sync", handlers.Sync(a.SyncService))

	a.Router.Post("/query", handlers.RAG(a.RAGService))
}

func New() (*App, error) {

	app := &App{
		Router: chi.NewRouter(),
	}

	if err := app.registerDependencies(); err != nil {
		return nil, err
	}

	app.registerRoutes()

	return app, nil
}
