package usecase

import (
	"errors"
	"go-wmb/model"
	"go-wmb/repository"
)

type CustomerOrderUseCase interface {
	Order(code string, name string, table int) error
}

type customerOrderUseCase struct {
	repo repository.CustomerOrderRepo
}

func (c *customerOrderUseCase) Order(code string, name string, table int) error {
	newOrder := model.NewOrder(code, name, table)
	err := c.repo.Order(newOrder)
	if err != nil {
		return errors.New("Failed Update Struct")
	}
	return nil
}

func NewOrderUseCase(repo repository.CustomerOrderRepo) *customerOrderUseCase {
	return &customerOrderUseCase{repo: repo}
}
