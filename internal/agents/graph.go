package agents

import (
	"context"
	"fmt"
)

type Graph struct {
	edges []Edge
	nodes map[string]Node
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[string]Node),
		edges: make([]Edge, 0),
	}
}

func (g *Graph) AddNode(node Node) error {
	if _, exists := g.nodes[node.Name()]; exists {
		return fmt.Errorf("node %q already exists", node.Name())
	}

	g.nodes[node.Name()] = node
	return nil
}

func (g *Graph) AddEdge(edge Edge) error {
	if _, ok := g.nodes[edge.From]; !ok {
		return fmt.Errorf("node %q does not exist", edge.From)
	}

	if _, ok := g.nodes[edge.To]; !ok {
		return fmt.Errorf("node %q does not exist", edge.To)
	}
	for _, e := range g.edges {
		if e.To == edge.To && e.From == edge.From {
			return fmt.Errorf("edge from %q to %q already exists", edge.From, edge.To)
		}
	}
	g.edges = append(g.edges, edge)
	return nil
}

func (g *Graph) Execute(ctx context.Context, state *State) error {

	//defining the start point
	current := "route"

	for {
		found := false
		//retriving the node
		node, ok := g.nodes[current]
		if !ok {
			return fmt.Errorf("node %q not found", current)
		}

		//Executing the node
		if err := node.Execute(ctx, state); err != nil {
			return err
		}

		//Finding the next node
		for _, edge := range g.edges {
			if edge.From == current {
				current = edge.To
				found = true
				break
			}
		}

		if !found {
			break
		}

	}
	return nil
}
