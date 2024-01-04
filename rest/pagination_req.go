package rest

type PaginationReq struct {
	// Limit number of resource in the response
	Limit int64 `json:"limit" query:"limit" extensions:"x-order=101" example:"200"`
	// Offset number of resource in the response
	Offset int64 `json:"offset" query:"offset" extensions:"x-order=102" example:"0"`
}
