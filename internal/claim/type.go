package claim

import timestamppb "google.golang.org/protobuf/types/known/timestamppb"

type ClaimStatus int32

const (
	CLAIM_STATUS_NEW      ClaimStatus = 0
	CLAIM_STATUS_PENDING  ClaimStatus = 1
	CLAIM_STATUS_COVERAGE ClaimStatus = 2
	CLAIM_STATUS_SUCCESS  ClaimStatus = 3
	CLAIM_STATUS_ERROR    ClaimStatus = 4
)

type OpeningTypeCode int32

const (
	OPENING_TYPE_CODE_WR OpeningTypeCode = 0
	OPENING_TYPE_CODE_WS OpeningTypeCode = 1
	OPENING_TYPE_CODE_DR OpeningTypeCode = 2
	OPENING_TYPE_CODE_SI OpeningTypeCode = 3
	OPENING_TYPE_CODE_VN OpeningTypeCode = 4
	OPENING_TYPE_CODE_QT OpeningTypeCode = 5
	OPENING_TYPE_CODE_PT OpeningTypeCode = 6
	OPENING_TYPE_CODE_BK OpeningTypeCode = 7
	OPENING_TYPE_CODE_RF OpeningTypeCode = 8
)

type PositionCode int32

const (
	POSITION_CODE_B PositionCode = 0
	POSITION_CODE_F PositionCode = 1
	POSITION_CODE_M PositionCode = 2
	POSITION_CODE_R PositionCode = 3
)

type RelativeLocationCode int32

const (
	RELATIVE_LOCATION_CODE_IN RelativeLocationCode = 0
	RELATIVE_LOCATION_CODE_LO RelativeLocationCode = 1
	RELATIVE_LOCATION_CODE_OU RelativeLocationCode = 2
	RELATIVE_LOCATION_CODE_UP RelativeLocationCode = 3
)

type SideCode int32

const (
	SIDE_CODE_L SideCode = 0
	SIDE_CODE_R SideCode = 1
	SIDE_CODE_C SideCode = 2
)

type VehicleOwnership int32

const (
	VEHICLE_OWNERSHIP_OWNER       VehicleOwnership = 0
	VEHICLE_OWNERSHIP_COMMERCIAL  VehicleOwnership = 1
	VEHICLE_OWNERSHIP_RENTAL      VehicleOwnership = 2
	VEHICLE_OWNERSHIP_BORROWED    VehicleOwnership = 3
	VEHICLE_OWNERSHIP_THIRD_PARTY VehicleOwnership = 4
)

type VehicleType int32

const (
	VEHICLE_TYPE_STANDARD VehicleType = 0
	VEHICLE_TYPE_RV       VehicleType = 1
)

type CauseOfLossCode int32

const (
	CAUSE_OF_LOSS_101 CauseOfLossCode = 0
	CAUSE_OF_LOSS_105 CauseOfLossCode = 1
	CAUSE_OF_LOSS_111 CauseOfLossCode = 2
	CAUSE_OF_LOSS_121 CauseOfLossCode = 3
	CAUSE_OF_LOSS_131 CauseOfLossCode = 4
	CAUSE_OF_LOSS_199 CauseOfLossCode = 5
	CAUSE_OF_LOSS_201 CauseOfLossCode = 6
	CAUSE_OF_LOSS_301 CauseOfLossCode = 7
	CAUSE_OF_LOSS_311 CauseOfLossCode = 8
	CAUSE_OF_LOSS_321 CauseOfLossCode = 9
	CAUSE_OF_LOSS_331 CauseOfLossCode = 10
	CAUSE_OF_LOSS_341 CauseOfLossCode = 11
	CAUSE_OF_LOSS_401 CauseOfLossCode = 12
	CAUSE_OF_LOSS_411 CauseOfLossCode = 13
	CAUSE_OF_LOSS_421 CauseOfLossCode = 14
	CAUSE_OF_LOSS_431 CauseOfLossCode = 15
	CAUSE_OF_LOSS_499 CauseOfLossCode = 16
	CAUSE_OF_LOSS_901 CauseOfLossCode = 17
)

