package agents

import "context"

type SyncNode struct {
	agent *Agent
}

func NewSyncNode(agent *Agent) *SyncNode {
	return &SyncNode{
		agent: agent,
	}
}

func (n *SyncNode) Name() string {
	return "sync"
}

func (n *SyncNode) Execute(ctx context.Context, state *State) error {
	return n.agent.SyncGmail(ctx, state)
}
