package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sid-shaha/restAPIExample/service"
)

type PostHandler struct {
	PostService service.PostService
}

func NewPostHandler(service service.PostService) *PostHandler {
	return &PostHandler{PostService: service}
}
func (h *PostHandler) CreatePost(c *gin.Context) {

}
func (h *PostHandler) GetAllPosts(c *gin.Context) {
	//return
}

func (h *PostHandler) GetPostByID(c *gin.Context) {

}
func (h *PostHandler) UpdatePost(c *gin.Context) {

}
func (h *PostHandler) DeletePost(c *gin.Context) {

}
