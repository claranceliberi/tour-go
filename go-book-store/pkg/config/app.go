package config

import (
	"github.com/jinzhu/gorm" // indirect
	"github.com/go-sql-driver/mysql" // indirect
)

var (
	db * gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql", "liberi:liberi/learngo?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	db = d

}

func getDB() *gorm.DB{
	return db;
}