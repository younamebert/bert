package models

import "bert/shopbert/database"

type ProductDetailsVersion struct {
	VersionId  int64  `json:"version_id"`
	DetailsId  int64  `json:"details_id"`
	ProductId  int64  `json:"product_id"`
	ProductImg string `json:"productImg"`
	*database.Com
	Num   int     `json:"num"`
	Price float64 `json:"price"`
	Attr  string  `json:"attr"`
}

func WhereGetAllsVersion(a int64) (u []*ProductDetailsVersion, err error) {

	err = database.DB.Where("product_id = ?", a).Find(&u).Error
	if err != nil {
		return nil, err
	}
	return

}
