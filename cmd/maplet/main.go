package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/abhinavdevarakonda/maplet/internal/analyzer"
	"github.com/abhinavdevarakonda/maplet/internal/export"
	"github.com/abhinavdevarakonda/maplet/internal/server"
	mcpserver "github.com/mark3labs/mcp-go/server"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("	maplet analyze <path>")
		fmt.Println("	maplet export <path>")
		fmt.Println("	maplet serve <path>")
		fmt.Println("	maplet mcp <path>")
		return
	}

	command := os.Args[1]

	path := "."
	if len(os.Args) > 2 {
		path = os.Args[2]
	}

	switch command {
	case "analyze":
		result := analyzer.Analyze(path)

		fmt.Printf(
			"structure graph: %d nodes, %d edges\n",
			len(result.Structure.Nodes),
			len(result.Structure.Edges),
		)

		fmt.Printf(
			"call graph: %d functions\n",
			len(result.Call.Edges),
		)

	case "export":
		result := analyzer.Analyze(path)

		eg := export.FromGraph(result.Structure)
		eg.CallEdges = export.FromCallGraph(result.Call)

		data, err := json.MarshalIndent(eg, "", " ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(data))

	case "serve":
		result := analyzer.Analyze(path)

		eg := export.FromGraph(result.Structure)

		srv := server.New(eg)
		if err := srv.Start("localhost:6767"); err != nil {
			panic(err)
		}

	case "mcp":
		result := analyzer.Analyze(path)
		mcpSrv := server.NewMCPServer(result)
		stdioSrv := mcpserver.NewStdioServer(mcpSrv)
		if err := stdioSrv.Listen(context.Background(), os.Stdin, os.Stdout); err != nil {
			panic(err)
		}

	default:
		fmt.Println("unknown command:", command)
	}
}

