package manager

import (
	"go-wmb/repository"
)

type RepoManager interface {
	CustomerOrderRepo() repository.CustomerRepo
	ShowListFoodRepo() repository.ShowListRepo
	ShowListTableRepo() repository.ShowListRepo
	CustomerPaymentRepo() repository.CustomerRepo
}

type repoManager struct {
	infra Infra
}

func (r *repoManager) CustomerOrderRepo() repository.CustomerRepo {
	return repository.NewCustomerRepo(r.infra.SqlDb())
}

func (r *repoManager) CustomerPaymentRepo() repository.CustomerRepo {
	return repository.NewPaymentRepo(r.infra.SqlDb())
}

func (r *repoManager) ShowListFoodRepo() repository.ShowListRepo {
	return repository.ListFoodRepo(r.infra.SqlDb())
}

func (r *repoManager) ShowListTableRepo() repository.ShowListRepo {
	return repository.ListTableRepo(r.infra.SqlDb())
}

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
