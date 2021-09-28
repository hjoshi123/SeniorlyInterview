package pizza

type OrderService interface {
	CreateOrder(*Order) (*Order, error)
	TrackOrder(string) (string, error)
	UpdateOrder(string) (*Order, error)
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

func (sv *Service) UpdateOrder(chef string) (*Order, error) {
	return sv.repository.UpdateOrder(chef)
}
