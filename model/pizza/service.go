package pizza

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/hjoshi123/seniorly_interview/utils"
)

type OrderService interface {
	CreateOrder(*Order) (*Order, error)
	TrackOrder(string) (string, error)
	UpdateOrder(uint, string) (*Order, error)
	GetOrderByMobileNumber(string) (*Order, *utils.AppError)
}

type Service struct {
	repository OrderRepository
}

func NewService(repo OrderRepository) *Service {
	return &Service{repository: repo}
}

func (sv *Service) CreateOrder(order *Order) (*Order, error) {
	prevOrder, err := sv.repository.GetOrderByMobileNumber(order.Mobile)
	if err != nil && err.Type == utils.NotFound {
		return sv.repository.CreateOrder(order)
	} else {
		if prevOrder.Status == "done" {
			return sv.repository.CreateOrder(order)
		}
		return prevOrder, errors.New("user exists with an order in previous")
	}
}

func (sv *Service) TrackOrder(mobileNumber string) (string, error) {
	return sv.repository.TrackOrder(mobileNumber)
}

func (sv *Service) GetOrderByMobileNumber(mobileNumber string) (*Order, *utils.AppError) {
	return sv.repository.GetOrderByMobileNumber(mobileNumber)
}

func (sv *Service) UpdateOrder(id uint, status string) (*Order, error) {
	order, err := sv.repository.UpdateOrder(id, status)
	if err != nil {
		return nil, err
	}

	if order.Status == "done" {
		file, _ := json.Marshal(order)

		log.Println("Sending notification to user")

		os.Mkdir("/tmp/test", 0777)

		fileName := path.Join("/tmp/test", "file.json")

		err = ioutil.WriteFile(fileName, file, 0666)
		log.Printf("Order of user with order id %d is done by chef %s", order.ID, order.Chef)
		if err != nil {
			log.Println("Notification error")
		}
	}

	return order, nil
}
