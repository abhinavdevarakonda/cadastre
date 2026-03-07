package analyzer

import (
	"sort"

	"github.com/abhinavdevarakonda/maplet/internal/graph"
)

// ImpactAnalysis returns the set of function IDs that are transitively
// affected if the given startID changes. It walks the reverse call graph
// (AdjIn) via BFS: if A calls B and B calls C, then changing C impacts B and A.
//
// The start node itself is NOT included in the result — only its callers are.
// Results are returned sorted alphabetically.
type ImpactResult struct {
	ID   string
	Line int
}

func ImpactAnalysis(g *graph.Graph, startID string) []ImpactResult {
	var results []ImpactResult
	for _, e := range g.Edges {
		if e.Type == graph.CallsEdge && e.To == startID {
			results = append(results, ImpactResult{
				ID:   e.From,
				Line: e.Line,
			})
		}
	}
	sort.Slice(results, func(i, j int) bool {
		if results[i].ID != results[j].ID {
			return results[i].ID < results[j].ID
		}
		return results[i].Line < results[j].Line
	})
	return results
}
