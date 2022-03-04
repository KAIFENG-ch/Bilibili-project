package model

var Admins = map[string]string{
	"admin1": "123456",
	"admin2": "7890ab",
}

type Admin struct {
	Username string
	Password string
}