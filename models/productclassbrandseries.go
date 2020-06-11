package models

import "bert/shopbert/database"

type ProductClassBrandSeries struct {
	SeriesId   int    `json:"series_id"`
	BrandValue string `json:"brand_value"`
	BrandId    int    `json:"brand_id"`
	SeriesName string `json:"series_name"`
	SeriesImg  string `json:"series_img"`
	ClassId    int    `json:"class_id"`
	*database.Com
}

func GetProductClassBrandSeries() (s []*ProductClassBrandSeries, err error) {
	err = database.DB.Find(&s).Error
	if err != nil {
		return nil, err
	}
	return
}
