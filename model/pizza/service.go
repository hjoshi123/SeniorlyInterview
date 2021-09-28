package pizza

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type OrderService interface {
	CreateOrder(*Order) (*Order, error)
	TrackOrder(string) (string, error)
	UpdateOrder(uint, string) (*Order, error)
}

type Service struct {
	repository OrderRepository
}

func NewService(repo OrderRepository) *Service {
	return &Service{repository: repo}
}

func (sv *Service) CreateOrder(order *Order) (*Order, error) {
	return sv.repository.CreateOrder(order)
}

func (sv *Service) TrackOrder(mobileNumber string) (string, error) {
	return sv.repository.TrackOrder(mobileNumber)
}

func (sv *Service) UpdateOrder(id uint, status string) (*Order, error) {
	order, err := sv.repository.UpdateOrder(id, status)
	if err != nil {
		return nil, err
	}

	if order.Status == "done" {
		file, _ := json.MarshalIndent(order, "", " ")

		log.Println("Sending notification to user")

		os.Mkdir("test", 0777)

		fileName := path.Join("test", "file.txt")

		err = ioutil.WriteFile(fileName, file, 0666)
		if err != nil {
			log.Println("Notification error")
		}
	}

	return order, nil
}
