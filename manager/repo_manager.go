package manager

import "go-wmb/repository"

type RepoManager interface {
	CustomerRepo() repository.CustomerRepo
	CashierRepo() repository.CashierRepo
}

type repoManager struct {
	infra Infra
}

func (r *repoManager) CashierRepo() repository.CashierRepo {
	return repository.ListFoodRepo(r.infra.SqlDb())
}

func (r *repoManager) CustomerRepo() repository.CustomerRepo {
	return repository.NewCustomerRepo(r.infra.SqlDb())
}

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
