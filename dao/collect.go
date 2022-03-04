package dao

import (
	"Bilibili-project/model"
	"Bilibili-project/serialize"
	"strconv"
)

func Collect(UserID uint, VideoID string) *serialize.Base {
	var UserInfo model.User
	var VideoInfo model.Video
	id,_ := strconv.Atoi(VideoID)
	model.DB.Model(&model.User{}).Where("id = ?", UserID).First(&UserInfo)
	model.DB.Model(&model.Video{}).Where("id = ?", id).First(&VideoInfo)
	err := model.DB.Model(&UserInfo).Association("Collect").Append(&VideoInfo)
	if err != nil{
		panic(err)
	}
	return &serialize.Base{
		Status: 200,
		Msg: "ok",
		Data: "收藏成功！",
	}
}

func DeleteCollect(UserID uint, VideoID string) *serialize.Base {
	var userInfo model.User
	var videoInfo model.Video
	id,_ := strconv.Atoi(VideoID)
	model.DB.Model(&model.User{}).Where("id = ?", UserID).First(&userInfo)
	model.DB.Model(&model.Video{}).Where("id = ?", id).First(&videoInfo)
	err := model.DB.Model(&userInfo).Association("Collect").Delete(&videoInfo)
	if err != nil {
		panic(err)
	}
	return &serialize.Base{
		Status: 200,
		Msg: "ok",
		Data: "删除成功！",
	}
}


