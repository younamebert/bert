package models

import "bert/shopbert/database"

type UserDetail struct {
	Id           int    `json:"id"`
	Tel          string `json:"tel"`
	Identity     string `json:"identity"`
	Email        string `json:"email"`
	Birthdaytime string `json:"birthdaytime"`
	Gender       int    `json:"gender"`
	Topimg       string `json:"topimg"`
	*database.Com
	UserId   int    `json:"user_id"`
	Zip      string `json:"zip"`
	Nickname string `json:"nickname"`
}

func ListOne(id string) (u *UserDetail, err error) {
	u = new(UserDetail)
	err = database.DB.Where("tel = ?", id).First(u).Error
	if err != nil {
		panic(err)
	}
	return
}
