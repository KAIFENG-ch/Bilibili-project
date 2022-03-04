package controller

import (
	"Bilibili-project/dao"
	"Bilibili-project/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Comment(c *gin.Context) {
	var CommentData dao.CommentData
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	err := c.ShouldBind(&CommentData)
	id,_ := strconv.Atoi(c.Param("id"))
	res := CommentData.Comment(claims.Id, claims.Username, uint(id))
	if err != nil{
		c.JSON(400, err)
	}
	c.JSON(200, res)
}

func Reply(c *gin.Context) {
	var ReplyData dao.CommentData
	claims,_ := util.ParseToken(c.GetHeader("Authorization"))
	err := c.ShouldBind(&ReplyData)
	id,_ := strconv.Atoi(c.Param("id"))
	res := ReplyData.Reply(claims.Id, claims.Username, uint(id))
	if err != nil{
		c.JSON(400, err)
	}
	c.JSON(200, res)
}

func ReadComment(c *gin.Context) {
	res := dao.ReadComment(c.Param("id"), c.DefaultQuery("page", "1"))
	c.JSON(200, res)
}

func ReadReply(c *gin.Context) {
	res := dao.ReadReply(c.Param("id"), c.DefaultQuery("page", "1"))
	c.JSON(200, res)
}

func DeleteComment(c *gin.Context) {
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	res := dao.DeleteComment(claims.Id, c.Param("id"))
	c.JSON(200, res)
}

func DeleteReply(c *gin.Context) {
	claims,_ := util.ParseToken(c.GetHeader("Authorization"))
	res := dao.DeleteReply(claims.Id, c.Param("id"))
	c.JSON(200, res)
}