package main

import (
	"bert/shopbert/database"
	"bert/shopbert/routers"
)

func main() {
	// fmt.Print(vender.Userlimit("aaa"))
	// 初始化mysql
	err := database.InitMySQL()
	if err != nil {
		panic(err)
	}
	// 初始化redis
	err = database.InitRedis()
	if err != nil {
		panic(err)
	}

	defer database.Close()
	// 注册路由

	r := routers.SetupRouter()
	r.Run()
}
