package service

import (
	"message_board/dao"
	"message_board/model"
)

func AddPost(post model.Post) error {
	return dao.InsetPost(post)
}
func ChangePost(post model.Post) error {
	return dao.ChangePost(post)
}
func GetPost() ([]model.Post, error) {
	return dao.SelectPosts()
}
func GetPostById(postId int) (model.Post, error) {
	return dao.SelectPostByPostId(postId)
}
