package agents

//building a graph
func BuildGraph(agent *Agent) *Graph {
	graph := NewGraph()

	graph.AddNode(NewRouteNode(agent))
	graph.AddNode(NewSyncNode(agent))
	graph.AddNode(NewResponseNode(agent))

	graph.AddEdge(Edge{
		From: "route",
		To:   "sync",
	})

	graph.AddEdge(Edge{
		From: "sync",
		To:   "response",
	})

	return graph
}
