package controller

import (
	"Bilibili-project/dao"
	"Bilibili-project/util"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context)  {
	var registerMsg dao.UserRegister
	err := c.ShouldBind(&registerMsg)
	res := registerMsg.Register()
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, res)
}

func Login(c *gin.Context) {
	var loginMsg dao.UserLogin
	err := c.ShouldBind(&loginMsg)
	res := loginMsg.Login()
	if err != nil{
		c.JSON(400, err)
	}
	c.JSON(200,res)
}

func Update(c *gin.Context)  {
	var updateMsg dao.UserData
	claims, err := util.ParseToken(c.GetHeader("Authorization"))
	if err != nil{
		c.JSON(400, err)
	}
	_ = c.ShouldBind(&updateMsg)
	file,err := c.FormFile("headphoto")
	if err != nil{
		c.JSON(400, err)
	}
	res := updateMsg.Update(claims.Id, file)
	err = c.SaveUploadedFile(file, file.Filename)
	if err != nil{
		c.JSON(400, err)
	}
	c.JSON(200, res)
}

func Read(c *gin.Context) {
	claims, err := util.ParseToken(c.GetHeader("Authorization"))
	if err != nil{
		c.JSON(400, err)
	}
	res := dao.ReadUser(claims.Id)
	c.JSON(400, res)
}

func Black(c *gin.Context) {
	claims, err := util.ParseToken(c.GetHeader("Authorization"))
	if err != nil{
		c.JSON(400, err)
	}
	res := dao.Black(claims.Id, c.Param("id"))
	c.JSON(200, res)
}

func ReadBlack(c *gin.Context) {
	claims,err := util.ParseToken(c.GetHeader("Authorization"))
	if err != nil{
		c.JSON(400, err)
	}
	res := dao.ReadBlack(claims.Id)
	c.JSON(200, res)
}

func DeleteBlack(c *gin.Context) {
	claims, err := util.ParseToken(c.GetHeader("Authorization"))
	if err != nil{
		c.JSON(400, err)
	}
	res := dao.DeleteBlack(claims.Id, c.Param("id"))
	c.JSON(200, res)
}