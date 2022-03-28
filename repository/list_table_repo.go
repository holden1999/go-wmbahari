package repository

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"go-wmb/model"
)

type ListTableRepo interface {
	GetTable() []model.TableList
}

type listTableRepo struct {
	db *sqlx.DB
}

func (l *listTableRepo) GetTable() []model.TableList {
	tableData := []model.TableList{}
	err := l.db.Select(&tableData, "select * from master_table")
	if err != nil {
		errors.New("Failed showing Table List")
	}
	return tableData
}

func NewListTableRepo(db *sqlx.DB) ListTableRepo {
	return &listTableRepo{db: db}
}
