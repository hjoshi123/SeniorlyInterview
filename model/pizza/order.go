package pizza

type Order struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	PizzaType string `json:"type"`
	Mobile    string `json:"mobile"`
	Status    string `json:"status" gorm:"default:preparing"`
	Chef      string `json:"chef"`
}
