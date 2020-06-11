package models

import "bert/shopbert/database"

type ProductClassBrand struct {
	BrandId    int    `json:"brand_id"`
	BrandValue string `json:"brand_value"`
	BrandLogo  string `json:"brand_logo"`
	BrandTel   string `json:"brand_tel"`
	BrandDesc  string `json:"brand_desc"`
	BrandWeb   string `json:"brand_web"`
	ClassId    int    `json:"class_id"`
	*database.Com
}

func GetProductClassBrandAlls() (b []*ProductClassBrand, err error) {
	err = database.DB.Find(&b).Error
	if err != nil {
		return nil, err
	}
	return
}
