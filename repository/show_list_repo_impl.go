package repository

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"go-wmb/model"
)

type showListRepoImpl struct {
	db *sqlx.DB
}

func (c *showListRepoImpl) GetTable() []model.TableList {
	tableData := []model.TableList{}
	err := c.db.Select(&tableData, "select * from master_table")
	if err != nil {
		errors.New("Failed showing Table List")
	}
	return tableData
}

func (c *showListRepoImpl) GetFood() []model.FoodList {
	foodData := []model.FoodList{}
	err := c.db.Select(&foodData, "select * from master_food")
	if err != nil {
		errors.New("Failed Showing Food List")
	}
	return foodData
}

func ListFoodRepo(db *sqlx.DB) ShowListRepo {
	return &showListRepoImpl{db: db}
}
func ListTableRepo(db *sqlx.DB) ShowListRepo {
	return &showListRepoImpl{db: db}
}
