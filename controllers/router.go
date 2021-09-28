package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	controller "github.com/hjoshi123/seniorly_interview/controllers/pizza"
	model "github.com/hjoshi123/seniorly_interview/model/pizza"
)

func LogRequest(c string) gin.HandlerFunc {
	return func(c1 *gin.Context) {
		log.Println("Middleware" + c)
		c1.Next()
		log.Println("Exit log")
	}
}

func NewHTTPHandler(service model.OrderService) http.Handler {
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("pizzaType", controller.VerifyPizzaType)
		v.RegisterValidation("updateStatus", controller.VerifyUpdateStatus)
	}

	api := router.Group("/api")
	api.Use(LogRequest("share"))

	controller.NewRoutesFactory(api)(service)

	return router
}
