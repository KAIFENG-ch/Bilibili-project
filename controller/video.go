package controller

import (
	"Bilibili-project/dao"
	"Bilibili-project/util"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var VideoData dao.VideoData
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	err := c.ShouldBind(&VideoData)
	if err != nil{
		c.JSON(400, err)
	}
	file, err := c.FormFile("videofile")
	if err != nil{
		c.JSON(400, err)
	}
	res := VideoData.Create(claims.Id, file)
	err = c.SaveUploadedFile(file, file.Filename)
	if err != nil{
		c.JSON(400, err)
	}
	c.JSON(200, res)
}

func ReadOneVideo(c *gin.Context) {
	res := dao.ReadOneVideo(c.Param("id"))
	c.JSON(200, res)
}

func ReadVideo(c *gin.Context) {
	res := dao.ReadVideo(c.Param("id"), c.DefaultQuery("page", "1"))
	c.JSON(200, res)
}

func DeleteVideo(c *gin.Context) {
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	res := dao.DeleteVideo(claims.Id, c.Param("id"))
	c.JSON(200, res)
}

func Like(c *gin.Context) {
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	res := dao.Like(claims.Id, c.Param("id"))
	c.JSON(200, res)
}
