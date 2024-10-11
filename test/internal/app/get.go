package app

import (
	"github.com/NewGlassbiller/gb-go-common/gbnet/common"
	"github.com/NewGlassbiller/gb-go-common/gbpolicy"
	"github.com/NewGlassbiller/gb-services-insurance/pkg/proto/main/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

// Get handles GetRequest and returns GetResponse
func (cmd Command) Get(authUser *common.UserInfo, req *dto.GetRequest) (*dto.GetResponse, error) {
	// TODO: Add your logic here for Get
	return nil, nil
}
