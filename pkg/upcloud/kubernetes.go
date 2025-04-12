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

// GetKubernetesClusters retrieves all kubernetes clusters.
func GetKubernetesClusters(svc *service.Service) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_kubernetes_clusters",
			mcp.WithDescription("Get all kubernetes clusters"),
		), func(ctx context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			clusters, err := svc.GetKubernetesClusters(ctx, &upreq.GetKubernetesClustersRequest{})
			if err != nil {
				return nil, fmt.Errorf("failed to get kubernetes clusters: %w", err)
			}

			r, err := json.Marshal(clusters)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal kubernetes clusters: %w", err)
			}
			return mcp.NewToolResultText(string(r)), nil
		}
}

// GetKubernetesCluster retrieves kubernetes cluster by UUID.
func GetKubernetesCluster(svc *service.Service) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_kubernetes_cluster",
			mcp.WithDescription("Get kubernetes cluster"),
			mcp.WithString("uuid", mcp.Required(), mcp.Description("The UUID of the kubernetes cluster")),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			uuid, err := requiredParam[string](request, "uuid")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			cluster, err := svc.GetKubernetesCluster(ctx, &upreq.GetKubernetesClusterRequest{
				UUID: uuid,
			})
			if err != nil {
				return nil, fmt.Errorf("failed to get kubernetes cluster: %w", err)
			}

			r, err := json.Marshal(cluster)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal kubernetes cluster: %w", err)
			}
			return mcp.NewToolResultText(string(r)), nil
		}
}

// GetKubernetesPlans retrieves all kubernetes plans.
func GetKubernetesPlans(svc *service.Service) (tool mcp.Tool, handler server.ToolHandlerFunc) {
	return mcp.NewTool("get_kubernetes_plans",
			mcp.WithDescription("Get kubernetes plans"),
		), func(ctx context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			plans, err := svc.GetKubernetesPlans(ctx, &upreq.GetKubernetesPlansRequest{})
			if err != nil {
				return nil, fmt.Errorf("failed to get kubernetes plans: %w", err)
			}

			r, err := json.Marshal(plans)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal kubernetes plans: %w", err)
			}
			return mcp.NewToolResultText(string(r)), nil
		}
}
