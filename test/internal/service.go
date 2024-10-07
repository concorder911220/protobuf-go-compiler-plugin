package claim

import (
	"context"

	"github.com/NewGlassbiller/gb-go-common/gbnet/common"
	"github.com/NewGlassbiller/gb-services-insurance/internal/claim/app"
	dto "github.com/NewGlassbiller/gb-services-insurance/pkg/proto/claim/dto"
	grpcgen "github.com/NewGlassbiller/gb-services-insurance/pkg/proto/claim/grpc"
)

type MainService struct {
	grpcgen.UnimplementedMainServer
	command *app.Command
	query   *app.Query
}

func NewMainService(command *app.Command, query *app.Query) *MainService {
	return &MainService{
		command: command,
		query:   query,
	}
}
func (s *MainService) list(ctx context.Context, req *dto.ListRequest) (*dto.ListResponse, error) {
	authUser, _ := ctx.Value(common.CtxUserKey{}).(common.UserInfo)
	return s.command.list(&authUser, req)
}
