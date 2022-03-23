package delivery

import (
	"fmt"
	"go-wmb/delivery/utils"
	"go-wmb/usecase"
	"os"
)

func OrderForm(usecase usecase.CustomerUseCase) {
	var (
		foodOrder    string
		nameOrder    string
		foodQuantity string
	)

	utils.CreateHeader()
	fmt.Println("Menu Order Makanan")
	fmt.Println("Nama Customer : ")
	fmt.Scan(&nameOrder)
	fmt.Println("Pilihan Makanan : ")
	fmt.Scan(&foodOrder)
	fmt.Println("Masukkan Jumlah porsi : ")
	fmt.Scan(&foodQuantity)
	utils.CreateHeader()

	if err := usecase.Order(foodOrder, nameOrder, foodQuantity); err != nil {
		os.Exit(1)
	}
	MainMenu()
}
