package midd

import (
	"bert/shopbert/vender"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	Cache string = "1"
)

// 中间件设置api 请求
func Apicontext() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		/* new header methods
		post    eq add
		get     eq select
		options eq *()
		put     eq update
		delete  eq delete
		*/
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,PUT,DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		if c.Query("cache") == Cache {
			c.JSON(http.StatusOK, gin.H{
				"status": http.StatusOK,
				"msg":    "cache",
			})
			c.Abort()
		}

		// 用户是否登录
		usertoken := c.Query("usertoken")
		query := c.Query("query")
		str, _ := vender.RedisGet(usertoken)

		if str != "200" && query != "login" {
			c.JSON(http.StatusOK, gin.H{
				"status": http.StatusOK,
				"msg":    "not usertoken failed err",
			})
			c.Abort()
		}

		// pages := c.Query("pages")
		// // api 请求限制
		// if usertoken != "" && pages != "" {
		// 	// alls := usertoken + pages
		// 	res := vender.Userlimit()

		// 	if res != "1" {
		// 		c.JSON(http.StatusOK, gin.H{
		// 			"status": http.StatusOK,
		// 			"msg":    "请求频繁",
		// 			"error":  res,
		// 		})
		// 		c.Abort()
		// 	}
		// }
		c.Next()
	}
}
