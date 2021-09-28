package pizza

import "github.com/hjoshi123/seniorly_interview/utils"

type OrderRepository interface {
	CreateOrder(*Order) (*Order, error)
	TrackOrder(string) (string, error)
	UpdateOrder(uint, string) (*Order, error)
	GetOrderByMobileNumber(string) (*Order, *utils.AppError)
}
