package manager

import "go-wmb/usecase"

type UseCaseManager interface {
	CustomerOrderUseCase() usecase.OrderUseCase
	FoodListUseCase() usecase.FoodUseCase
	TableListUseCase() usecase.TableUseCase
	CustomerPaymentUseCase() usecase.PaymentUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) TableListUseCase() usecase.TableUseCase {
	return usecase.NewTableListUseCase(u.repo.ShowListTableRepo())
}

func (u *useCaseManager) FoodListUseCase() usecase.FoodUseCase {
	return usecase.NewFoodListUseCase(u.repo.ShowListFoodRepo())
}

func (u *useCaseManager) CustomerOrderUseCase() usecase.OrderUseCase {
	return usecase.NewOrderUseCase(u.repo.CustomerOrderRepo())
}

func (u *useCaseManager) CustomerPaymentUseCase() usecase.PaymentUseCase {
	return usecase.NewPaymentUseCase(u.repo.CustomerPaymentRepo())
}

func NewUseCaseManager(manager RepoManager) *useCaseManager {
	return &useCaseManager{repo: manager}
}
