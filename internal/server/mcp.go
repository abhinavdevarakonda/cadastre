package server

import (
	"context"
	"fmt"
	"strings"

	"github.com/abhinavdevarakonda/maplet/internal/analyzer"
	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
)

func NewMCPServer(result analyzer.Result) *mcpserver.MCPServer {
	s := mcpserver.NewMCPServer(
		"maplet",
		"1.0.0",
	)

	// List all nodes
	s.AddTool(mcp.NewTool("list_nodes",
		mcp.WithDescription("List all nodes in the project structure (directories, files, functions)"),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		nodes := []string{}
		for id, node := range result.Structure.Nodes {
			nodes = append(nodes, fmt.Sprintf("%s (%s): %s", id, node.Type, node.Name))
		}
		return mcp.NewToolResultText(strings.Join(nodes, "\n")), nil
	})

	// Get node details
	s.AddTool(mcp.NewTool("get_node_details",
		mcp.WithDescription("Get detailed information about a specific node"),
		mcp.WithString("id", mcp.Description("The ID of the node"), mcp.Required()),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := request.RequireString("id")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		node, exists := result.Structure.Nodes[id]
		if !exists {
			return mcp.NewToolResultError(fmt.Sprintf("node %s not found", id)), nil
		}

		details := fmt.Sprintf("ID: %s\nType: %s\nName: %s\nPath: %s", node.ID, node.Type, node.Name, node.Path)
		return mcp.NewToolResultText(details), nil
	})

	// Get callers
	s.AddTool(mcp.NewTool("get_callers",
		mcp.WithDescription("Find functions that call the given function"),
		mcp.WithString("function_id", mcp.Description("The ID of the function (e.g., 'package.Receiver.Name')"), mcp.Required()),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		fnID, err := request.RequireString("function_id")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		callers := []string{}
		for caller, callees := range result.Call.Edges {
			for _, callee := range callees {
				if callee == fnID {
					callers = append(callers, caller)
				}
			}
		}

		if len(callers) == 0 {
			return mcp.NewToolResultText("No callers found"), nil
		}
		return mcp.NewToolResultText(strings.Join(callers, "\n")), nil
	})

	// Get callees
	s.AddTool(mcp.NewTool("get_callees",
		mcp.WithDescription("Find functions called by the given function"),
		mcp.WithString("function_id", mcp.Description("The ID of the function"), mcp.Required()),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		fnID, err := request.RequireString("function_id")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		callees, exists := result.Call.Edges[fnID]
		if !exists || len(callees) == 0 {
			return mcp.NewToolResultText("No callees found"), nil
		}

		return mcp.NewToolResultText(strings.Join(callees, "\n")), nil
	})

	return s
}
