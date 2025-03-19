package web

import (
	"github.com/gin-gonic/gin"
	"github.com/sid-shaha/restAPIExample/handlers"
)

type RouterConfig struct {
	PostHandler *handlers.PostHandler
}

func NewRouter(config RouterConfig) *gin.Engine {
	router := gin.Default()

	router.Use(gin.Recovery(), gin.Logger())

	RegisterPostRoutes(router, config.PostHandler)

	return router
}
