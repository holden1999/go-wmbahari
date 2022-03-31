package repository

import (
	"errors"
	"go-wmb/model"
	"gorm.io/gorm"
)

type OrderRepo interface {
	Order(newOrder model.Order) error
}

type orderRepo struct {
	db *gorm.DB
}

func (o *orderRepo) Order(newOrder model.Order) error {
	err := o.db.Select("Name", "DateTime").Create(&newOrder)
	o.db.Model(&newOrder).Where("TableNumber = ?", newOrder.TableNumber).Update("IsReserved", true)
	if err != nil {
		return errors.New("Query Error")
	}
	return nil
}

func NewCustomerOrderRepo(db *gorm.DB) OrderRepo {
	return &orderRepo{db: db}
}
