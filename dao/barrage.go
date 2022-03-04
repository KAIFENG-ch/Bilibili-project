package dao

import (
	"Bilibili-project/model"
	"Bilibili-project/serialize"
	"strconv"
)

type BarrageData struct {
	Content string `json:"content" form:"content"`
	Color string `json:"color" form:"color"`
	Time uint `json:"time" form:"time"`
}

func (receiver BarrageData) CreateBarrage(UserID uint, VideoID string) *serialize.Base {
	var BarrageInfo model.Barrage
	BarrageInfo = model.Barrage{
		UserID: UserID,
		VideoID: VideoID,
		Content: receiver.Content,
		Color: receiver.Color,
		Time: receiver.Time,
	}
	model.DB.Create(&BarrageInfo)
	return &serialize.Base{
		Status: 200,
		Msg: "ok",
		Data: "发送成功！",
	}
}

func ReadBarrage(page string) *serialize.Base {
	var BarrageInfo []model.Barrage
	pg,_ := strconv.Atoi(page)
	model.DB.Model(&model.Barrage{}).Limit(10).Offset((pg - 1) * 10).Find(&BarrageInfo)
	return serialize.BuildBarrages(BarrageInfo)
}

func DeleteBarrage(UserID uint, BarrageID string) *serialize.Base {
	var Barrage model.Barrage
	id,_ := strconv.Atoi(BarrageID)
	model.DB.Model(&model.Barrage{}).Where("user_id = ? and id = ?", UserID, id).
		Delete(&Barrage)
	return &serialize.Base{
		Status: 200,
		Msg: "ok",
		Data: "删除成功！",
	}
}