package service

import "github.com/sid-shaha/restAPIExample/models"

func GetAllPosts() ([]models.Post, error) {
	return make([]models.Post, 0), nil
}

func GetPostByID(id int) (models.Post, error) {
	return models.Post{}, nil
}
func UpdatePost(id int, post *models.Post) {

}
func DeletePost(id int) {

}
