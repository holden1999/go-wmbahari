package repository

type CashierRepo interface {
	GetFood() error
	ListTable() error
	OrderTable(number string)
}
