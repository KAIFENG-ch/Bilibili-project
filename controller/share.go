package controller

import (
	"Bilibili-project/dao"
	"Bilibili-project/util"
	"github.com/gin-gonic/gin"
)

func Share(c *gin.Context) {
	claims, err := util.ParseToken(c.GetHeader("Authorization"))
	if err != nil{
		panic(err)
	}
	res := dao.Share(claims.Username, c.Param("id"))
	c.JSON(200, res)
}

func ReadShare(c *gin.Context) {
	res := dao.ReadShare(c.Param("id"))
	c.JSON(200, res)
}
