package pizza

type OrderRepository interface {
	CreateOrder(*Order) (*Order, error)
	TrackOrder(string) (string, error)
	UpdateOrder(string) (*Order, error)
}
