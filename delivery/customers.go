package delivery

type CustomerDelivery interface {
	Order(food, name string)
	Payment(code string)
}
