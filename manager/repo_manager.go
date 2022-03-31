package manager

import (
	"go-wmb/repository"
)

type RepoManager interface {
	CustomerOrderRepo() repository.OrderRepo
	ShowListFoodRepo() repository.FoodRepo
	ShowListTableRepo() repository.TableRepo
	CustomerPaymentRepo() repository.PaymentRepo
}

type repoManager struct {
	infra Infra
}

func (r *repoManager) CustomerOrderRepo() repository.OrderRepo {
	return repository.NewCustomerOrderRepo(r.infra.SqlDb())
}

func (r *repoManager) CustomerPaymentRepo() repository.PaymentRepo {
	return repository.NewCustomerPaymentRepo(r.infra.SqlDb())
}

func (r *repoManager) ShowListFoodRepo() repository.FoodRepo {
	return repository.NewListFoodRepo(r.infra.SqlDb())
}

func (r *repoManager) ShowListTableRepo() repository.TableRepo {
	return repository.NewListTableRepo(r.infra.SqlDb())
}

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
