package usecase

import (
	"go-wmb/model"
	"go-wmb/repository"
)

type CustomerPaymentUseCase interface {
	Payment(tableNumber int) []model.Payment
}

type customerPaymentUseCase struct {
	repo repository.CustomerOrderRepo
}

func (c *customerPaymentUseCase) Payment(table int) []model.Payment {
	newTable := model.NewPayment(table)
	return c.repo.Payment(newTable)
}

func NewPaymentUseCase(repo repository.CustomerOrderRepo) *customerPaymentUseCase {
	return &customerPaymentUseCase{repo: repo}
}
