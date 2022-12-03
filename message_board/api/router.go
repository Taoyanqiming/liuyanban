package api

import "github.com/gin-gonic/gin"

func Initrouter() {
	r := gin.Default()
	u := r.Group("/user")
	{
		u.POST("/register", Register)      //注册
		u.POST("/login", Login)            //登录
		u.PUT("/password", changePassword) //修改密码
		u.POST("/post", addPost)           //留言
		u.PUT("changepost", post_id)
		u.GET("/briefpost", briefPosts)
		u.POST("addcomment", addComment)
		u.PUT("/changecomment", changeComment)
		u.POST("/deletecomment", deleteComment)
	}
	r.Run()
}
