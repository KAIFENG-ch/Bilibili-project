package controller

import (
	"Bilibili-project/dao"
	"Bilibili-project/util"
	"github.com/gin-gonic/gin"
)

func CreateBarrage(c *gin.Context) {
	var Barrage dao.BarrageData
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	err := c.ShouldBind(&Barrage)
	if err != nil{
		c.JSON(400, err)
	}
	res := Barrage.CreateBarrage(claims.Id, c.Param("id"))
	c.JSON(200, res)
}

func ReadBarrage(c *gin.Context) {
	res := dao.ReadBarrage(c.DefaultQuery("page", "1"))
	c.JSON(200, res)
}

func DeleteBarrage(c *gin.Context)  {
	claims,_ := util.ParseToken(c.GetHeader("Authorization"))
	res := dao.DeleteBarrage(claims.Id, c.Param("id"))
	c.JSON(200, res)
}
