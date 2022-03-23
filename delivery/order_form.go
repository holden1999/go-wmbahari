package delivery

import (
	"fmt"
	"go-wmb/usecase"
	"os"
)

func OrderForm(usecase usecase.CustomerUseCase) {
	var (
		foodOrder string
		nameOrder string
	)
	fmt.Println("Nama Customer")
	fmt.Scanln(&nameOrder)
	fmt.Println("Pilihan Makanan")
	fmt.Scanln(&foodOrder)

	if err := usecase.Order(foodOrder, nameOrder); err != nil {
		os.Exit(1)
	}
	MainMenu()
}
