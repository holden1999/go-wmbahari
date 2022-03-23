package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-wmb/model"
)

type customerRepoImpl struct {
	customerDb *sqlx.DB
}

func (c *customerRepoImpl) Order(newOrder model.Order) error {
	err := c.customerDb.MustExec("insert into customer_order (customer_name, food_order) values ($1, $2)", newOrder.GetOrderName(), newOrder.GetFoodName())
	fmt.Println(newOrder.GetFoodName())
	if err != nil {
		return errors.New("Query Error")
	}
	return nil
}

func NewCustomerRepo(customerDb *sqlx.DB) CustomerRepo {
	return &customerRepoImpl{customerDb: customerDb}
}
