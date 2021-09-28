package storage

import (
	"log"

	"github.com/hjoshi123/seniorly_interview/config"
	"github.com/jinzhu/gorm"
)

func NewDB(params ...string) *gorm.DB {
	var err error
	conString := config.GetPostgresConnectionString()

	log.Print(conString)

	DB, err := gorm.Open(config.GetDBType(), conString)

	if err != nil {
		log.Panic(err)
	}

	return DB
}
