package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(DBASE string, DBHOST string, DBPORT string, DBUSERNAME string, DBPASSWORD string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUSERNAME, DBPASSWORD, DBHOST, DBPORT, DBASE)
	fmt.Println(dsn)
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	var Condb *gorm.DB
	Condb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	//	database.AutoMigrate(&Account{})
	Condb.AutoMigrate(&BlogPost{})
	DB = Condb
}
