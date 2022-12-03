package api

import (
	"github.com/gin-gonic/gin"
	"message_board/util"
)

func auth(ctx *gin.Context) {
	username, err := ctx.Cookie("username")
	if err != nil {
		util.RespOK(ctx)
		return
	}
	ctx.Set("username", username)
	ctx.Next()
}
