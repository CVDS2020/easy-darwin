package models

import (
	"github.com/suy/easy-darwin/db"
	"github.com/suy/easy-darwin/utils"
)

func Init() (err error) {
	err = db.Init()
	if err != nil {
		return
	}
	db.Mysql.AutoMigrate(User{}, Stream{})
	count := 0
	sec := utils.Conf().Section("http")
	defUser := sec.Key("default_username").MustString("admin")
	defPass := sec.Key("default_password").MustString("admin")
	db.Mysql.Model(User{}).Where("username = ?", defUser).Count(&count)
	if count == 0 {
		db.Mysql.Create(&User{
			Username: defUser,
			Password: utils.MD5(defPass),
		})
	}
	return
}

func Close() {
	db.Close()
}
