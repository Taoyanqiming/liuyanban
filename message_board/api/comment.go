package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message_board/model"
	"message_board/service"
	"message_board/util"
	"strconv"
	"time"
)

func addComment(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)
	txt := ctx.PostForm("txt")
	postIdString := ctx.PostForm("post_id")
	postId, err := strconv.Atoi(postIdString)
	if err != nil {
		fmt.Println("post id to int error", err)
		util.ResParamErr(ctx)
		return
	}
	comment := model.Comment{
		PostId:      postId,
		Txt:         txt,
		UserName:    username,
		CommentTime: time.Now(),
	}

	err = service.AddComment(comment)
	if err != nil {
		fmt.Println("add comment err:", err)
		util.ResInterErr(ctx)
		return
	}
	util.RespOK(ctx)
}
func deleteComment(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)
	txt := ctx.PostForm("txt")
	postIdString := ctx.PostForm("post_id")
	postId, err := strconv.Atoi(postIdString)
	if err != nil {
		fmt.Println("post id to int error")
		util.ResParamErr(ctx) //id错误
		return
	}
	comment := model.Comment{
		PostId:      postId,
		Txt:         txt,
		UserName:    username,
		CommentTime: time.Now(),
	}
	err = service.DeleteComment(comment)
	if err != nil {
		fmt.Println("delete common error", err)
		util.RespInternalErr(ctx)
		return
	}
}
func changeComment(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)
	txt := ctx.PostForm("txt")
	postIdString := ctx.PostForm("post_id")
	postId, err := strconv.Atoi(postIdString)
	if err != nil {
		fmt.Println("post id to int error")
		util.ResParamErr(ctx) //id有误
		return
	}
	comment := model.Comment{
		PostId:      postId,
		Txt:         txt,
		UserName:    username,
		CommentTime: time.Now(),
	}
	err = service.Changecomment(comment)
	if err != nil {
		fmt.Println("change comment default:")
		util.ResInterErr(ctx) //失败

	}
}
