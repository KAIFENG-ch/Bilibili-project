package dao

import (
	"Bilibili-project/model"
	"Bilibili-project/serialize"
	"strconv"
)

type Search struct {
	Content string `json:"content" form:"content"`
}

type Topic struct {
	CreatedAt uint `json:"created_at" form:"created_at"`
	UpdatedAt uint `json:"updated_at" form:"updated_at"`
	VideoID uint `json:"video_id" form:"video_id"`
	Username string `json:"username" form:"username"`
}

func (receiver Search) SearchVideo(page string) *serialize.Base {
	var videos []model.Video
	pg,_ := strconv.Atoi(page)
	model.DB.Model(&model.Video{}).Limit(10).Offset((pg - 1) * 10).
		Where("title like ? or content like ? or username like ? ",
			"%" + receiver.Content + "%", "%" + receiver.Content + "%",
			"%" + receiver.Content + "%").
		Find(&videos)
	model.RDB.LPush("search", receiver.Content)
	return &serialize.Base{
		Status: 200,
		Msg: "ok",
		Data: serialize.Datalist{
			Item: videos,
			Total: len(videos),
		},
	}
}

func (receiver Search) SearchUser(page string) *serialize.Base {
	var users []ReadUsers
	pg,_ := strconv.Atoi(page)
	model.DB.Model(&model.User{}).Limit(10).Offset((pg - 1) * 10).
		Where("name like ? or signature like ? or gender = ? or age = ? or birthday = ?",
			"%" + receiver.Content + "%", "%" + receiver.Content + "%",
			receiver.Content, receiver.Content, receiver.Content).
		Find(&users)
	model.RDB.LPush("search", receiver.Content)
	return &serialize.Base{
		Status: 200,
		Msg: "ok",
		Data: serialize.Datalist{
			Item: users,
			Total: len(users),
		},
	}
}

func HistorySearch() *serialize.Base {
	return &serialize.Base{
		Status: 200,
		Msg: "ok",
		Data: serialize.Datalist{
			Item: model.RDB.LRange("search", 0, 10).Val(),
			Total: 10,
		},
	}
}

func (receiver Topic) GetTopic(page string) *serialize.Base {
	result := model.DB.Model(&model.Video{})
	pg,_ := strconv.Atoi(page)
	var searchResult []model.Video
	if receiver.VideoID > 0{
		result = result.Where("id = ?", receiver.VideoID)
	}
	if receiver.CreatedAt > 0{
		result = result.Order("created_at desc")
	}
	if receiver.UpdatedAt > 0{
		result = result.Order("updated_at desc")
	} else if receiver.Username != ""{
		result = result.Where("username = ?", receiver.Username)
	}
	result = result.Limit(10).Offset((pg - 1) * 10).Find(&searchResult)
	return &serialize.Base{
		Status: 200,
		Msg: "ok",
		Data: serialize.Datalist{
			Item: searchResult,
			Total: len(searchResult),
		},
	}
}