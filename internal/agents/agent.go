package agents

import (
	"fmt"

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

func (a *Agent) Route(ctx context.Context, state *State) error {
    ...
}

func (a *Agent) SyncGmail(ctx context.Context, state *State) error{
if err:=a.syncService.Sync(state.limit);err!=nil{
	return fmt.Errorf(err)
}
}


func (a *Agent) RetrieveContext(ctx context.Context, state *State) error{
answer,err:=a.ragService.Answer(state.UserQuestion)
if err!=nil{
	return fmt.Errorf(err)
}
state.Response=answer
}


func (a *Agent) GenerateResponse(ctx context.Context, state *State) error{
	
}
