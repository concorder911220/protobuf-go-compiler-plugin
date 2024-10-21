package claim

import (
	"context"
	"github.com/NewGlassbiller/gb-services-insurance/internal/claim/app"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"reflect"
)

type Service struct {
	grpcgen.UnimplementedServer
	command *app.Command
	query   *app.Query
}

func NewService(command *app.Command, query *app.Query) *Service {
	return &Service{
		command: command,
		query:   query,
	}
}
func (s *ClaimService) CreateClaim(ctx context.Context, req *dto.CreateClaimRequest) (*dto.CreateClaimResponse, error) {
	authUser, _ := ctx.Value(common.CtxUserKey{}).(common.UserInfo)
	return s.command.CreateClaim(&authUser, req)
}
func (s *ClaimService) UpdateClaim(ctx context.Context, req *dto.UpdateClaimRequest) (*dto.UpdateClaimResponse, error) {
	authUser, _ := ctx.Value(common.CtxUserKey{}).(common.UserInfo)
	return s.command.UpdateClaim(&authUser, req)
}
func (s *ClaimService) GetClaim(ctx context.Context, req *dto.GetClaimRequest) (*dto.GetClaimResponse, error) {
	authUser, _ := ctx.Value(common.CtxUserKey{}).(common.UserInfo)
	return s.command.GetClaim(&authUser, req)
}
func (s *ClaimService) GetJobLastClaim(ctx context.Context, req *dto.GetJobLastClaimRequest) (*dto.GetClaimResponse, error) {
	authUser, _ := ctx.Value(common.CtxUserKey{}).(common.UserInfo)
	return s.command.GetJobLastClaim(&authUser, req)
}
func (s *ClaimService) ListClaim(ctx context.Context, req *dto.ListClaimRequest) (*dto.ListClaimResponse, error) {
	authUser, _ := ctx.Value(common.CtxUserKey{}).(common.UserInfo)
	return s.command.ListClaim(&authUser, req)
}
func (s *ClaimService) AcknowledgeClaim(ctx context.Context, req *dto.AcknowledgeClaimRequest) (*dto.AcknowledgeClaimResponse, error) {
	authUser, _ := ctx.Value(common.CtxUserKey{}).(common.UserInfo)
	return s.command.AcknowledgeClaim(&authUser, req)
}
func (s *ClaimService) UpdateCoverage(ctx context.Context, req *dto.UpdateCoverageRequest) (*dto.UpdateCoverageResponse, error) {
	authUser, _ := ctx.Value(common.CtxUserKey{}).(common.UserInfo)
	return s.command.UpdateCoverage(&authUser, req)
}
func (s *ClaimService) ListInsurer(ctx context.Context, req *dto.ListInsurerRequest) (*dto.ListInsurerResponse, error) {
	authUser, _ := ctx.Value(common.CtxUserKey{}).(common.UserInfo)
	return s.command.ListInsurer(&authUser, req)
}
