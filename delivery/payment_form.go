package delivery

import (
	"fmt"
	"go-wmb/usecase"
	"os"
)

func PaymentForm(usecase usecase.CustomerUseCase) {
	var (
		tableNumber int
		choice      string
	)
	fmt.Println("Pembayaran Nomor Meja : ")
	fmt.Scan(&tableNumber)
	for idx, detail := range usecase.Payment(tableNumber) {
		fmt.Printf("%d %s %s %s %d\n", idx+1, detail.GetTableNumber(), detail.GetOrderName(), detail.GetFoodCode(), detail.GetFoodName(), detail.GetFoodPrice())
	}
	fmt.Print("Kembali ke Main Menu (y/n) ? ")
	fmt.Scan(&choice)
	if choice == "y" {
		MainMenu()
	} else {
		os.Exit(1)
	}
}
