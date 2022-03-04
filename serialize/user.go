package serialize

import "Bilibili-project/model"

type User struct {
	ID  uint `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Email string `json:"email" form:"email"`
	Age uint `json:"age" form:"age"`
	Birthday int64 `json:"birthday"`
	Signature string `json:"signature" form:"signature"`
	CreatedAt int64 `json:"created_at"`
}

func BuildUser(user model.User) User  {
	return User{
		Username: user.Name,
		Email: user.Email,
		Age: user.Age,
		Birthday: user.Birthday,
		Signature: user.Signature,
		CreatedAt: user.CreatedAt.Unix(),
	}
}

func BuildUsers(data []model.User) (Users []User) {
	for _, user := range data{
		user := BuildUser(user)
		Users = append(Users, user)
	}
	return Users
}
