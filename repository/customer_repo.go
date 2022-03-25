package repository

import (
	"go-wmb/model"
)

type CustomerRepo interface {
	Order(newOrder model.Order) error
	Payment(newOrder model.Payment) []model.Payment
}
