package mypkg

import (
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const(
	DBNAME = "study"
	USERNAME = "root"
	PASSWORD = ""
	IPADDRESS = "127.0.0.1"
	PORTNUMBER = "3306"
)

func DatabaseConnection()(*gorm.DB, error){
	dsn := USERNAME + ":" + PASSWORD + "@tcp(" + IPADDRESS + ":" + PORTNUMBER + ")/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db,err
}
//DB初期化
func DbInit(){
	db, err := DatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&Todo{})
    }
	
//DB追加
func DbInsert(text string, status string) {
	db, err := DatabaseConnection()
	if err != nil {
		panic("データベース開けず")
	}
	db.Create(&Todo{Text: text, Status: status})
}