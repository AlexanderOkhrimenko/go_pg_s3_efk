package handler

import (
	"github.com/gin-gonic/gin"
	"go_pg_s3_efk/internal/service"
	"log"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	gin.SetMode(gin.DebugMode)

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/url.insert", h.addUrl)
		v1.GET("/ping", h.Ping)

	}
	err := router.Run(":8080")
	if err != nil {
		log.Println(err)
	}
	return router
}
