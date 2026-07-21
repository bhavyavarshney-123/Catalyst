package agent

import (
	"github.com/bhavyavarshney-123/catalyst/internal/repository"
	"github.com/bhavyavarshney-123/catalyst/internal/services/rag"
	"github.com/bhavyavarshney-123/catalyst/internal/services/sync"
)

type Agent struct {
	ragService  *rag.RAGService
	syncService *sync.SyncService
	Repo        *repository.OpportunityRepository
}

func NewAgent(ragService *rag.RAGService, syncService *sync.SyncService, Repo *repository.OpportunityRepository) *Agent {
	return &Agent{ragService: ragService, syncService: syncService, Repo: Repo}
}


func (a *Agent) SyncGmail(ctx context.Context, state State) (State, error) {
    ...
}

func (a *Agent) ExtractOpportunities(ctx context.Context, state State) (State, error) {
    ...
}

func (a *Agent) FindRelevantOpportunities(ctx context.Context, state State) (State, error) {
    ...
}


func (a *Agent) GenerateResponse(ctx context.Context, state State) (State, error){
	
}
