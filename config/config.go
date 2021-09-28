package config

import (
	"fmt"
	"os"
)

func GetDBType() string {
	return "postgres"
}

func GetPostgresConnectionString() string {
	var (
		DBUser     = os.Getenv("POSTGRES_USER")
		DBPassword = os.Getenv("POSTGRES_PASSWORD")
		DBName     = os.Getenv("POSTGRES_DB")
		DBHost     = "database"
		DBPort     = "5432"
	)

	dataBase := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		DBHost,
		DBPort,
		DBUser,
		DBName,
		DBPassword)
	return dataBase
}
