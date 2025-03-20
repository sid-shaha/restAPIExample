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
func (p *PostServiceImpl) CreateNewPost(post models.Post) error {
	return p.postRepository.AddNewPost(post)
}
func (p *PostServiceImpl) GetAllPosts() ([]models.Post, error) {
	return p.postRepository.GetAllPosts()
}
func (p *PostServiceImpl) GetPostByID(id int) (models.Post, error) {
	return p.postRepository.GetPostByID(id)
}
func (p *PostServiceImpl) DeletePost(id int) error {
	return p.postRepository.DeletePost(id)
}
func (p *PostServiceImpl) UpdatePost(id int, post *models.Post) error {
	return p.postRepository.UpdatePost(id, post)
}
