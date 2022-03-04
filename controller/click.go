package controller

import (
	"Bilibili-project/dao"
	"github.com/gin-gonic/gin"
)

func ReadClickList(c *gin.Context) {
	res := dao.ReadClickList()
	c.JSON(200, res)
}
