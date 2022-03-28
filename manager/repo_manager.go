package manager

import (
	"go-wmb/repository"
)

type RepoManager interface {
	CustomerOrderRepo() repository.CustomerOrderRepo
	ShowListFoodRepo() repository.ListFoodRepo
	ShowListTableRepo() repository.ListTableRepo
	CustomerPaymentRepo() repository.CustomerPaymentRepo
}

type repoManager struct {
	infra Infra
}

func (r *repoManager) CustomerOrderRepo() repository.CustomerOrderRepo {
	return repository.NewCustomerOrderRepo(r.infra.SqlDb())
}

func (r *repoManager) CustomerPaymentRepo() repository.CustomerPaymentRepo {
	return repository.NewCustomerPaymentRepo(r.infra.SqlDb())
}

func (r *repoManager) ShowListFoodRepo() repository.ListFoodRepo {
	return repository.NewListFoodRepo(r.infra.SqlDb())
}

func (r *repoManager) ShowListTableRepo() repository.ListTableRepo {
	return repository.NewListTableRepo(r.infra.SqlDb())
}

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
