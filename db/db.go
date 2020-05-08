package db

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Init(connString string) {
	db, err := gorm.Open("mysql", connString)
	db.LogMode(true)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.DB().SetMaxIdleConns(50)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

}
