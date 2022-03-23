package usecase

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"go-wmb/repository"
)

type CashierUseCase interface {
	GetFood() error
	ListTable() sqlx.DB
	OrderTable(number string)
}

type cashierUseCase struct {
	repo repository.CashierRepo
}

func (c *cashierUseCase) ListTable() sqlx.DB {
	//TODO implement me
	panic("implement me")
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
