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
			if err != nil && err.Error() == "user exists with an order in previous" {
				c.Error(err)
				c.JSON(http.StatusOK, gin.H{
					"message": "Only one order can be active",
					"order":   newOrder,
				})
				return
			}

			log.Println(newOrder)
			c.JSON(http.StatusCreated, gin.H{
				"message": "Order created",
				"order":   newOrder,
			})
		})

		group.GET("/track_pizza", func(c *gin.Context) {
			mobileNumber := c.Query("mobile")

			status, err := service.TrackOrder(mobileNumber)
			if err != nil {
				c.Error(err)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "Your order status is " + status,
			})
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

			c.JSON(http.StatusOK, gin.H{
				"message": "Updated successfully",
				"order": update,
			})
		})
	}

	return orderRoutesGroup
}
