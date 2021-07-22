package singleprocess

import (
	"context"

	pb "github.com/hashicorp/waypoint/internal/server/gen"
)

func (s *service) UI_ListDeployments(
	ctx context.Context,
	req *pb.UI_ListDeploymentsRequest,
) (*pb.UI_ListDeploymentsResponse, error) {
	var result []*pb.UI_DeploymentBundle

	return &pb.UI_ListDeploymentsResponse{Deployments: result}, nil
}

// GetDeployment returns a Deployment based on ID
func (s *service) UI_GetDeployment(
	ctx context.Context,
	req *pb.UI_GetDeploymentRequest,
) (*pb.UI_DeploymentBundle, error) {
	return nil, nil
}
