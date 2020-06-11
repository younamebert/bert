package models

import "bert/shopbert/database"

type ProductDetails struct {
	DetailsId         int64  `json:"details_id"`
	ProductId         int64  `json:"product_id"`
	DetailsImg        string `json:"details_img"`
	DetailsBottomdata string `json:"details_bottomdata"`
	DetailsJsondata   string `json:"details_jsondata"`
	*database.Com
	DetailsBottomimg string  `json:"details_bottomimg"`
	Price            float64 `json:"price"`
	ProductName      string  `json:"name"`
}

func WhereGetAlls(a int64) (s *ProductDetails, err error) {
	s = new(ProductDetails)
	err = database.DB.Where("details_id = ?", a).First(&s).Error
	if err != nil {
		return nil, err
	}
	return
}
