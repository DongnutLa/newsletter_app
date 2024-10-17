package domain

type Pagination struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
	HasNext  bool  `json:"hasNext"`
	Length   int64 `json:"length"`
}

type PaginatedResponse[T any] struct {
	Metadata Pagination `json:"metadata"`
	Data     []T        `json:"data"`
}

type PaginationsParams struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
}
