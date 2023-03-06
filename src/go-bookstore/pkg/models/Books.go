package modeld

import (
	"github.com/Emc002/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string ` json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
