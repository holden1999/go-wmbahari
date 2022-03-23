package usecase

import (
	"errors"
	"go-wmb/model"
	"go-wmb/repository"
)

type CustomerUseCase interface {
	Order(food, name string) error
}

type customerUseCase struct {
	repo repository.CustomerRepo
}

func (c *customerUseCase) Order(food, name string) error {
	newOrder := model.NewOrder(food, name)
	err := c.repo.Order(newOrder)
	if err != nil {
		return errors.New("Order Failed")
	}
	return nil
}

func NewOrderUseCase(repo repository.CustomerRepo) *customerUseCase {
	return &customerUseCase{repo: repo}
}

func ListFoodUseCase(repo repository.CashierRepo) *cashierUseCase {
	return &cashierUseCase{repo: repo}
}
