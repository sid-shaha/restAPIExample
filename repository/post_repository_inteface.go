package repository

import "github.com/sid-shaha/restAPIExample/models"

type PostRepository interface {
	GetAllPosts() ([]models.Post, error)
	GetPostByID(id int) (models.Post, error)
	DeletePost() error
	UpdatePost(id int, post *models.Post) error
}
