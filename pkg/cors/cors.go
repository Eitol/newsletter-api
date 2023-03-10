package cors

import (
	"github.com/gin-gonic/gin"
	corsMid "github.com/itsjamie/gin-cors"
)

// AllowCORS enables Cross-Origin Resource Sharing
// See https://fetch.spec.whatwg.org/#http-cors-protocol and
// https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS for further info
func AllowCORS() gin.HandlerFunc {
	return corsMid.Middleware(
		corsMid.Config{
			Origins:         "*",
			RequestHeaders:  "Origin, Authorization, Content-Type, api_key",
			ExposedHeaders:  "",
			Methods:         "*",
			Credentials:     false,
			ValidateHeaders: false,
		},
	)
}
