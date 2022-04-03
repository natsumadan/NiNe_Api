package datbase

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

//Database Connection File
func Connection() {
	data, err := gorm.Open("mysql", "root:Happy@1998877@1/nineapi?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = data
}

func GetDB() *gorm.DB {
	return db
}
