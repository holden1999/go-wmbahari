package usecase

import (
	"go-wmb/model"
	"go-wmb/repository"
)

type TableUseCase interface {
	GetTable() ([]model.Table, error)
}

type tableUseCase struct {
	repo repository.TableRepo
}

func (t *tableUseCase) GetTable() ([]model.Table, error) {
	return t.repo.GetTable()
}

func NewTableListUseCase(repo repository.TableRepo) TableUseCase {
	return &tableUseCase{repo: repo}
}
