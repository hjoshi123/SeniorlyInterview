package pizza

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hjoshi123/seniorly_interview/model/pizza"
	"github.com/hjoshi123/seniorly_interview/utils"
)

func NewRoutesFactory(group *gin.RouterGroup) func(service pizza.OrderService) {
	orderRoutesGroup := func(service pizza.OrderService) {
		group.POST("/buy_pizza", func(c *gin.Context) {
			order, err := Bind(c)
			if err != nil {
				appError := utils.NewAppError(err, utils.ValidationError)
				c.Error(appError)
				return
			}

			log.Println(order.Mobile)
			newOrder, err := service.CreateOrder(order)
			if err != nil {
				c.Error(err)
				return
			}

			log.Println(newOrder)
			c.JSON(http.StatusCreated, newOrder)
		})

		group.GET("/track_pizza", func(c *gin.Context) {
			mobileNumber := c.Query("mobile")

			status, err := service.TrackOrder(mobileNumber)
			if err != nil {
				c.Error(err)
				return
			}

			c.JSON(http.StatusOK, status)
		})

		group.PUT("/update_status", func(c *gin.Context) {
			chef, status, err := UpdateBind(c)

			if err != nil {
				appError := utils.NewAppError(err, utils.ValidationError)
				c.Error(appError)
				return
			}

			update, err := service.UpdateOrder(chef, status)
			if err != nil {
				c.Error(err)
				return
			}

			c.JSON(http.StatusOK, update)
		})
	}

	return orderRoutesGroup
}
