package callgraph

type Graph struct {
	// Edges maps a caller ID to a list of callee IDs
	Edges map[string][]string
}

func New() *Graph {
	return &Graph{
		Edges: make(map[string][]string),
	}
}

func (g *Graph) AddEdge(from, to string) {
	g.Edges[from] = append(g.Edges[from], to)
}
