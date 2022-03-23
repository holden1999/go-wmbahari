package delivery

type CashierDelivery interface {
	ListFood()
	ListTable()
	OrderTable(number string)
}
