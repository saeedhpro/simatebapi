package requests

type DeleteMultipleItemRequest struct {
	Ids []uint64 `json:"ids"`
}
