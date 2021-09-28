package pizza

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	model "github.com/hjoshi123/seniorly_interview/model/pizza"
)

type OrderValidator struct {
	PizzaType string `binding:"required,pizzaType" json:"type"`
	Mobile    string `binding:"required,min=10,max=10" json:"mobile"`
}

type UpdateStatus struct {
	ID     uint    `binding:"required" json:"id"`
	Status string `binding:"required,updateStatus" json:"status"`
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

var VerifyUpdateStatus validator.Func = func(fl validator.FieldLevel) bool {
	updateStatus, ok := fl.Field().Interface().(string)
	if ok {
		if updateStatus == "dough-prep" || updateStatus == "oven-bake" || updateStatus == "topping-art" || updateStatus == "done" {
			return true
		}
	}

	return false
}

func Bind(c *gin.Context) (*model.Order, error) {
	var validator OrderValidator
	log.Println("Enterring validator")
	if err := c.ShouldBindJSON(&validator); err != nil {
		return nil, err
	}

	order := &model.Order{
		PizzaType: validator.PizzaType,
		Mobile:    validator.Mobile,
	}

	return order, nil
}

func UpdateBind(c *gin.Context) (uint, string, error) {
	var validator UpdateStatus
	if err := c.ShouldBindJSON(&validator); err != nil {
		return 0, "", err
	}

	return validator.ID, validator.Status, nil
}
