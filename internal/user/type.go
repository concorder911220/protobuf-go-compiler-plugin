package user

type UserStatus int32

const (
	USER_STATUS_NEW      UserStatus = 0
	USER_STATUS_PENDING  UserStatus = 1
	USER_STATUS_COVERAGE UserStatus = 2
	USER_STATUS_SUCCESS  UserStatus = 3
	USER_STATUS_ERROR    UserStatus = 4
)
