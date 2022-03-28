package repository

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"go-wmb/model"
)

type CustomerOrderRepo interface {
	Order(newOrder model.Order) error
}

type customerOrderRepo struct {
	db *sqlx.DB
}

func (c *customerOrderRepo) Order(newOrder model.Order) error {
	err := c.db.MustExec("insert into customer_orders ( table_number, person, food_code) values ($1, $2, $3)", newOrder.GetTableNumber(), newOrder.GetOrderName(), newOrder.GetFoodName())
	c.db.MustExec("update master_table set available_status = false where id = $1 ", newOrder.GetTableNumber())
	if err != nil {
		return errors.New("Query Error")
	}
	return nil
}

func NewCustomerOrderRepo(db *sqlx.DB) CustomerOrderRepo {
	return &customerOrderRepo{db: db}
}
