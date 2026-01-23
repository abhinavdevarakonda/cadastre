package export

import (
	"github.com/abhinavdevarakonda/maplet/internal/callgraph"
	"github.com/abhinavdevarakonda/maplet/internal/graph"
)

type Graph struct {
	Nodes     []Node     `json:"nodes"`
	Edges     []Edge     `json:"edges"`
	CallEdges []CallEdge `json:"callEdges"`
}

type Node struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
	Path string `json:"path"`
}

type Edge struct {
	From string `json:"from"`
	To   string `json:"to"`
	Type string `json:"type"`
}

type CallEdge struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func FromGraph(g *graph.Graph) Graph {
	out := Graph{
		Nodes: make([]Node, 0, len(g.Nodes)),
		Edges: make([]Edge, 0, len(g.Edges)),
	}

	for _, n := range g.Nodes {
		out.Nodes = append(out.Nodes, Node{
			ID:   n.ID,
			Type: string(n.Type),
			Name: n.Name,
			Path: n.Path,
		})
	}

	for _, e := range g.Edges {
		out.Edges = append(out.Edges, Edge{
			From: e.From,
			To:   e.To,
			Type: string(e.Type),
		})
	}

	return out
}

func FromCallGraph(cg *callgraph.Graph) []CallEdge {
	out := []CallEdge{}

	for from, tos := range cg.Edges {
		for _, to := range tos {
			out = append(out, CallEdge{
				From: from,
				To:   to,
			})
		}
	}

	return out
}
