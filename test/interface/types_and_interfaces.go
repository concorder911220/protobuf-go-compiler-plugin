package app

type CategoryData struct {
	id   PaymentCategoryId
	name string
}
type ListRequest struct {
}
type ListResponse struct {
	category []*CategoryData
}
type Main interface {
	list(req *ListRequest) (*ListResponse, error)
}
