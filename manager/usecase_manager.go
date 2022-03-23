package manager

import "go-wmb/usecase"

type UseCaseManager interface {
	CustomerOrderUseCase() usecase.CustomerUseCase
	ListFoodUseCase() usecase.CashierUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) ListFoodUseCase() usecase.CashierUseCase {
	return usecase.ListFoodUseCase(u.repo.CashierRepo())
}

func (u *useCaseManager) OrderTableUseCase() usecase.CashierUseCase {
	//TODO implement me
	panic("implement me")
}

func (u *useCaseManager) CustomerOrderUseCase() usecase.CustomerUseCase {
	return usecase.NewOrderUseCase(u.repo.CustomerRepo())
}

func NewUseCaseManager(manager RepoManager) *useCaseManager {
	return &useCaseManager{
		repo: manager,
	}
}
