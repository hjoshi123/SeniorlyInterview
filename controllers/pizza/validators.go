package pizza

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hjoshi123/seniorly_interview/model/pizza"
)

type OrderValidator struct {
	PizzaType string `binding:"required,pizzaType" json:"type"`
	Mobile string `binding:"required,min=10,max=10" json:"mobile"`
}

var VerifyPizzaType validator.Func = func(fl validator.FieldLevel) bool {
	pizzaType, ok := fl.Field().Interface().(string)
	if ok {
		if pizzaType == "Veggie Lovers" || pizzaType == "Meat Lovers" {
			return true
		}
	}

	return false
}

func Bind(c *gin.Context) (*pizza.Order, error) {
	var validator OrderValidator
	if err := c.ShouldBindJSON(&validator); err != nil {
		return nil, err
	}

	order := &pizza.Order{
		PizzaType: validator.PizzaType,
		Mobile: validator.Mobile,
	}

	return order, nil
}