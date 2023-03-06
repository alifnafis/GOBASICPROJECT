package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "Emc002:PureAtom99,1%/simplerest?charset=utf8&parseTime=True&local")
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
