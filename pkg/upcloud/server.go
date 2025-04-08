package upcloud

import (
	"fmt"

	"github.com/UpCloudLtd/upcloud-go-api/v8/upcloud/client"
	"github.com/UpCloudLtd/upcloud-go-api/v8/upcloud/service"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// NewServer creates a new Upcloud MCP server with the specified Upcloud client and logger.
func NewServer(client *client.Client, version string, readOnly bool) *server.MCPServer {
	// Create a new MCP server
	s := server.NewMCPServer(
		"upcloud-mcp-server",
		version,
		server.WithResourceCapabilities(true, true),
		server.WithLogging())

	svc := service.New(client)

	// account
	s.AddTool(getAccount(svc))
	s.AddTool(getAccountDetails(svc))
	s.AddTool(getAccountList(svc))

	// database
	s.AddTool(getDatabase(svc))

	return s
}

func requiredParam[T comparable](r mcp.CallToolRequest, p string) (T, error) {
	var zero T

	// Check if the parameter is present in the request
	if _, ok := r.Params.Arguments[p]; !ok {
		return zero, fmt.Errorf("missing required parameter: %s", p)
	}

	// Check if the parameter is of the expected type
	if _, ok := r.Params.Arguments[p].(T); !ok {
		return zero, fmt.Errorf("parameter %s is not of type %T", p, zero)
	}

	if r.Params.Arguments[p].(T) == zero {
		return zero, fmt.Errorf("missing required parameter: %s", p)
	}

	return r.Params.Arguments[p].(T), nil
}
