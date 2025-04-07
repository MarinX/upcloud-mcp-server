package upcloud

import (
	"github.com/UpCloudLtd/upcloud-go-api/v8/upcloud/client"
	"github.com/UpCloudLtd/upcloud-go-api/v8/upcloud/service"
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

	s.AddTool(getAccount(svc))

	return s
}
