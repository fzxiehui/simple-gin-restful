package global

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/spf13/viper"
	"log"
)

var DB *gorm.DB

func InitDB() {
	var err error
	// DB, err = gorm.Open(viper.GetString("db.type"), viper.GetString("db.url"))
	DB, err = gorm.Open(sqlite.Open(viper.GetString("db.url")), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
}

func CloseDB() {
	db, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}

func GetDB() *gorm.DB {
	return DB
}

