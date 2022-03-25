package manager

import "go-wmb/usecase"

type UseCaseManager interface {
	CustomerOrderUseCase() usecase.CustomerUseCase
	FoodListUseCase() usecase.ShowListUseCase
	TableListUseCase() usecase.ShowListUseCase
	CustomerPaymentUseCase() usecase.CustomerUseCase
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

func (u *useCaseManager) CustomerOrderUseCase() usecase.CustomerUseCase {
	return usecase.NewOrderUseCase(u.repo.CustomerOrderRepo())
}

func (u *useCaseManager) CustomerPaymentUseCase() usecase.CustomerUseCase {
	return usecase.NewPaymentUseCase(u.repo.CustomerPaymentRepo())
}

func NewUseCaseManager(manager RepoManager) *useCaseManager {
	return &useCaseManager{repo: manager}
}
