package dao

import (
	"Bilibili-project/model"
	"Bilibili-project/serialize"
	"strconv"
)

func Share(username string, VideoID string) *serialize.Base {
	id,_ := strconv.Atoi(VideoID)
	var video model.Video
	model.DB.Model(&model.Video{}).Where("id = ?", id).First(&video)
	model.RDB.LPush(video.Title, username)
	return &serialize.Base{
		Status: 200,
		Msg: "ok",
		Data: "分享成功！",
	}
}

func ReadShare(VideoID string) *serialize.Base {
	id,_ := strconv.Atoi(VideoID)
	var video model.Video
	model.DB.Model(&model.Video{}).Where("id = ?", id).First(&video)
	shareInfo := model.RDB.LRange(video.Title, 0, -1)
	return &serialize.Base{
		Status: 200,
		Msg: "ok",
		Data: serialize.Datalist{
			Item: shareInfo.Val(),
			Total: len(shareInfo.Val()),
		},
	}
}
