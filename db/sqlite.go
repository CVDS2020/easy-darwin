package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/suy/easy-darwin/utils"
)

type Model struct {
	ID        string         `structs:"id" gorm:"primary_key" form:"id" json:"id"`
	CreatedAt utils.DateTime `structs:"-" json:"createdAt" gorm:"type:datetime"`
	UpdatedAt utils.DateTime `structs:"-" json:"updatedAt" gorm:"type:datetime"`
	// DeletedAt *time.Time `sql:"index" structs:"-"`
}

var Mysql *gorm.DB

func Init() (err error) {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTablename string) string {
		return "t_" + defaultTablename
	}
	dbDsn := utils.Conf().Section("db").Key("dsn").MustString("root:css66018@(localhost)/easydarwin?charset=utf8&parseTime=True&loc=Local")
	Mysql, err = gorm.Open("mysql", dbDsn)
	if err != nil {
		return
	}
	// Sqlite cannot handle concurrent writes, so we limit sqlite to one connection.
	// see https://github.com/mattn/go-sqlite3/issues/274
	Mysql.DB().SetMaxOpenConns(1)
	Mysql.SetLogger(DefaultGormLogger)
	Mysql.LogMode(false)
	return
}

func Close() {
	if Mysql != nil {
		Mysql.Close()
		Mysql = nil
	}
}
