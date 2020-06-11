package controller

import (
	"bert/shopbert/models"
	"bert/shopbert/tools"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPersonageHander(c *gin.Context) {
	// msg :=
	info := tools.DePwdCode(c.Query("usertoken"))

	data, err := models.ListOne(info)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"error":  err,
		})
		c.Abort()
	}

	f := make(map[string]interface{})

	f["name"] = data.Nickname
	f["gender"] = data.Gender
	f["topimg"] = data.Topimg
	f["email"] = data.Email
	f["zip"] = data.Zip
	f["age"] = data.Birthdaytime
	f["usertoken"] = c.Query("usertoken")

	m, err := json.Marshal(f)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"error":  err,
		})
		c.Abort()
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   tools.EnPwdCode(string(m[:])),
	})
}

// put user details 2020/6/5
func PutPersonageHander(c *gin.Context) {
	// query := c.Query("")
	// c.Query("")
}
