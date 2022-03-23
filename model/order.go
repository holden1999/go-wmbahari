package model

import "github.com/jmoiron/sqlx"

type Order struct {
	foodName  string
	orderName string
}

func (o *Order) GetFoodName() string {
	return o.foodName
}

func (o *Order) GetOrderName() string {
	return o.orderName
}

func (f *Order) Order(newOrder Order) sqlx.DB {
	return sqlx.DB{}
}

func Search(foodCode string) string {
	return ""
}

func NewOrder(foodName, orderName string) Order {
	return Order{
		foodName:  foodName,
		orderName: orderName,
	}
}
