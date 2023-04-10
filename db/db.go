package db

import (
	"github.com/joshuaautawi/shuttle/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/shuttle_db"
	DB, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&user.User{})
	return DB

}
