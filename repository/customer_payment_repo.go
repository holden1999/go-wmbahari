package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go-wmb/model"
)

type CustomerPaymentRepo interface {
	Payment(newPayment model.Payment) error
}

type customerPaymentRepo struct {
	db *sqlx.DB
}

func (c *customerPaymentRepo) Payment(newPayment model.Payment) error {
	err := c.db.MustExec("select customer_orders ( table_number, person, food_code) values ($1, $2, $3)", newOrder.GetTableNumber(), newOrder.GetOrderName(), newOrder.GetFoodName())
	c.db.MustExec("update master_table set available_status = false where id = $1 ", newOrder.GetTableNumber())
	if err != nil {
		return errors.New("Query Error")
	}
	return nil
}

func NewCustomerPaymentRepo(db *sqlx.DB) CustomerPaymentRepo {
	return &customerPaymentRepo{db: db}
}
