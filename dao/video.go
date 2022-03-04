package dao

import (
	"Bilibili-project/model"
	"Bilibili-project/serialize"
	"mime/multipart"
	"strconv"
	"time"
)

type VideoData struct {
	Title string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}

func (receiver VideoData) Create(ID uint, VideoFile *multipart.FileHeader) *serialize.Base {
	var VideoData model.Video
	var UserData model.User
	model.DB.Model(&model.User{}).Where("id = ?", ID).First(&UserData)
	model.RDB.Set(receiver.Title + " scanned", 0, 100000 * time.Hour)
	VideoData = model.Video{
		Title: receiver.Title,
		Content: receiver.Content,
		VideoFile: VideoFile.Filename,
		UserID: ID,
		Username: UserData.Name,
	}
	model.DB.Create(&VideoData)
	return &serialize.Base{
		Status: 200,
		Msg: "ok",
		Data: "上传成功！",
	}
}

func ReadOneVideo(videoID string) *serialize.Base {
	id,_ := strconv.Atoi(videoID)
	var video model.Video
	model.DB.Model(&model.Video{}).Where("id = ?", id).First(&video)
	model.RDB.Incr(video.Title + " scanned")
	str,_ := model.RDB.Get(video.Title + " scanned").Result()
	click,_ := strconv.Atoi(str)
	model.DB.Model(&video).Update("click", click)
	return &serialize.Base{
		Status: 200,
		Msg: "OK",
		Data: serialize.Video{
			Title: video.Title,
			Content: video.Content,
			Video: video.VideoFile,
			Author: video.Username,
			Click: uint(click),
		},
	}
}

func ReadVideo(userID string, page string) *serialize.Base {
	var videoInfo []model.Video
	pg,_ := strconv.Atoi(page)
	model.DB.Model(&model.Video{}).Where("user_id = ?", userID).
		Limit(10).Offset((pg - 1) * 10).Find(&videoInfo)
	return serialize.BuildVideos(videoInfo)
}

func DeleteVideo(UserID uint, vid string) *serialize.Base {
	var deleteVideo model.Video
	VideoID,_ := strconv.Atoi(vid)
	model.DB.Model(&model.Video{}).Where("user_id = ? and id = ?", UserID, VideoID).
		Delete(&deleteVideo)
	return &serialize.Base{
		Status: 200,
		Msg: "ok",
		Data: "删除成功！",
	}
}

func Like(UserID uint, VideoID string) *serialize.Base {
	id,_ := strconv.Atoi(VideoID)
	value, err := model.RDB.SIsMember("video" + VideoID, UserID).Result()
	if err != nil{
		panic(err)
	}
	if value == false{
		_, err := model.RDB.SAdd("video" + VideoID, UserID).Result()
		model.DB.Model(model.Video{}).Where("id = ?", id).
			Update("like", model.RDB.SCard("video" + VideoID).Val())
		if err != nil{
			panic(err)
		}
		return &serialize.Base{
			Status: 200,
			Msg: "ok",
			Data: "点赞成功！",
		}
	} else {
		_, err := model.RDB.SRem("video" + VideoID, UserID).Result()
		model.DB.Model(model.Video{}).Where("id = ?", id).
			Update("like", model.RDB.SCard("video" + VideoID).Val())
		if err != nil{
			panic(err)
		}
		return &serialize.Base{
			Status: 200,
			Msg: "ok",
			Data: "取消成功！",
		}
	}
}

