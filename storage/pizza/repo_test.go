package pizza

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	model "github.com/hjoshi123/seniorly_interview/model/pizza"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when opening a stub  database connection", err)
	}

	gormDB, err := gorm.Open("sqlite3", db)
	gormDB.LogMode(true)
	if err != nil {
		t.Fatalf("an error %s was not expected when opening a stub  database connection", err)
	}

	order := &model.Order{
		ID:        12,
		PizzaType: "Veggie Lover",
		Mobile:    "8762202041",
		Chef:      "Hemant",
		Status:    "preparing",
	}

	query := `"INSERT INTO "orders" ("id","pizza_type","mobile","status","chef") VALUES (?,?,?,?,?)"`

	mock.ExpectBegin()
	mock.ExpectExec(query).WithArgs(order.ID, order.PizzaType, order.Mobile, order.Status, order.Chef).WillReturnResult(sqlmock.NewResult(12, 1))
	mock.ExpectCommit()

	repo := New(gormDB)
	recvd, err := repo.CreateOrder(order)
	t.Log(recvd)
	assert.NoError(t, err)
	// assert.Equal(t, uint(12), recvOrder.ID)
}
