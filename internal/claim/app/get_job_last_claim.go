package app

import (
	"github.com/NewGlassbiller/gb-services-insurance/internal/claim"
	"github.com/NewGlassbiller/gb-services-insurance/internal/claim/domain"
	"gorm.io/gorm"
	"time"
)

func (query *Query) GetJobLastClaim(req *dto.GetJobLastClaimRequest) (*dto.GetClaimResponse, error) {
	// TODO: Add your logic here for GetJobLastClaim
	return nil, nil
}
