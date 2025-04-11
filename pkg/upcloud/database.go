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

// GetDatabase retrieves the managed database information.
func GetDatabase(svc *service.Service) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_managed_database",
			mcp.WithDescription("Get managed database"),
			mcp.WithString("uuid", mcp.Required(), mcp.Description("The UUID of the database")),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			uuid, err := requiredParam[string](request, "uuid")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			db, err := svc.GetManagedDatabase(ctx, &upreq.GetManagedDatabaseRequest{UUID: uuid})
			if err != nil {
				return nil, fmt.Errorf("failed to get database: %w", err)
			}
			r, err := json.Marshal(db)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal database: %w", err)
			}
			return mcp.NewToolResultText(string(r)), nil
		}
}
