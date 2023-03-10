package handler

type Pagination struct {
	Page             int    `json:"page" form:"page"`
	NumberOfPages    int    `json:"numberOfPages"`
	PaginationString string `json:"paginationString"`
	MaxPageSize      int    `json:"maxPageSize" form:"maxPageSize"`
	TotalElements    int    `json:"totalElements"`
}
