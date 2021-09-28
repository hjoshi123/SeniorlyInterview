package pizza

import (
	"math/rand"

	model "github.com/hjoshi123/seniorly_interview/model/pizza"
	"github.com/hjoshi123/seniorly_interview/utils"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

const (
	createError = "Error in creating new order"
	readError   = "Error in finding order in the database"
)

type Store struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Store {
	db.AutoMigrate(&model.Order{})

	return &Store{db: db}
}

func (s *Store) CreateOrder(order *model.Order) (*model.Order, error) {
	chefs := []string{"Mark", "hemant", "Hello"}

	chefRandom := chefs[rand.Intn(len(chefs))]

	order.Chef = chefRandom

	if err := s.db.Create(order).Error; err != nil {
		return nil, err
	}

	return order, nil
}

func (s *Store) TrackOrder(mobileNumber string) (string, error) {
	res := &model.Order{}

	query := s.db.Where("mobile = ?", mobileNumber).First(res)

	if query.RecordNotFound() {
		err := utils.NewAppErrorWithType(utils.NotFound)
		return "", err
	}

	if err := query.Error; err != nil {
		errRet := utils.NewAppError(errors.Wrap(err))
	}
}
