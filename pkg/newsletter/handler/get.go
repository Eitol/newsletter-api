package handler

import (
	"github.com/Eitol/newsletter-api/pkg/newsletter"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strconv"
)

type rawGetParams struct {
	Page        string
	MaxPageSize string
	UserId      string
	BlogId      string
	Interests   []string
}

type parsedGetParams struct {
	Page        int
	MaxPageSize int
	UserId      uuid.UUID
	BlogID      uuid.UUID
	Interests   []newsletter.Interest
}

func adaptGetParams(raw rawGetParams) (*parsedGetParams, error) {
	if raw.Page == "" {
		return nil, errEmptyPageParam
	}
	pageInt, err := strconv.Atoi(raw.Page)
	if err != nil {
		return nil, errInvalidPageParam
	}
	if pageInt < 1 {
		return nil, errPageParamLT1
	}
	if raw.MaxPageSize == "" {
		return nil, errEmptyMaxPageSizeParam
	}
	maxPageSizeInt, err := strconv.Atoi(raw.MaxPageSize)
	if err != nil {
		return nil, errInvalidMaxPageSizeParam
	}
	userId := uuid.Nil
	if raw.UserId != "" {
		userId, err = uuid.Parse(raw.UserId)
		if err != nil {
			return nil, errInvalidUserIdParam
		}
	}
	blogId := uuid.Nil
	if raw.BlogId != "" {
		blogId, err = uuid.Parse(raw.BlogId)
		if err != nil {
			return nil, errInvalidBlogIdParam
		}
	}
	var interests []newsletter.Interest
	for _, interest := range raw.Interests {
		interests = append(interests, newsletter.Interest(interest))
	}
	return &parsedGetParams{
		Page:        pageInt,
		MaxPageSize: maxPageSizeInt,
		UserId:      userId,
		BlogID:      blogId,
		Interests:   interests,
	}, nil
}

func extractGetParams(ctx *gin.Context) rawGetParams {
	page := ctx.Query("page")
	maxPageSize := ctx.Query("maxPageSize")
	userId := ctx.Query("userId")
	blogId := ctx.Query("blogId")
	interests := ctx.QueryArray("interests")
	return rawGetParams{
		Page:        page,
		MaxPageSize: maxPageSize,
		UserId:      userId,
		BlogId:      blogId,
		Interests:   interests,
	}
}

// nolint:lll // godoc
// Get godoc
// @Summary      Read subscriptions
// @Tags         subscriptions
// @Router       /subscriptions [get]
// @Param        page	        query  int		 true   "Result page (>=1)" example(1)
// @Param        maxPageSize	query  int		 true   "Max page size"     example(10)
// @Param        userId	        query  string	 false  "User ID"           example(1)
// @Param        blogId	        query  string	 false  "Blog ID"           example(1)
// @Param        interests	    query  []string  false  "Interests"         example(["tech","sports"])
// @Produce      json
// @Success      200  {array}  handler.ResponseDoc
// nolint:gocyclo //error checking branches
func (h *handler) Get(ctx *gin.Context) {
	rawParams := extractGetParams(ctx)
	adaptedReq, err := adaptGetParams(rawParams)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	srvResp, err := h.svc.Get(
		ctx, adaptedReq.UserId, adaptedReq.BlogID, adaptedReq.Interests,
		adaptedReq.Page, adaptedReq.MaxPageSize,
	)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	adaptedResponse, err := adaptGetResponse(srvResp, rawParams, *adaptedReq)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, adaptedResponse)
}

func adaptGetResponse(
	srvResp *newsletter.Result[*newsletter.Subscription],
	raw rawGetParams, parsed parsedGetParams,
) ([]ResponseDoc, error) {
	var adaptedResult []*ResultsDoc
	for _, e := range srvResp.Page.Elements {
		interestsStr := make([]string, len(e.Interests))
		for i, interest := range e.Interests {
			interestsStr[i] = string(interest)
		}
		adaptedResult = append(adaptedResult, &ResultsDoc{
			UserID:    e.UserID.String(),
			BlogID:    e.BlogID.String(),
			Interests: interestsStr,
		})

	}
	return []ResponseDoc{
		{
			Filter: &FilterDoc{
				UserID:    raw.UserId,
				BlogID:    raw.BlogId,
				Interests: raw.Interests,
			},
			Pagination: &PaginationDoc{
				Page:             parsed.Page,
				NumberOfPages:    srvResp.Pages,
				TotalElements:    srvResp.Total,
				PaginationString: "", // TODO: implement
				MaxPageSize:      parsed.MaxPageSize,
			},
			Results: adaptedResult,
		},
	}, nil
}
