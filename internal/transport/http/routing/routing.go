package routing

import (
	"github.com/alonsoF100/shotlink/internal/transport/http/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	router.POST("/api/shorten", handler.CreateShortURL)
	router.GET("/:code", handler.Redirect)
	router.GET("/api/links/:code", handler.GetLinkInfo)

	return router
}
