package agents

import (
	"context"
	"fmt"
)

type ResponseNode struct {
	agent *Agent
}

func NewResponseNode(agent *Agent) *ResponseNode {
	return &ResponseNode{
		agent: agent,
	}
}

func (n *ResponseNode) Name() string {
	return "response"
}

func (n *ResponseNode) Execute(ctx context.Context, state *State) error {
	fmt.Println("Executing Response Node")
	return n.agent.GenerateResponse(ctx, state)
}
