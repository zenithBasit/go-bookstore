package models

import (
	"bookstore/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	if err := db.First(b).Error; err != nil && gorm.ErrRecordNotFound == err {
		db.Create(&b)
	}
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(ID int64) (*Book, *gorm.DB) {
	var Book Book
	db := db.Where("ID=?", ID).Find(Book)
	return &Book, db
}

func DeleteBook(ID int64) Book {
	var Book Book
	db.Where("ID=?", ID).Delete(Book)
	return Book
}
