package service

import "github.com/sid-shaha/restAPIExample/models"

type PostService interface {
	GetAllPosts() ([]models.Post, error)
	GetPostByID(id int) (models.Post, error)
	DeletePost(id int) error
	UpdatePost(id int, post *models.Post) error
}
