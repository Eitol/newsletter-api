package handler

import (
	"github.com/Eitol/newsletter-api/pkg/newsletter"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type PostRequestBody struct {
	UserID    uuid.UUID `json:"userId"`
	BlogID    uuid.UUID `json:"blogId"`
	Interests []string  `json:"interests"`
}

type PostResponseBody struct {
	SubscriptionID uuid.UUID `json:"subscriptionId" example:"e020e7f8-79e6-4d16-80ce-7cbf86cefe1f"`
}

// Post godoc
// @Summary Create a new subscription
// @Tags subscriptions
// @Router /subscriptions [post]
// @Param body PostRequestBody
// @Produce json
// @Success 201 handler.PostResponseBody
func (h *handler) Post(ctx *gin.Context) {
	var reqBody PostRequestBody
	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	interests := make([]newsletter.Interest, len(reqBody.Interests))
	for i, interest := range reqBody.Interests {
		interests[i] = newsletter.Interest(interest)
	}
	newSubscriptionUUID, err := h.svc.Post(ctx, reqBody.UserID, reqBody.BlogID, interests)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := PostResponseBody{
		SubscriptionID: newSubscriptionUUID,
	}
	ctx.JSON(http.StatusCreated, response)
}
