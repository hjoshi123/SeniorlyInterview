package main

import (
	"net/http"

	"github.com/hjoshi123/seniorly_interview/controllers"
	"github.com/hjoshi123/seniorly_interview/model/pizza"
	"github.com/hjoshi123/seniorly_interview/storage"
	orderStore "github.com/hjoshi123/seniorly_interview/storage/pizza"
)

// func init() {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatalf("Error loading env file")
// 		os.Exit(0)
// 	}
// }

func main() {
	db, err := storage.NewDB()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	orderRepo := orderStore.New(db)
	orderService := pizza.NewService(orderRepo)

	httpRouter := controllers.NewHTTPHandler(orderService)

	err = http.ListenAndServe(":8080", httpRouter)
	if err != nil {
		panic(err)
	}
}
