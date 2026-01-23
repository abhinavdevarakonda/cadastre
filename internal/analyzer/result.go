package analyzer

import (
	"github.com/abhinavdevarakonda/maplet/internal/callgraph"
	"github.com/abhinavdevarakonda/maplet/internal/graph"
)

type Result struct {
	Structure *graph.Graph
	Call      *callgraph.Graph
}

