package usecase

import (
	"go-wmb/model"
	"go-wmb/repository"
)

type ListTableUseCase interface {
	GetTable() []model.TableList
}

type listTableUseCase struct {
	repo repository.ListFoodRepo
}

func (c *listTableUseCase) GetTable() []model.TableList {
	return c.repo.GetTable()
}

func NewTableListUseCase(repo repository.ListFoodRepo) ListTableUseCase {
	return &listTableUseCase{repo: repo}
}
