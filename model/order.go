package model

type Order struct {
	FoodCode    string
	OrderName   string
	TableNumber int
}

func (o *Order) GetFoodName() string {
	return o.FoodCode
}

func (o *Order) GetOrderName() string {
	return o.OrderName
}

func (o *Order) GetTableNumber() int {
	return o.TableNumber
}

//func (f *Order) Order(Order) sqlx.DB {
//	return sqlx.DB{}
//}

func NewOrder(foodCode, orderName string, tableNumber int) Order {
	return Order{
		FoodCode:    foodCode,
		OrderName:   orderName,
		TableNumber: tableNumber,
	}
}
