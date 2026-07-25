package agents

import (
	"context"

	"github.com/bhavyavarshney-123/catalyst/internal/repository"
	"github.com/bhavyavarshney-123/catalyst/internal/services/rag"
	"github.com/bhavyavarshney-123/catalyst/internal/services/sync"
)

type Agent struct {
	ragService  *rag.RAGService
	syncService *sync.SyncService
	//browserService *browser.Service
	repo *repository.OpportunityRepository
}

func NewAgent(ragService *rag.RAGService, syncService *sync.SyncService, Repo *repository.OpportunityRepository) *Agent {
	return &Agent{ragService: ragService, syncService: syncService, repo: Repo}
}

//nodes of the Agent

func (a *Agent) Route(ctx context.Context, state *State) error {
	state.NeedsSync = true
	return nil
}

func (a *Agent) SyncGmail(ctx context.Context, state *State) error {
	if err := a.syncService.Sync(state.Limit); err != nil {
		return err
	}

	return nil
}

func (a *Agent) GenerateResponse(ctx context.Context, state *State) error {
	answer, err := a.ragService.Answer(state.UserQuestion)
	if err != nil {
		return err
	}
	state.Response = answer
	return nil
}
