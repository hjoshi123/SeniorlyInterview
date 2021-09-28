package pizza

import (
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

			newOrder, err := service.CreateOrder(order)
			if err != nil {
				c.Error(err)
				return
			}

			c.JSON(http.StatusCreated, newOrder)
		})
	}

	return orderRoutesGroup
}
