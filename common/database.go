package common

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbEngine *gorm.DB

func init() {
	dsn := "host=localhost user=postgres password=root dbname=gin_web port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("The database connection failed!")
	}
	dbEngine = db
}

func GetDb() *gorm.DB {
	return dbEngine
}
