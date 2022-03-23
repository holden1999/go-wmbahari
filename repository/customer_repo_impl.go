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
	var inc int
	c.customerDb.Select(&inc, "select count(*) from customer_table")
	err := c.customerDb.MustExec("insert into customer_orders (id, table_num, person, foods, f_quantity ) values ($1, $2, $3, $4, $5)", int(inc), int(inc), newOrder.GetOrderName(), newOrder.GetFoodName(), newOrder.GetFoodQuantity())
	fmt.Println(newOrder.GetFoodName())
	if err != nil {
		return errors.New("Query Error")
	}
	return nil
}

func NewCustomerRepo(customerDb *sqlx.DB) CustomerRepo {
	return &customerRepoImpl{customerDb: customerDb}
}
