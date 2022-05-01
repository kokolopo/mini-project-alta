package database

import (
	"fmt"
	"order_kafe/config"
	"order_kafe/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitDatabase(config config.Config) {

	//dsn := "username:password@tcp(host:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := config.DB_USERNAME + ":" + config.DB_PASSWORD + "@tcp(" + config.DB_HOST + ":" + config.DB_PORT + ")/" + config.DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	} else {
		fmt.Println("koneksi ke database berhasil!!!")
	}
	DB.AutoMigrate(&user.User{})
}
