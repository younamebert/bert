package models

import (
	"bert/shopbert/database"
)

type Buyrush struct {
	Id            int    `json:"id"`
	OriginalPrice string `json:"original_price"`
	RulingPrice   string `json:"ruling_price"`
	Goodsnum      int    `json:"num"`
	Goodsname     string `json:"goodsname"`
	Goodsattrbute string `json:"goodsattrbute"`
	Goodsimg      string `json:"goodsimg"`
	Starttime     string `json:"starttime"`
	*database.Com
}

func GetAllsBuyrush() (b []*Buyrush, err error) {
	// b = new(b)
	err = database.DB.Find(&b).Error
	if err != nil {
		return nil, err
	}
	return
}
