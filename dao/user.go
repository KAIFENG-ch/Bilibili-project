package dao

import (
	"Bilibili-project/model"
	"Bilibili-project/serialize"
	"Bilibili-project/util"
	"golang.org/x/crypto/bcrypt"
	"mime/multipart"
	"strconv"
	"time"
)

type UserRegister struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type UserLogin struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type UserData struct {
	Username string `json:"username" form:"username"`
	Gender string `json:"gender" form:"gender"`
	Email string `json:"email" form:"email"`
	Age uint `json:"age" form:"age"`
	Birthday int64 `json:"birthday" form:"birthday"`
	Signature string `json:"signature" form:"signature"`
}

type UserHeadPhoto struct {
	HeadPhoto string `json:"head_photo" file:"head_photo"`
}

type ChangePwd struct {
	OldPwd string `json:"old_pwd" form:"old_pwd"`
	NewPwd string `json:"new_pwd" form:"new_pwd"`
	NewPwdAgain string `json:"new_pwd_again" form:"new_pwd_again"`
}

type ReadUsers struct {
	ID uint
	Name string
	Age uint
	Gender string
	Email string
	Birthday uint
	Signature string
	CreatedAt time.Time
}

func (receiver UserRegister) Register() *serialize.Base {
	var user model.User
	var count int64
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(receiver.Password),bcrypt.DefaultCost)
	user = model.User{
		Name:     receiver.Username,
		Password: string(hashPassword),
	}
	model.DB.Where("name = ?", receiver.Username).First(&model.User{}).Count(&count)
	if count > 0{
		return &serialize.Base{
			Status: 400,
			Msg: "fail",
			Data: "用户已存在",
		}
	}
	model.DB.Create(&user)
	return &serialize.Base{
		Status: 200,
		Msg: "ok",
		Data: user.Name,
	}
}

func (receiver UserLogin) Login() *serialize.Base {
	var userMsg model.User
	model.DB.Model(&model.User{}).Where("name = ?", receiver.Username).First(&userMsg)
	if userMsg.ID == 0{
		return &serialize.Base{
			Status: 400,
			Msg: "fail",
			Data: "用户不存在",
		}
	}
	result := bcrypt.CompareHashAndPassword([]byte(userMsg.Password), []byte(receiver.Password))
	if result != nil{
		return &serialize.Base{
			Status: 400,
			Msg: "fail",
			Data: "密码错误！",
		}
	}
	token, err := util.CreateToken(userMsg)
	if err != nil {
		panic(err)
	}
	return &serialize.Base{
		Status: 200,
		Msg: "ok",
		Data: token,
	}
}

func (receiver UserData) Update(ID uint, HeadFile *multipart.FileHeader) *serialize.Base {
	model.DB.Model(&model.User{}).Where("id = ?", ID).
		Updates(map[string]interface{}{"name": receiver.Username, "age": receiver.Age,
			"head_photo": HeadFile.Filename, "gender": receiver.Gender, "email": receiver.Email,
			"birthday": receiver.Birthday, "signature": receiver.Signature})
	return &serialize.Base{
		Status: 200,
		Msg: "ok",
		Data: "update successful",
	}
}

func (receiver ChangePwd) ChangePwd(ID uint) *serialize.Base {
	var user model.User
	model.DB.Model(&model.User{}).Where("id = ?", ID).First(&user)
	result := bcrypt.CompareHashAndPassword([]byte(receiver.OldPwd), []byte(user.Password))
	if result != nil{
		return &serialize.Base{
			Status: 400,
			Msg: "fail",
			Data: "密码错误！",
		}
	}
	if receiver.NewPwd != receiver.NewPwdAgain {
		return &serialize.Base{
			Status: 400,
			Msg: "fail",
			Data: "两次密码不一致！",
		}
	}
	HashedNewPwd, err := bcrypt.GenerateFromPassword([]byte(receiver.OldPwd), bcrypt.DefaultCost)
	if err != nil{
		panic(err)
	}
	model.DB.Model(&model.User{}).Where("id = ?", ID).Update("password", HashedNewPwd)
	return &serialize.Base{
		Status: 200,
		Msg: "ok",
		Data: "修改成功！",
	}
}

func ReadUser(ID uint) *serialize.User {
	var userMsg model.User
	model.DB.Where("id = ?", ID).First(&userMsg)
	return &serialize.User{
		ID:  userMsg.ID,
		Username: userMsg.Name,
		Email: userMsg.Email,
		Age: userMsg.Age,
		Birthday: userMsg.Birthday,
		Signature: userMsg.Signature,
		CreatedAt: userMsg.CreatedAt.Unix(),
	}
}


func Black(UserID uint, BlackID string) *serialize.Base {
	var blackUser model.User
	var user model.User
	id,_ := strconv.Atoi(BlackID)
	model.DB.Model(&model.User{}).Where("id = ?", id).First(&blackUser)
	model.DB.Model(&model.User{}).Where("id = ?", UserID).First(&user)
	err := model.DB.Model(&user).Association("Black").Append(&blackUser)
	if err != nil {
		panic(err)
	}
	return &serialize.Base{
		Status: 200,
		Msg: "ok",
		Data: "拉黑成功！",
	}
}

func ReadBlack(UserID uint) *serialize.Base {
	var users []model.User
	model.DB.Model(&model.User{}).Where("id = ?", UserID).
		Preload("Black").Find(&users)
	return &serialize.Base{
		Status: 200,
		Msg: "ok",
		Data: serialize.Datalist{
			Item: serialize.BuildUsers(users),
			Total: len(users),
		},
	}
}

func DeleteBlack(UserID uint, BlackID string) *serialize.Base {
	var deleteUser model.User
	var user model.User
	id,_ := strconv.Atoi(BlackID)
	model.DB.Model(&model.User{}).Where("id = ?", id).First(&deleteUser)
	model.DB.Model(&model.User{}).Where("id = ?", UserID).First(&user)
	err := model.DB.Model(&user).Association("Black").Delete(&deleteUser)
	if err != nil {
		panic(err)
	}
	return &serialize.Base{
		Status: 200,
		Msg: "ok",
		Data: "已删除",
	}
}