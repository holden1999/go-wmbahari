package usecase

import (
	"go-wmb/model"
	"go-wmb/repository"
)

type PaymentUseCase interface {
	Payment(tableNumber int) ([]model.Payment, error)
}

type paymentUseCase struct {
	repo repository.PaymentRepo
}

func (p *paymentUseCase) Payment(table int) ([]model.Payment, error) {
	newTable := model.NewPayment(table)
	return p.repo.Payment(newTable)
}

func NewPaymentUseCase(repo repository.PaymentRepo) *paymentUseCase {
	return &paymentUseCase{repo: repo}
}
