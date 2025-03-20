package repository

import (
	"fmt"

	"github.com/sid-shaha/restAPIExample/models"
)

type PostRepositoryIMImpl struct {
	postsMap *map[int]models.Post
}

func NewPostIMRepository(posts *map[int]models.Post) *PostRepositoryIMImpl {
	return &PostRepositoryIMImpl{
		postsMap: posts,
	}
}
func (p *PostRepositoryIMImpl) AddNewPost(post models.Post) error {
	postMap := *p.postsMap
	val, ok := postMap[post.Id]

	if ok {
		return fmt.Errorf("POST %v Already exist with the given id %v", val, post.Id)
	}
	postMap[post.Id] = post
	return nil
}
func (p *PostRepositoryIMImpl) GetAllPosts() ([]models.Post, error) {
	postMap := *p.postsMap
	var allPosts []models.Post

	for _, v := range postMap {
		allPosts = append(allPosts, v)
	}

	return allPosts, nil
}
func (p *PostRepositoryIMImpl) GetPostByID(id int) (models.Post, error) {
	postMap := *p.postsMap
	val, ok := postMap[id]

	if !ok {
		return models.Post{}, fmt.Errorf("NO POST FOUND BY ID %v", id)
	}
	return val, nil
}
func (p *PostRepositoryIMImpl) DeletePost(id int) error {
	delete(*p.postsMap, id)
	return nil
}
func (p *PostRepositoryIMImpl) UpdatePost(id int, post *models.Post) error {
	postMap := *p.postsMap
	postMap[id] = *post
	return nil
}
