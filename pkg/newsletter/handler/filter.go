package handler

type Filter struct {
	UserID    string   `json:"userId" form:"userId"`
	BlogID    string   `json:"blogId" form:"blogId"`
	Interests []string `json:"interests" form:"interests"`
}
