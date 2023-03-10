package handler

type Response struct {
	FilterUserID    string      `json:"userId,omitempty"`
	FilterBlogID    string      `json:"blogId,omitempty"`
	FilterInterests []string    `json:"interests,omitempty"`
	Pagination      *Pagination `json:"pagination"`
	Results         []*Result   `json:"results"`
}

type Result struct {
	UserID    string   `json:"userId"`
	BlogID    string   `json:"blogId"`
	Interests []string `json:"interests"`
}
