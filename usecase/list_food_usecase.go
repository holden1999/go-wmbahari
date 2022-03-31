package usecase

import (
	"go-wmb/model"
	"go-wmb/repository"
)

type FoodUseCase interface {
	GetFood() ([]model.Food, error)
}

type foodUseCase struct {
	repo repository.FoodRepo
}

func (f *foodUseCase) GetFood() ([]model.Food, error) {
	return f.repo.GetFood()
}

func NewFoodListUseCase(repo repository.FoodRepo) FoodUseCase {
	return &foodUseCase{repo: repo}
}
