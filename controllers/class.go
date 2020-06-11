package controller

import (
	"bert/shopbert/models"
	"bert/shopbert/tools"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Obj struct {
	SeriesId string `json:"series_id"`
	Img      string `json:"img"`
	Name     string `json:"name"`
}

type BrandAlls struct {
	BrandName string `json:"brand_name"`
	BrandId   int    `json:"brand_id"`
	ObjAlls   *[]Obj
}

func GetClassHander(c *gin.Context) {
	// 接收数据
	res := c.Query("token")
	// 解密
	resa := []byte(tools.DePwdCode(res))

	query := &tools.Tabname{}

	json.Unmarshal(resa, &query)

	if query.Query != "class" {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
			"err":    "failed query not eq getclass",
		})
		c.Abort()
	}

	// 获取商品数据
	data, err := models.ClassAlls()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"err":    err,
		})
		c.Abort()
	}

	brand, err := models.GetProductClassBrandAlls()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"err":    err,
		})
		c.Abort()
	}

	series, err := models.GetProductClassBrandSeries()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"err":    err,
		})
		c.Abort()
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   data,
		"brand":  brand,
		"series": series,
	})

}