type CreateClaimRequest struct {
	job_id                  int32
	provider_name           string
	customer                *Customer
	vehicle                 *Vehicle
	insurance               *Insurance
	agent                   *Agent
	additional_info         *AdditionalInfo
	acknowledgement_details *AcknowledgementDetails
	damage_info             []*DamageInfo
}
type CreateClaimResponse struct {
	id int32
}
type UpdateClaimRequest struct {
	id       int32
	archived *bool
}
type UpdateClaimResponse struct {
}
type GetClaimRequest struct {
	id int32
}
type GetJobLastClaimRequest struct {
	job_id int32
}
type GetClaimResponse struct {
	claim *ClaimData
}
type ListClaimRequest struct {
	page_size        int32
	page             int32
	order_by         *OrderBy
	shop_id          *int32
	status           *ClaimStatus
	id               *string
	job_id           *string
	reference_number *string
	insurer_name     *string
	insurer_phone    *string
	customer_name    *string
	vehicle          *string
	submitted_dt     *string
	archived         *bool
}
type OrderBy struct {
	field string
	desc  bool
}
type ListClaimResponse struct {
	claim       []*ClaimShortData
	total_count int32
}
type AcknowledgeClaimRequest struct {
	id                      int32
	acknowledgement_details *AcknowledgementDetails
}
type AcknowledgeClaimResponse struct {
}
type ClaimData struct {
	id                      int32
	job_id                  int32
	created_user_id         int32
	created_dt              *timestamppb.Timestamp
	unique_id               string
	shop_id                 int32
	provider_name           string
	customer                *Customer
	vehicle                 *Vehicle
	insurance               *Insurance
	agent                   *Agent
	additional_info         *AdditionalInfo
	damage_info             []*DamageInfo
	result                  *ClaimResult
	acknowledgement_details *AcknowledgementDetails
	archived                bool
}
type Customer struct {
	id             int32
	first_name     string
	last_name      string
	phone1         string
	phone2         string
	state          string
	city           string
	street_address string
	postal_code    string
}
type Vehicle struct {
	id           int32
	vin          string
	year         int32
	make         string
	model        string
	plate_number string
	image_url    string
	thumb_url    string
	ownership    VehicleOwnership
	nags_id      string
	vehicle_type VehicleType
	number       int32
	style        string
	make_id      int32
	model_id     int32
}
type Insurance struct {
	insurer_id    string
	policy_number string
	policy_state  string
	date_of_loss  *timestamppb.Timestamp
}
type Agent struct {
	first_name  string
	last_name   string
	phone       string
	agency_name string
}
type AdditionalInfo struct {
	glass_only_damage        bool
	bodily_injury            bool
	origination_contact_name string
	cause_of_loss_code       CauseOfLossCode
	destination_pid          int32
	subrogation_contact_name string
	subrogation_data         string
}
type DamageInfo struct {
	opening_type_code      OpeningTypeCode
	position_code          PositionCode
	relative_location_code RelativeLocationCode
	side_code              SideCode
	glass_damage_quantity  int32
	repair_qualification   bool
}
type ClaimResult struct {
	reference_number          string
	error_code                *string
	status                    ClaimStatus
	error_message             string
	insurer_phone             string
	coverage_response_dt      **timestamppb.Timestamp
	claim_request             string
	claim_response            string
	coverage_response         string
	acknowledgement_dt        **timestamppb.Timestamp
	acknowledgement_request   string
	acknowledgement_response  string
	coverage_response_vehicle *CoverageResponseVehicle
}
type CoverageResponseVehicle struct {
	error_code         string
	error_text         string
	deductible         interface{}
	nags_id            string
	number             int32
	alternative_number string
	description        string
}
type ClaimShortData struct {
	job_id           int32
	id               int32
	submitted_dt     *timestamppb.Timestamp
	status           ClaimStatus
	error_message    string
	reference_number string
	customer_name    string
	vehicle          string
	insurer_name     string
	insurer_phone    string
	archived         bool
}
type UpdateCoverageRequest struct {
	provider_name     string
	coverage_response string
}
type UpdateCoverageResponse struct {
}
type AcknowledgementDetails struct {
	mobile_indicator          bool
	acceptance_contact        string
	work_location_postal_code string
	requested_appointment_dt  *timestamppb.Timestamp
}
type ListInsurerRequest struct {
	provider_name string
	active        *bool
	enol          *bool
}
type ListInsurerResponse struct {
	insurer []*InsurerData
}
type InsurerData struct {
	provider_name string
	external_id   string
	name          string
	active        bool
	enol          bool
	last_updated  string
}
type Claim interface {
	CreateClaim(req *CreateClaimRequest) (*CreateClaimResponse, error)
	UpdateClaim(req *UpdateClaimRequest) (*UpdateClaimResponse, error)
	GetClaim(req *GetClaimRequest) (*GetClaimResponse, error)
	GetJobLastClaim(req *GetJobLastClaimRequest) (*GetClaimResponse, error)
	ListClaim(req *ListClaimRequest) (*ListClaimResponse, error)
	AcknowledgeClaim(req *AcknowledgeClaimRequest) (*AcknowledgeClaimResponse, error)
	UpdateCoverage(req *UpdateCoverageRequest) (*UpdateCoverageResponse, error)
	ListInsurer(req *ListInsurerRequest) (*ListInsurerResponse, error)
}
