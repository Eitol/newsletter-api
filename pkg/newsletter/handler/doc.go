package handler

type ResponseDoc struct {
	Filter     *FilterDoc     `json:"filter"`
	Pagination *PaginationDoc `json:"pagination"`
	Results    []*ResultsDoc  `json:"results"`
}

type FilterDoc struct {
	UserID    string   `json:"userId" example:"e020e7f8-79e6-4d16-80ce-7cbf86cefe1f"`
	BlogID    string   `json:"blogId" example:"b4f2e2d9-e399-46d0-a458-ef75527896c1"`
	Interests []string `json:"interests" example:"tech,sports"`
}

//nolint:lll // Long documentation lines
type ResultsDoc struct {
	UserID    string   `json:"userId" example:"e020e7f8-79e6-4d16-80ce-7cbf86cefe1f"`
	BlogID    string   `json:"blogId" example:"b4f2e2d9-e399-46d0-a458-ef75527896c1"`
	Interests []string `json:"interests" example:"tech,sports"`
}

type PaginationDoc struct {
	Page             int    `json:"page"`
	NumberOfPages    int    `json:"numberOfPages"`
	PaginationString string `json:"paginationString" example:"1/1"`
	MaxPageSize      int    `json:"maxPageSize"`
	TotalElements    int    `json:"totalElements"`
}
