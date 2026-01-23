package analyzer

import (
	"path/filepath"
	"strings"

	"github.com/abhinavdevarakonda/maplet/internal/callgraph"
	"github.com/abhinavdevarakonda/maplet/internal/graph"
)

func Build(scan *ScanResult, symbols []Symbol, facts []Fact) Result {
	structGraph := graph.New()
	callGraph := callgraph.New()

	// add dirs and files to structure graph
	for _, dir := range scan.Directories {
		structGraph.AddNode(&graph.Node{
			ID:   dir,
			Type: graph.DirectoryNode,
			Name: filepath.Base(dir),
			Path: dir,
		})
		if dir != scan.Root {
			structGraph.AddEdge(filepath.Dir(dir), dir, graph.ContainsEdge)
		}
	}

	for _, file := range scan.Files {
		structGraph.AddNode(&graph.Node{
			ID:   file,
			Type: graph.FileNode,
			Name: filepath.Base(file),
			Path: file,
		})
		structGraph.AddEdge(filepath.Dir(file), file, graph.ContainsEdge)
	}

	// add symbols to structure graph and create a lookup table
	symbolTable := make(SymbolTable)
	for _, sym := range symbols {
		symbolTable[sym.ID] = sym

		structGraph.AddNode(&graph.Node{
			ID:   sym.ID,
			Type: graph.FunctionNode,
			Name: sym.Name,
			Path: sym.Path,
		})
		structGraph.AddEdge(sym.Path, sym.ID, graph.ContainsEdge)
	}

	// resolve facts to build the call graph
	for _, fact := range facts {
		callerID := findCaller(fact, symbols)
		if callerID == "" {
			continue
		}

		calleeID := findCallee(fact, symbols)
		if calleeID == "" {
			continue
		}

		callGraph.AddEdge(callerID, calleeID)
	}

	return Result{
		Structure: structGraph,
		Call:      callGraph,
	}
}

func findCaller(f Fact, symbols []Symbol) string {
	for _, sym := range symbols {
		if sym.Path == f.Path &&
			f.Line >= sym.StartLine &&
			f.Line <= sym.EndLine {
			return sym.ID
		}
	}
	return ""
}

func findCallee(f Fact, symbols []Symbol) string {
	for _, sym := range symbols {
		if sym.Name != f.CalleeName {
			continue
		}

		// same package heuristic (very basic)
		if filepath.Dir(sym.Path) == filepath.Dir(f.Path) {
			return sym.ID
		}

		// qualifier heuristic
		if f.CalleeQualifier != "" && strings.Contains(sym.ID, f.CalleeQualifier) {
			return sym.ID
		}
	}
	return ""
}
