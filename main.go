package main

import (
	"Bilibili-project/conf"
	"Bilibili-project/controller"
)

func main()  {
	conf.InitConfig()
	r := controller.RouterInit()
	_ = r.Run(":8000")
}
