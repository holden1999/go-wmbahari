package usecase

import (
	"errors"
	"go-wmb/model"
	"go-wmb/repository"
)

type CustomerUseCase interface {
	Order(code string, name string, table int) error
	Payment(tableNumber int) []model.Payment
}

type customerUseCase struct {
	repo repository.CustomerRepo
}

func (c *customerUseCase) Payment(table int) []model.Payment {
	newTable := model.NewPayment(table)
	return c.repo.Payment(newTable)
}

func (c *customerUseCase) Order(code string, name string, table int) error {
	newOrder := model.NewOrder(code, name, table)
	err := c.repo.Order(newOrder)
	if err != nil {
		return errors.New("Failed Update Struct")
	}
	return nil
}

func NewOrderUseCase(repo repository.CustomerRepo) *customerUseCase {
	return &customerUseCase{repo: repo}
}

func NewPaymentUseCase(repo repository.CustomerRepo) *customerUseCase {
	return &customerUseCase{repo: repo}
}
