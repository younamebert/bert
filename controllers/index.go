package controller

import (
	"bert/shopbert/models"
	"bert/shopbert/tools"
	"bert/shopbert/vender"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Name     string `form:"user" json:"name"`
	Password string `form:"password" json:"password"`
}

func LoginHander(c *gin.Context) {

	msg := c.Query("token")

	// json
	loginJson := []byte(tools.DePwdCode(msg))

	login := &Login{}
	json.Unmarshal(loginJson, &login)

	data, err := models.LoginUser(login.Name, login.Password)

	if err != nil {
		panic(err)
	}

	if data.Name != login.Name {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
		})
	}

	// var b bytes.Buffer
	// gz := gzip.NewWriter(&b)
	// if _, err := gz.Write([]byte("YourDataHere")); err != nil {
	// 	panic(err)
	// }
	// if err := gz.Flush(); err != nil {
	// 	panic(err)
	// }
	// if err := gz.Close(); err != nil {
	// 	panic(err)
	// }
	// fmt.Println(b)

	vender.RedisSet(tools.EnPwdCode(string(data.Tel)), "200")

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   tools.EnPwdCode(string(data.Tel)),
	})
}
