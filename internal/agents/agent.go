package agent

import (
	"github.com/bhavyavarshney-123/catalyst/internal/repository"
	"github.com/bhavyavarshney-123/catalyst/internal/services/rag"
	"github.com/bhavyavarshney-123/catalyst/internal/services/sync"
)

type Agent struct {
    ragService     *rag.RAGService
    syncService    *sync.SyncService
    browserService *browser.Service
    repo           *repository.OpportunityRepository
}
func NewAgent(ragService *rag.RAGService, syncService *sync.SyncService,browserService *browser.Service,Repo *repository.OpportunityRepository) *Agent {
	return &Agent{ragService: ragService, syncService: syncService,browserService: browserService, Repo: Repo}
}

//nodes of the Agent

func (a *Agent) Route(ctx context.Context, state State) (State, error) {
    ...
}

func (a *Agent) SyncGmail(ctx context.Context, state State) (State, error) {
    ...
}


func (a *Agent) RetrieveContext(ctx context.Context, state State) (State, error) {
    ...
}


func (a *Agent) GenerateResponse(ctx context.Context, state State) (State, error){
	
}
