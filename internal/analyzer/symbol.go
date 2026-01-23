package analyzer

// eg. func, struct, etc.
type SymbolKind string

// just for type safety
const (
	FunctionSymbol SymbolKind = "function"
	StructSymbol   SymbolKind = "struct"
	// will add things like VariableSymbol etc. 
)

// symbol is a definition found in the code
type Symbol struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Kind      SymbolKind `json:"kind"`
	Path      string     `json:"path"`
	StartLine int        `json:"startLine"`
	EndLine   int        `json:"endLine"`

	// can add language specific metadata here
}

type SymbolTable map[string]Symbol

// fact is just an interaction in the code
type Fact struct {
	Path            string `json:"path"`
	Line            int    `json:"line"`
	CalleeName      string `json:"calleeName"`
	CalleeQualifier string `json:"calleeQualifier,omitempty"`
}

// SymbolExtractor extracts symbols from files
type SymbolExtractor interface {
	ExtractSymbols(files []string) ([]Symbol, error)
}

// FactExtractor extracts facts from files
type FactExtractor interface {
	ExtractFacts(files []string) ([]Fact, error)
}

// LanguageExtractor combines both for a complete language plugin
type LanguageExtractor interface {
	SymbolExtractor
	FactExtractor
}
