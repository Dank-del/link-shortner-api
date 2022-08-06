package database

import (
	"log"

	"github.com/Dank-del/link-shortner-api/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func StartDatabase() {
	db, err := gorm.Open(sqlite.Open("server.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		log.Fatalln(err)
	}
	Session = db
	err = Session.AutoMigrate(&types.ShortenedLink{})
	if err != nil {
		log.Fatalln(err)
	}

}
