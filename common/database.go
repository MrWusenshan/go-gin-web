package common

import (
	"fmt"

	"github.com/spf13/viper"
	_ "go-gin-web/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbEngine *gorm.DB

func init() {
	host := viper.Get("database.postgresql.host")
	user := viper.Get("database.postgresql.user")
	password := viper.Get("database.postgresql.password")
	dbname := viper.Get("database.postgresql.dbname")
	port := viper.Get("database.postgresql.port")
	sslmode := viper.Get("database.postgresql.sslmodel")
	timezone := viper.Get("database.postgresql.timezone")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("The database connection failed!")
	}
	dbEngine = db
}

func GetDb() *gorm.DB {
	return dbEngine
}
