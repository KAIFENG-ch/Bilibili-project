package controller

import (
	"Bilibili-project/dao"
	"github.com/gin-gonic/gin"
)

func AdminLogin(c *gin.Context) {
	var loginInfo dao.AdminLogin
	err := c.ShouldBind(&loginInfo)
	if err != nil{
		c.JSON(400, err)
	}
	res := loginInfo.AdminLogin()
	c.JSON(200, res)
}

func Check(c *gin.Context) {
	res := dao.Check(c.Param("id"))
	c.JSON(200, res)
}

func Forbid(c *gin.Context) {
	res := dao.Forbid(c.Param("id"))
	c.JSON(200, res)
}

func UnForbid(c *gin.Context) {
	res := dao.UnForbid(c.Param("id"))
	c.JSON(200, res)
}

func ManageComment(c *gin.Context) {
	res := dao.ManageComment(c.Param("id"))
	c.JSON(200, res)
}
