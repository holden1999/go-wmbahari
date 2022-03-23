package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type cashierRepoImpl struct {
	cashierDb *sqlx.DB
}

func (c *cashierRepoImpl) OrderTable(number string) {
	//TODO implement me
	panic("implement me")
}

type tables struct {
	foodCode string `db:"f_code"`
	foodName string `db:"f_name"`
	price    int    `db:"price"`
}

func (c *cashierRepoImpl) GetFood() error {
	tables := []tables{}
	err := c.cashierDb.Select(tables, "select * from master_foods")
	if err != nil {
		fmt.Println(tables)
		errors.New("Failed Show Table")
	}
	return nil
}

func ListFoodRepo(cashierDb *sqlx.DB) CashierRepo {
	return &cashierRepoImpl{cashierDb: cashierDb}
}
