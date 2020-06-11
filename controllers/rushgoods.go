package controller

import (
	"bert/shopbert/models"
	"bert/shopbert/tools"
	"bert/shopbert/vender"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// websocket
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 消息中间件
func midd() {

}

// 秒杀开始 redis减少
func redisbuy() {

}

// 数据写入
func addOrder() {

}

// 商品数据 分发流量
func GetRushgoodsHander(c *gin.Context) {

	if vender.RedisExists("rushgoods") == "1" {

		data, _ := vender.RedisGet("rushgoods")
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"msg":    "已写入缓存",
			"data":   data,
		})
		return
	}

	data, err := models.GetAllsBuyrush()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"error":  err,
		})
	}

	// vender.RedisLPush("")
	for i := 0; i < len(data); i++ {
		// vender.RedisLPush("")
		for n := 0; n < data[i].Goodsnum; n++ {
			strkey := strconv.Itoa(data[i].Id)
			vender.RedisLPush(strkey, n)
		}
	}

	a, _ := json.Marshal(data)

	vender.RedisSet("rushgoods", tools.EnPwdCode(string(a[:])))

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   tools.EnPwdCode(string(a[:])),
	})

}
