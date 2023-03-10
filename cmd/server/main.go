package main

import (
	"os"

	"github.com/Eitol/newsletter-api/pkg/cors"
	"github.com/Eitol/newsletter-api/pkg/newsletter/handler"
	"github.com/gin-gonic/gin"
)

// @contact.name                Grupo MContigo
// @title                       Newsletter API
// @version                     1.0
// @description                 Newsletter API
func main() {
	r := gin.Default()

	r.Use(cors.AllowCORS())

	newsletterHandler := handler.Build()

	libGroup := r.Group("/newsletter")
	subscriptionsGroup := libGroup.Group("/subscriptions")
	subscriptionsGroup.GET("", newsletterHandler.Get)
	subscriptionsGroup.POST("", newsletterHandler.Post)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := r.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
