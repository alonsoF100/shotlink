package routing

import (
	"github.com/alonsoF100/shotlink/internal/transport/http/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(service handler.Service, baseURL string) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	handler := handler.New(service, baseURL)

	router.POST("/shorten", handler.CreateShortURL)
	router.GET("/:code", handler.Redirect)
	
	return router
}
