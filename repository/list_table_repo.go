package repository

import (
	"go-wmb/model"
	"gorm.io/gorm"
)

type TableRepo interface {
	GetTable() ([]model.Table, error)
}

type tableRepo struct {
	db *gorm.DB
}

func (t *tableRepo) GetTable() ([]model.Table, error) {
	tables := make([]model.Table, 0)
	err := t.db.Find(&tables).Error
	if err != nil {
		return nil, err
	}
	return tables, nil
}

func NewListTableRepo(db *gorm.DB) TableRepo {
	return &tableRepo{db: db}
}
