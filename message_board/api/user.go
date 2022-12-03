package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"message_board/model"
	"message_board/service"
	"message_board/util"
)

func Register(c *gin.Context) {
	username := c.PostForm("name")
	password := c.PostForm("password") //接下来判断

	if username == "" || password == "" {
		//响应resp
		util.ResParamErr(c)
		return
	}

	u, err := service.SearchUserByUserName(username) //注册时判断账户是否已经存在？？
	if err != nil && err != sql.ErrNoRows {
		util.RespInternalErr(c)
		return
	}

	if u.UserName != "" {
		util.NormErr(c, 300, "账户已存在")
		return
	}

	arr := service.CreateUser(model.User{
		UserName: username,
		PassWord: username,
	})
	if arr != nil {
		util.ResParamErr(c)
		return
	}
	util.RespOK(c)
}

func changePassword(ctx *gin.Context) {
	oldPassword := ctx.PostForm("old_password")
	newPassword := ctx.PostForm("new_password")
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)
	flag, err := service.IsPasswordCorrect(username, oldPassword) //判断
	if err != nil {
		fmt.Println("密码错误：", err)
		util.ResParamErr(ctx)
		return

	}
	if !flag {
		util.Err(ctx)
		return
	}
	err = service.ChangePassword(username, newPassword)
	if err != nil {
		fmt.Println("修改密码失败", err)
		util.ResParamErr(ctx)
		return
	}
	util.RespOK(ctx)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	_, err := service.IsPasswordCorrect(username, password) //判断
	if err != nil {
		fmt.Println("密码错误：", err)
		util.ResParamErr(c)
		return

	}
}
