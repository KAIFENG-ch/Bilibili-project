package controller

import (
	"Bilibili-project/dao"
	"Bilibili-project/util"
	"github.com/gin-gonic/gin"
)

func Collect(c *gin.Context) {
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	res := dao.Collect(claims.Id, c.Param("id"))
	c.JSON(200, res)
}

func DeleteCollect(c *gin.Context) {
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	res := dao.DeleteCollect(claims.Id, c.Param("id"))
	c.JSON(200, res)
}
