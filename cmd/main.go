package main

import (
	"zadaie/internal/database"
	"zadaie/internal/handler"
	"zadaie/internal/repository"
	"zadaie/internal/service"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "zadaie/docs"
)

func main() {
	database.Connect()

	repo := repository.SubscriptionRepository{
		DB: database.DB,
	}

	svc := service.SubscriptionService{
		Repo: &repo,
	}

	h := handler.SubscriptionHandler{
		Service: &svc,
	}

	r := gin.Default()

	r.POST("/subscriptions", h.Create)
	r.GET("/subscriptions", h.GetAll)
	r.GET("/subscriptions/total", h.GetTotalCost)

	r.GET("/subscriptions/:id", h.GetByID)
	r.PUT("/subscriptions/:id", h.Update)
	r.DELETE("/subscriptions/:id", h.Delete)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
