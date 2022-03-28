package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go-wmb/model"
)

type ListFoodRepo interface {
	GetFood() []model.FoodList
}

type listFoodRepo struct {
	db *sqlx.DB
}

func (c *listFoodRepo) GetFood() []model.FoodList {
	foodData := []model.FoodList{}
	err := c.db.Select(&foodData, "select * from master_food")
	if err != nil {
		errors.New("Failed Showing Food List")
	}
	return foodData
}

func NewListFoodRepo(db *sqlx.DB) ListFoodRepo {
	return &listFoodRepo{db: db}
}
