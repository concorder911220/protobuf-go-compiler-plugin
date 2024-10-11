package app

import (
	"github.com/NewGlassbiller/gb-go-common/gbnet/common"
	"github.com/NewGlassbiller/gb-go-common/gbpolicy"
	"github.com/NewGlassbiller/gb-services-insurance/pkg/proto/main/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

// list handles ListRequest and returns ListResponse
func (cmd Command) list(authUser *common.UserInfo, req *dto.ListRequest) (*dto.ListResponse, error) {
	// TODO: Add your logic here for list
	return nil, nil
}
