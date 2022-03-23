package repository

type CashierRepo interface {
	OrderTable(number string)
	GetFood() error
}
