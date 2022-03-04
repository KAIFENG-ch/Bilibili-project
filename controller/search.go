package controller

import (
	"Bilibili-project/dao"
	"github.com/gin-gonic/gin"
)

func SearchVideo(c *gin.Context) {
	var search dao.Search
	err := c.ShouldBind(&search)
	if err != nil{
		c.JSON(400, err)
	}
	res := search.SearchVideo(c.DefaultQuery("page", "1"))
	c.JSON(200, res)
}

func SearchUser(c *gin.Context) {
	var search dao.Search
	err := c.ShouldBind(&search)
	if err != nil{
		c.JSON(400, err)
	}
	res := search.SearchUser(c.DefaultQuery("page", "1"))
	c.JSON(200, res)
}

func SearchHistory(c *gin.Context) {
	res := dao.HistorySearch()
	c.JSON(200, res)
}

func GetTopic(c *gin.Context) {
	var topic dao.Topic
	err := c.ShouldBind(&topic)
	if err != nil{
		c.JSON(400, err)
	}
	res := topic.GetTopic(c.Param("page"))
	c.JSON(200, res)
}