package storage

import (
	"log"

	"github.com/hjoshi123/seniorly_interview/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Import GORM postgres dialect for its side effects, according to GORM docs.
)

func NewDB(params ...string) (*gorm.DB, error) {
	var err error
	conString := config.GetPostgresConnectionString()

	log.Print(conString)

	DB, err := gorm.Open(config.GetDBType(), conString)

	return DB, err
}
