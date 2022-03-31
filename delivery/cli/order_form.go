package cli

import (
	"fmt"
	"go-wmb/delivery/utils"
	"go-wmb/usecase"
	"os"
)

func OrderForm(usecase usecase.OrderUseCase) {
	var (
		foodCode    string
		nameOrder   string
		tableNumber int
	)

	utils.CreateHeader()
	fmt.Println("Menu Order Makanan")

	fmt.Println("Nama Customer : ")
	fmt.Scan(&nameOrder)
	fmt.Println("Kode Makanan : ")
	fmt.Scan(&foodCode)
	fmt.Println("Pilih Meja : ")
	fmt.Scan(&tableNumber)
	utils.CreateHeader()

	if err := usecase.Order(foodCode, nameOrder, tableNumber); err != nil {
		os.Exit(1)
	}
	MainMenu()
}
