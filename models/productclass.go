package models

import "bert/shopbert/database"

type ProductClass struct {
	ClassId   int    `json:"class_id"`
	ClassName string `json:"class_name"`
	ClassCode string `json:"class_code"`
	FatherId  int    `json:"father_id"`
	IconImg   string `json:"icon_img"`
	*database.Com
}

func ClassAlls() (p []*ProductClass, err error) {
	err = database.DB.Find(&p).Error
	if err != nil {
		return nil, err
	}
	return
}
