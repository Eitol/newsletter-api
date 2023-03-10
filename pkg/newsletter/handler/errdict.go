package handler

import "errors"

var (
	errEmptyPageParam          = errors.New("page parameter is empty")
	errInvalidPageParam        = errors.New("page parameter is invalid")
	errPageParamLT1            = errors.New("page parameter is less than 1")
	errEmptyMaxPageSizeParam   = errors.New("maxPageSize parameter is empty")
	errInvalidMaxPageSizeParam = errors.New("maxPageSize parameter is invalid")
	errInvalidUserIdParam      = errors.New("userId parameter is invalid, must be a valid UUID")
	errInvalidBlogIdParam      = errors.New("blogId parameter is invalid, must be a valid UUID")
)
