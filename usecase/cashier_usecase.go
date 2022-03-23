package usecase

import (
	"errors"
	"go-wmb/repository"
)

type CashierUseCase interface {
	GetFood() error
	ListTable() error
	OrderTable(number string)
}

type cashierUseCase struct {
	repo repository.CashierRepo
}

func (c *cashierUseCase) ListTable() error {

}

func (c *cashierUseCase) OrderTable(number string) {
	//TODO implement me
	panic("implement me")
}

func (c *cashierUseCase) GetFood() error {
	err := c.repo.GetFood()
	if err != nil {
		return errors.New("GetFailed")
	}
	return nil
}
