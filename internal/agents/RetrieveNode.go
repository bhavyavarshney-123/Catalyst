package agents

import "context"

type RetrieveNode struct {
	agent *Agent
}

func NewRetrieveNode(agent *Agent) *RetrieveNode {
	return &RetrieveNode{
		agent: agent,
	}
}

func (n *RetrieveNode) Name() string {
	return "GenerateResponse"
}

func (n *RetrieveNode) Execute(ctx context.Context, state *State) error {
	return n.agent.RetrieveContext(ctx, state)
}
