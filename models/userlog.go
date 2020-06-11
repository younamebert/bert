package models

import (
	"bert/shopbert/database"
)

type UserLog struct {
	UserlogId int    `json:"userlog_id"`
	UserId    int    `json:"user_id"`
	UserText  string `json:"user_text"`
	LogType   int    `json:"log_type"`
	*database.Com
}

func Getuserlogalls() (ulog []*UserLog, err error) {

	err = database.DB.Find(&ulog).Error
	if err != nil {
		return nil, err
	}
	return
}

func Adduserlog(add UserLog) {
	database.DB.Create(add)

}
