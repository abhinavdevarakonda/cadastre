package analyzer

import "fmt"

func Analyze(root string) Result {
	// Phase 1: FS Scan (Discovery)
	scan, err := Scan(root)
	if err != nil {
		panic(fmt.Sprintf("failed to scan directory: %v", err))
	}

	// Phase 2: Symbols (Definitions)
	// For now, we only have the Go extractor. 
	// Later we can dispatch based on file extensions.
	goExt := &GoExtractor{}
	symbols, err := goExt.ExtractSymbols(scan.Files)
	if err != nil {
		panic(fmt.Sprintf("failed to extract symbols: %v", err))
	}

	// Phase 3: Facts (Interactions)
	facts, err := goExt.ExtractFacts(scan.Files)
	if err != nil {
		panic(fmt.Sprintf("failed to extract facts: %v", err))
	}

	// Phase 4: Graphs (Synthesis)
	return Build(scan, symbols, facts)
}
