package dao

import (
	"Bilibili-project/model"
	"Bilibili-project/serialize"
	"Bilibili-project/util"
	"strconv"
)

type AdminLogin struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func (receiver AdminLogin) AdminLogin() *serialize.Base {
	loginInfo := model.Admin{
		Username: receiver.Username,
		Password: receiver.Password,
	}
	if model.Admins[loginInfo.Username] == loginInfo.Password{
		token,err := util.CreateAdminToken(loginInfo)
		if err != nil{
			panic(err)
		}
		return &serialize.Base{
			Status: 200,
			Msg: "ok",
			Data: token,
		}
	}
	return &serialize.Base{
		Status: 400,
		Msg: "fail",
		Data: "登陆失败！",
	}
}

func Check(VideoID string) *serialize.Base {
	id,_ := strconv.Atoi(VideoID)
	model.DB.Model(&model.Video{}).Where("id = ?", id).
		Update("status", true)
	return &serialize.Base{
		Status: 200,
		Msg: "ok",
		Data: "审核通过",
	}
}

func Forbid(UserID string) *serialize.Base {
	var user model.User
	id,_ := strconv.Atoi(UserID)
	//r := errors.Is(model.DB.Model(&model.User{}).Where("id = ?", id).First(&user).Error, gorm.ErrRecordNotFound)
	model.DB.Model(&model.User{}).Where("id = ?", id).First(&user)
	model.DB.Where("id = ?", UserID).Delete(&model.User{})
	return &serialize.Base{
		Status: 200,
		Msg: "ok",
		Data: "封号成功！",
	}
}

func UnForbid(UserID string) *serialize.Base {
	id,_ := strconv.Atoi(UserID)
	model.DB.Model(&model.User{}).Unscoped().Where("id = ?", id).Update("deleted_at", nil)
	return &serialize.Base{
		Status: 200,
		Msg: "ok",
		Data: "解封成功！",
	}
}

func ManageComment(CommentID string) *serialize.Base {
	id,_ := strconv.Atoi(CommentID)
	model.DB.Where("id = ?",id).Delete(&model.Comments{})
	return &serialize.Base{
		Status: 200,
		Msg: "ok",
		Data: "删除成功！",
	}
}