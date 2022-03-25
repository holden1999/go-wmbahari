package repository

import "go-wmb/model"

type ShowListRepo interface {
	GetFood() []model.FoodList
	GetTable() []model.TableList
}
