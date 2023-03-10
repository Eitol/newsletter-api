package httpserver

import (
	"github.com/Eitol/newsletter-api/pkg/cors"
	"github.com/Eitol/newsletter-api/pkg/newsletter/handler"
	"github.com/gin-gonic/gin"
	"strconv"
)

func RunHttpServer(port int) {
	r := gin.Default()

	r.Use(cors.AllowCORS())

	newsletterHandler := handler.Build()

	libGroup := r.Group("/newsletter")
	subscriptionsGroup := libGroup.Group("/subscriptions")
	subscriptionsGroup.GET("", newsletterHandler.Get)
	subscriptionsGroup.POST("", newsletterHandler.Post)

	err := r.Run(":" + strconv.Itoa(port))
	if err != nil {
		panic(err)
	}
}
