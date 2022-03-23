package delivery

import (
	"fmt"
	"go-wmb/usecase"
)

func ListFoodForm(usecase usecase.CashierUseCase) {
	fmt.Println(usecase.GetFood())
}
