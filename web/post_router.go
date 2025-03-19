package web

import (
	"github.com/gin-gonic/gin"
	"github.com/sid-shaha/restAPIExample/handlers"
)

// RegisterPostRoutes defines post-related routes.
func RegisterPostRoutes(router *gin.Engine, postHandler *handlers.PostHandler) {
	postRoutes := router.Group("/posts")
	{
		postRoutes.POST("/", postHandler.CreatePost)
		postRoutes.GET("/", postHandler.GetAllPosts)
		postRoutes.GET("/:id", postHandler.GetPostByID)
		postRoutes.PUT("/:id", postHandler.UpdatePost)
		postRoutes.DELETE("/:id", postHandler.DeletePost)
	}
}
