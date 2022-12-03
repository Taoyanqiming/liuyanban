package model

type Post struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	Text     string `json:"text"`
}
type PostDetail struct {
	Post
	Comment []Comment
}
