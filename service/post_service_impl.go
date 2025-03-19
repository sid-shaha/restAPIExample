package service

import (
	"github.com/sid-shaha/restAPIExample/models"
	"github.com/sid-shaha/restAPIExample/repository"
)

type PostServiceImpl struct {
	postRepository repository.PostRepository
}

func NewPostService(postRepo repository.PostRepository) *PostServiceImpl {
	return &PostServiceImpl{
		postRepository: postRepo,
	}
}

func (p *PostServiceImpl) GetAllPosts() ([]models.Post, error) {
	return make([]models.Post, 0), nil
}
func (p *PostServiceImpl) GetPostByID(id int) (models.Post, error) {
	return models.Post{}, nil
}
func (p *PostServiceImpl) DeletePost(id int) error {
	return nil
}
func (p *PostServiceImpl) UpdatePost(id int, post *models.Post) error {
	return nil
}
