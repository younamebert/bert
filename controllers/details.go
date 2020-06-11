package controller

import (
	"bert/shopbert/models"
	"bert/shopbert/tools"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Details struct {
	*tools.Tabname
	Productid int64 `json:"productid"`
}

func GetDetailsHander(c *gin.Context) {

	request := c.Query("token")
	requests := []byte(tools.DePwdCode(request))

	a := &Details{}
	json.Unmarshal(requests, &a)

	r := a.Productid
	info, err := models.WhereGetAlls(r)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
			"error":  err,
		})
		c.Abort()
	}

	infoversion, err := models.WhereGetAllsVersion(info.ProductId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
			"error":  err,
		})
	}
	// infojson, err := json.Unmarshal(&info)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"data":    info,
		"version": infoversion,
	})
}
