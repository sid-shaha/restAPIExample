package repository

import (
	"github.com/sid-shaha/restAPIExample/models"
)

type PostRepositoryImpl struct {
	postsMap *map[int]models.Post
}

func NewPostRepository(posts *map[int]models.Post) *PostRepositoryImpl {
	return &PostRepositoryImpl{
		postsMap: posts,
	}
}
func (p *PostRepositoryImpl) GetAllPosts() ([]models.Post, error) {
	return make([]models.Post, 0), nil
}
func (p *PostRepositoryImpl) GetPostByID(id int) (models.Post, error) {
	return models.Post{}, nil
}
func (p *PostRepositoryImpl) DeletePost() error {
	return nil
}
func (p *PostRepositoryImpl) UpdatePost(id int, post *models.Post) error {
	return nil
}
