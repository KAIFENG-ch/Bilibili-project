package dao

import (
	"Bilibili-project/model"
	"Bilibili-project/serialize"
)

func ReadClickList() *serialize.Base {
	var videos []model.Video
	model.DB.Order("click desc").Find(&videos)
	return &serialize.Base{
		Status: 200,
		Msg: "ok",
		Data: videos,
	}
}
