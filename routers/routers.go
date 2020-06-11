package routers

import (
	controller "bert/shopbert/controllers"
	"bert/shopbert/midd"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "static")
	r.Use(midd.Apicontext())

	//r.GET("/cache", controller.CacheHander)

	r.GET("/login", controller.LoginHander)

	r.GET("/class", controller.GetClassHander)

	r.GET("/details", controller.GetDetailsHander)

	/*
		personage 用户
	*/
	r.GET("/personage", controller.GetPersonageHander)
	r.PUT("/personage", controller.PutPersonageHander)

	/*

	 */
	r.GET("/rushgoods", controller.GetRushgoodsHander)
	return r
}
