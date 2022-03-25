package usecase

import (
	"go-wmb/model"
	"go-wmb/repository"
)

type ShowListUseCase interface {
	GetFood() []model.FoodList
	GetTable() []model.TableList
}

type showListUseCase struct {
	repo repository.ShowListRepo
}

func (c *showListUseCase) GetTable() []model.TableList {
	return c.repo.GetTable()
}

func (c *showListUseCase) GetFood() []model.FoodList {
	return c.repo.GetFood()
}

func NewFoodListUseCase(repo repository.ShowListRepo) ShowListUseCase {
	return &showListUseCase{repo: repo}
}
func NewTableListUseCase(repo repository.ShowListRepo) ShowListUseCase {
	return &showListUseCase{repo: repo}
}
