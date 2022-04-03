package models

import (
	"github.com/jinzhu/gorm"
	"ninApi/pkg/datbase"
)

var db *gorm.DB

type Tags struct {
	Tag1 string `json:"tag1"`
	Tag2 string `json:"tag2"`
	Tag3 string `json:"tag3"`
}

type Article struct {
	gorm.Model
	Title string `gorm:""json:"title"`
	Date  string `json:"date"`
	Body  string `json:"body"`
	Tags  Tags   `json:"tags"`
}

func init() {
	datbase.Connection()
	db = datbase.GetDB()
	db.AutoMigrate(&Article{})
}

//Function for Creating a new Article
func (a *Article) CreateArticle() *Article {
	db.NewRecord(a)
	db.Create(&a)
	return a
}

//Function to Get Article from database
func GetArticle(Id int64) (*Article, *gorm.DB) {
	var getArticle Article
	db := db.Where("ID=?", Id).Find(&getArticle)
	return &getArticle, db
}
