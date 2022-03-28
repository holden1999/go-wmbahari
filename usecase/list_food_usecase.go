package usecase

import (
	"go-wmb/model"
	"go-wmb/repository"
)

type ShowListUseCase interface {
	GetFood() []model.FoodList
}

type showListUseCase struct {
	repo repository.ListFoodRepo
}

func (c *showListUseCase) GetFood() []model.FoodList {
	return c.repo.GetFood()
}

func NewFoodListUseCase(repo repository.ListFoodRepo) ShowListUseCase {
	return &showListUseCase{repo: repo}
}
