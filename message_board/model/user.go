package model

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	PassWord string `json:"password"`
}
