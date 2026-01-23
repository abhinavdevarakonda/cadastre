package graph

type NodeType string

const (
	DirectoryNode NodeType = "directory"
	FileNode      NodeType = "file"
	FunctionNode  NodeType = "function"
)

type Node struct {
	ID   string
	Type NodeType
	Name string
	Path string
}

type EdgeType string

const ContainsEdge EdgeType = "contains"

type Edge struct {
	From string
	To   string
	Type EdgeType
}

type Graph struct {
	Nodes map[string]*Node
	Edges []*Edge
}

func New() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
		Edges: []*Edge{},
	}
}

func (g *Graph) AddNode(n *Node) {
	g.Nodes[n.ID] = n
}

func (g *Graph) AddEdge(from, to string, t EdgeType) {
	g.Edges = append(g.Edges, &Edge{
		From: from,
		To:   to,
		Type: t,
	})
}
