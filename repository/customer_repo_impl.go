package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-wmb/model"
)

type customerRepoImpl struct {
	db *sqlx.DB
}

func (c *customerRepoImpl) Payment(newPayment model.Payment) []model.Payment {
	var totalInvoice int
	invoiceDetail := []model.Payment{}
	err := c.db.Select(&invoiceDetail, "select * from customer_orders inner join master_food on code = food_code where table_number = $1", newPayment.GetTableNumber())
	c.db.Select(&totalInvoice, "select sum(price) from master_food ")
	c.db.MustExec("update master_table set available_status = true where id = $1 ", newPayment.GetTableNumber())
	if err != nil {
		errors.New("Failed showing Invoice detail")
	}
	fmt.Println(invoiceDetail)
	return invoiceDetail
}

func (c *customerRepoImpl) Order(newOrder model.Order) error {
	err := c.db.MustExec("insert into customer_orders ( table_number, person, food_code) values ($1, $2, $3)", newOrder.GetTableNumber(), newOrder.GetOrderName(), newOrder.GetFoodName())
	c.db.MustExec("update master_table set available_status = false where id = $1 ", newOrder.GetTableNumber())
	if err != nil {
		return errors.New("Query Error")
	}
	return nil
}

func NewCustomerRepo(db *sqlx.DB) CustomerRepo {
	return &customerRepoImpl{db: db}
}

func NewPaymentRepo(db *sqlx.DB) CustomerRepo {
	return &customerRepoImpl{db: db}
}
