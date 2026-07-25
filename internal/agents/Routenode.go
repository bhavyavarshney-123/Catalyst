package agents

import (
	"context"
	"fmt"
)

type RouteNode struct {
	agent *Agent
}

func NewRouteNode(agent *Agent) *RouteNode {
	return &RouteNode{
		agent: agent,
	}
}

func (n *RouteNode) Name() string {
	return "route"
}

func (n *RouteNode) Execute(ctx context.Context, state *State) error {
	fmt.Println("Executing Route Node")
	return n.agent.Route(ctx, state)
}
