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

// GetFirewallRules retrieves the firewall rules for a server.
func GetFirewallRules(svc *service.Service) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_firewall_rules",
			mcp.WithString("uuid", mcp.Required(), mcp.Description("The UUID of the server")),
			mcp.WithDescription("Get firewall rules"),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			uuid, err := requiredParam[string](request, "uuid")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			fwRules, err := svc.GetFirewallRules(ctx, &upreq.GetFirewallRulesRequest{
				ServerUUID: uuid,
			})
			if err != nil {
				return nil, fmt.Errorf("failed to get firewall rules: %w", err)
			}
			r, err := json.Marshal(fwRules)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal firewall rules: %w", err)
			}
			return mcp.NewToolResultText(string(r)), nil
		}
}

// GetFirewallRuleDetails retrieves the details of a specific firewall rule for a server.
// optional parameter: position
func GetFirewallRuleDetails(svc *service.Service) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_firewall_rule_details",
			mcp.WithString("uuid", mcp.Required(), mcp.Description("The UUID of the server")),
			mcp.WithNumber("position", mcp.Min(1), mcp.Description("The position of the firewall rule")),
			mcp.WithDescription("Get firewall rules details"),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			uuid, err := requiredParam[string](request, "uuid")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			position, err := RequiredInt(request, "position")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			fwRule, err := svc.GetFirewallRuleDetails(ctx, &upreq.GetFirewallRuleDetailsRequest{
				ServerUUID: uuid,
				Position:   position,
			})
			if err != nil {
				return nil, fmt.Errorf("failed to get firewall details: %w", err)
			}
			r, err := json.Marshal(fwRule)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal firewall details: %w", err)
			}
			return mcp.NewToolResultText(string(r)), nil
		}
}
