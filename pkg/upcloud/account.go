package upcloud

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/UpCloudLtd/upcloud-go-api/v8/upcloud/service"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func getAccount(svc *service.Service) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_account",
			mcp.WithDescription("Get current user account"),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			ac, err := svc.GetAccount(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get account: %w", err)
			}

			r, err := json.Marshal(ac)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal account: %w", err)
			}
			return mcp.NewToolResultText(string(r)), nil
		}
}
