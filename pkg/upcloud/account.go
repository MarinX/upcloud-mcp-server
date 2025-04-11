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

// GetAccount retrieves the current user account information.
func GetAccount(svc *service.Service) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_account",
			mcp.WithDescription("Get current user account"),
		), func(ctx context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

// GetAccountDetails retrieves the current user account details for a specific username.
func GetAccountDetails(svc *service.Service) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_account_details",
			mcp.WithDescription("Get current user account details"),
			mcp.WithString("username", mcp.Required(), mcp.Description("Username")),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			username, err := requiredParam[string](request, "username")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			acDetails, err := svc.GetAccountDetails(ctx, &upreq.GetAccountDetailsRequest{Username: username})
			if err != nil {
				return nil, fmt.Errorf("failed to get account details: %w", err)
			}
			r, err := json.Marshal(acDetails)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal account details: %w", err)
			}
			return mcp.NewToolResultText(string(r)), nil
		}
}

// GetAccountList retrieves the current user account list.
func GetAccountList(svc *service.Service) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_account_list",
			mcp.WithDescription("Get current user account list"),
		), func(ctx context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			list, err := svc.GetAccountList(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get account list: %w", err)
			}
			r, err := json.Marshal(list)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal account list: %w", err)
			}
			return mcp.NewToolResultText(string(r)), nil
		}
}
