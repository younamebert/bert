package models

import (
	"bert/shopbert/database"
	"bert/shopbert/tools"
)

// user Model
type UserLogin struct {
	Id       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Tel      string `json:"tel"`
	Email    string `json:"email"`
	*database.Com
}

/*
	User这个Model的增删改查操作都放在这里
*/
// CreateAUser 创建user
func CreateAUser(user *UserLogin) (err error) {
	err = database.DB.Create(&user).Error
	return
}

func GetAllUser() (userList []*UserLogin, err error) {
	if err = database.DB.Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

func LoginUser(name, password string) (user *UserLogin, err error) {
	user = new(UserLogin)
	err = database.DB.Where("name = ? ", name).Or("tel = ?", name).Or("email = ?", name).First(user).Error
	if err != nil {
		return nil, err
	}
	user.UpdateAt = tools.Rday()
	database.DB.Save(&user)

	if password != user.Password {
		return nil, err
	}
	// // 日志插入
	// r := UserLog{
	// 	UserId:   user.Id,
	// 	UserText: "[1]登录成功",
	// 	// LogType:
	// }
	// Adduserlog(r)
	return
}

func GetAUser(id string) (user *UserLogin, err error) {
	user = new(UserLogin)
	if err = database.DB.Debug().Where("id=?", id).First(user).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateAUser(user *UserLogin) (err error) {
	err = database.DB.Save(user).Error
	return
}

func DeleteAUser(id string) (err error) {
	err = database.DB.Where("id=?", id).Delete(&UserLogin{}).Error
	return
}
