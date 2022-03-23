package model

import "github.com/jmoiron/sqlx"

type Order struct {
	foodName     string
	orderName    string
	foodQuantity string
}

func (o *Order) GetFoodName() string {
	return o.foodName
}

func (o *Order) GetOrderName() string {
	return o.orderName
}
func (o *Order) GetFoodQuantity() string {
	return o.foodQuantity
}

func (f *Order) Order(newOrder Order) sqlx.DB {
	return sqlx.DB{}
}

func Search(foodCode string) string {
	return ""
}

func NewOrder(foodName, orderName, foodQuantity string) Order {
	return Order{
		foodName:     foodName,
		orderName:    orderName,
		foodQuantity: foodQuantity,
	}
}
