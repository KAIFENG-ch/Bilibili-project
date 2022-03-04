package dao

import (
	"Bilibili-project/model"
	"Bilibili-project/serialize"
	"strconv"
)

type CommentData struct {
	Content string `json:"content" form:"content"`
}

func (receiver CommentData) Comment(UserID uint, Username string, VideoID uint) *serialize.Base {
	var CommentData model.Comments
	//model.DB.Model(model.User{}).Where("id = ?", id).First(&user)
	CommentData = model.Comments{
		VideoID: VideoID,
		UserID: UserID,
		Username: Username,
		Content: receiver.Content,
	}
	model.DB.Create(&CommentData)
	return &serialize.Base{
		Status: 200,
		Msg: "ok",
		Data: "评论成功！",
	}
}

func (receiver CommentData) Reply (UserID uint, Username string, CommentID uint) *serialize.Base {
	var ReplyData model.Replies
	ReplyData = model.Replies{
		UserID: UserID,
		CommentsID: CommentID,
		Username: Username,
		Content: receiver.Content,
	}
	model.DB.Create(&ReplyData)
	return &serialize.Base{
		Status: 200,
		Msg: "ok",
		Data: "评论成功！",
	}
}

func ReadComment(VideoID string, page string) *serialize.Base {
	var ReadData []model.Comments
	pg,_ := strconv.Atoi(page)
	model.DB.Model(&model.Comments{}).Where("video_id = ?", VideoID).
		Limit(10).Offset((pg - 1) * 10).Find(&ReadData)
	return serialize.BuildComments(ReadData)
}

func ReadReply(CommentID string, page string) *serialize.Base {
	var ReadReply []model.Replies
	pg,_ := strconv.Atoi(page)
	model.DB.Model(&model.Replies{}).Where("comments_id = ?", CommentID).
		Limit(10).Offset((pg - 1) * 10).Find(&ReadReply)
	return serialize.BuildReplies(ReadReply)
}

func DeleteComment(UserID uint, CommentID string) *serialize.Base {
	var DeleteComment model.Comments
	model.DB.Model(&model.Comments{}).Where("user_id = ? and id = ?", UserID, CommentID).
		Delete(&DeleteComment)
	return &serialize.Base{
		Status: 200,
		Msg: "ok",
		Data: "删除成功！",
	}
}

func DeleteReply(UserID uint, ReplyID string) *serialize.Base {
	var DeleteReply model.Replies
	model.DB.Model(&model.Replies{}).Where("user_id = ? and id = ?", UserID, ReplyID).
		Delete(&DeleteReply)
	return &serialize.Base{
		Status: 200,
		Msg: "ok",
		Data: "删除成功！",
	}
}