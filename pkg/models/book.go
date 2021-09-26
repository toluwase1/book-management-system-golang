package models

import (
	"BOOK-MANAGEMENT-SYSTEM/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication string `json:""publication`
}

func init(){
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}
/*
Control first is in the routes
The routes gives the control to the controller
Controllers will give control to the model (Book.go)

1. User interacts with the route
2. The routes send control to the controller (where we have our logic)
3. Controllers now performs the operation with the database.
the operations of the database has to reside in the models file books.go
 */


func(b *Book) CreateBook() *Book{
	db.NewRecord(b) //our orm
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book
	db:= db.Where("ID", Id)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(&book)
	return book
}

//you can create update later
//for now we can use get, delete and post to update