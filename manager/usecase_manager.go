package manager

import "go-wmb/usecase"

type UseCaseManager interface {
	CustomerOrderUseCase() usecase.CustomerOrderUseCase
	FoodListUseCase() usecase.ShowListUseCase
	TableListUseCase() usecase.ShowListUseCase
	CustomerPaymentUseCase() usecase.CustomerOrderUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) TableListUseCase() usecase.ShowListUseCase {
	return usecase.NewTableListUseCase(u.repo.ShowListTableRepo())
}

func (u *useCaseManager) FoodListUseCase() usecase.ShowListUseCase {
	return usecase.NewFoodListUseCase(u.repo.ShowListFoodRepo())
}

func (u *useCaseManager) CustomerOrderUseCase() usecase.CustomerOrderUseCase {
	return usecase.NewOrderUseCase(u.repo.CustomerOrderRepo())
}

func (u *useCaseManager) CustomerPaymentUseCase() usecase.CustomerOrderUseCase {
	return usecase.NewPaymentUseCase(u.repo.CustomerPaymentRepo())
}

func NewUseCaseManager(manager RepoManager) *useCaseManager {
	return &useCaseManager{repo: manager}
}
