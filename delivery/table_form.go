package delivery

import (
	"fmt"
	"go-wmb/usecase"
)

func TableForm(usecase usecase.CashierUseCase) {
	var (
		orderTable string
	)
	fmt.Scanln(&orderTable)
	usecase.OrderTable(orderTable)
}
