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

