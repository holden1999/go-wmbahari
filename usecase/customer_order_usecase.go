package usecase

import (
	"errors"
	"go-wmb/model"
	"go-wmb/repository"
)

type OrderUseCase interface {
	Order(code string, name string, table int) error
}

type orderUseCase struct {
	repo repository.OrderRepo
}

func (c *orderUseCase) Order(code string, name string, table int) error {
	newOrder := model.NewOrder(code, name, table)
	err := c.repo.Order(newOrder)
	if err != nil {
		return errors.New("Failed Update Struct")
	}
	return nil
}

func NewOrderUseCase(repo repository.OrderRepo) *orderUseCase {
	return &orderUseCase{repo: repo}
}
