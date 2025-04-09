package upcloud

import (
	"context"
	"encoding/json"
	"fmt"

	upreq "github.com/UpCloudLtd/upcloud-go-api/v8/upcloud/request"

	"github.com/UpCloudLtd/upcloud-go-api/v8/upcloud/service"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func getServers(svc *service.Service) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_servers",
			mcp.WithDescription("Get servers"),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			srvs, err := svc.GetServers(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get servers: %w", err)
			}
			r, err := json.Marshal(srvs)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal servers: %w", err)
			}
			return mcp.NewToolResultText(string(r)), nil
		}
}

func getServerDetails(svc *service.Service) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_server_details",
			mcp.WithDescription("Get server details"),
			mcp.WithString("uuid", mcp.Required(), mcp.Description("The UUID of the server")),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			uuid, err := requiredParam[string](request, "uuid")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			details, err := svc.GetServerDetails(ctx, &upreq.GetServerDetailsRequest{
				UUID: uuid,
			})
			if err != nil {
				return nil, fmt.Errorf("failed to get servers: %w", err)
			}
			r, err := json.Marshal(details)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal servers: %w", err)
			}
			return mcp.NewToolResultText(string(r)), nil
		}
}
