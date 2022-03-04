package controller

import (
	"Bilibili-project/middleware"
	"github.com/gin-gonic/gin"
)

func RouterInit() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/bilibili/v1")
	{
		v1.POST("/user/register", Register)
		v1.POST("/user/login", Login)
		v1.GET("/user/readOneVideo/:id", ReadOneVideo)
		v1.GET("/user/readVideo/:id", ReadVideo)
		v1.GET("/user/readComment/:id", ReadComment)
		v1.GET("/user/readReply/:id", ReadReply)
		v1.GET("/user/readBarrage", ReadBarrage)
		v1.POST("/admin/login", AdminLogin)
		v1.GET("/clickList", ReadClickList)
		v1.GET("/user/searchVideo", SearchVideo)
		v1.GET("/user/searchUser", SearchUser)
		v1.GET("/user/searchHistory", SearchHistory)
		v1.GET("/user/topicSearch", GetTopic)
		loginRequired := v1.Group("/login")
		loginRequired.Use(middleware.JWT())
		{
			loginRequired.PUT("/update", Update)
			loginRequired.GET("/read", Read)
			loginRequired.POST("/video", Create)
			loginRequired.DELETE("/deleteVideo", DeleteVideo)
			loginRequired.POST("/like/:id", Like)
			loginRequired.POST("/comment/:id", Comment)
			loginRequired.POST("/reply/:id", Reply)
			loginRequired.GET("/ReadComment", ReadComment)
			loginRequired.POST("/share/:id", Share)
			loginRequired.DELETE("/deleteComment/:id", DeleteComment)
			loginRequired.DELETE("/deleteReply/:id", DeleteReply)
			loginRequired.POST("/barrage/:id", CreateBarrage)
			loginRequired.DELETE("/deleteBarrage/:id", DeleteBarrage)
			loginRequired.PUT("/check/:id", Check)
			loginRequired.DELETE("/forbid/:id", Forbid)
			loginRequired.PUT("/unForbid/:id", UnForbid)
			loginRequired.DELETE("/manageComment/:id", ManageComment)
			loginRequired.POST("/collect/:id", Collect)
			loginRequired.DELETE("/deleteCollect/:id", DeleteCollect)
			loginRequired.GET("/readShare/:id", ReadShare)
			loginRequired.POST("/black/:id", Black)
			loginRequired.GET("/readBlack", ReadBlack)
			loginRequired.DELETE("/deleteBlack/:id", DeleteBlack)
		}
	}
	return r
}
