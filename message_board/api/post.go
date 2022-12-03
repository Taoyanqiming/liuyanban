package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"message_board/model"
	"message_board/service"
	"message_board/util"
	"strconv"
)

func addPost(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)
	text := ctx.PostForm("text")

	post := model.Post{

		UserName: username,
		Text:     text,
	}

	err := service.AddPost(post)
	if err != nil {
		fmt.Println("insert post error", err)
		util.ResInterErr(ctx)
		return
	}
}

func post_id(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)
	text := ctx.PostForm("txt")
	post := model.Post{

		UserName: username,
		Text:     text,
	}
	err := service.ChangePost(post)
	if err != nil {
		fmt.Println("change post error", err)
		util.ResInterErr(ctx)
		return

	}
}
func briefPosts(ctx *gin.Context) {
	_, err := service.GetPost()
	if err != nil {
		fmt.Println("get posts err:", err)
		util.ResInterErr(ctx)
		return
	}
	util.RespOK(ctx)
}
func postDetail(ctx *gin.Context) {
	postIdString := ctx.Param("post_id")
	postId, err := strconv.Atoi(postIdString)
	if err != nil {
		fmt.Println("post id string to post id int error", err)
		util.RespOK(ctx)
	}
	posts, err := service.GetPostById(postId)
	if err != nil {
		fmt.Println("get post id by id err", err)
		util.ResInterErr(ctx)
		return
	}
	comments, err := service.GetPostComments(postId) //回复
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("get post comments err:", err)
			util.ResInterErr(ctx)
			return
		}
	}
	var PostDetail model.Post
	PostDetail.Post = posts
	PostDetail.Comment = comments
	fmt.Println("123")
	util.Err(ctx)
}
