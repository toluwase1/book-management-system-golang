package config

import (
	"github.com/jinzhu/gorm"
)
var (
	db *gorm.DB
)
func Connect() {
	d, err:= gorm.Open("mysql", "root:toluwase@tcp(127.0.0.1:3306)/golang-bookstore")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDb() *gorm.DB{
	return db
}