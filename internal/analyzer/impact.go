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
func ImpactAnalysis(g *graph.Graph, startID string) []string {
	visited := make(map[string]bool)
	queue := []string{startID}

	for len(queue) > 0 {
		id := queue[0]
		queue = queue[1:]

		for _, caller := range g.AdjIn[id] {
			if !visited[caller] {
				visited[caller] = true
				queue = append(queue, caller)
			}
		}
	}

	result := make([]string, 0, len(visited))
	for id := range visited {
		result = append(result, id)
	}
	sort.Strings(result)
	return result
}
