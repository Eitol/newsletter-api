package handler

import (
	"github.com/gin-gonic/gin"
)

// nolint:lll // godoc
// Get godoc
// @Summary      Read subscriptions
// @Tags         subscriptions
// @Router       /subscriptions [get]
// @Param        page	        query  int		 true   "Result page"                                   example(1)
// @Param        maxPageSize	query  int		 true   "Max page size"                                  example(10)
// @Param        userId	        query  string	 false  "User ID"                                        example(1)
// @Param        blogId	        query  string	 false  "Blog ID"                                        example(1)
// @Param        interests	    query  []string  false  "Interests"                                      example(["tech","sports"])
// @Produce      json
// @Success      200  {array}  handler.ResponseDoc
// nolint:gocyclo //error checking branches
func (h *handler) Get(ctx *gin.Context) {
	panic("implement me")
}
