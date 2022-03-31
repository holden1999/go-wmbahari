package repository

import (
	"go-wmb/model"
	"gorm.io/gorm"
)

type FoodRepo interface {
	GetFood() ([]model.Food, error)
}

type foodRepo struct {
	db *gorm.DB
}

func (f *foodRepo) GetFood() ([]model.Food, error) {
	foods := make([]model.Food, 0)
	err := f.db.Select(&foods).Error
	if err != nil {
		return nil, err
	}
	return foods, nil
}

func NewListFoodRepo(db *gorm.DB) FoodRepo {
	return &foodRepo{db: db}
}
