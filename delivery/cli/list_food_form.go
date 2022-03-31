package cli

import (
	"fmt"
	"go-wmb/usecase"
	"os"
)

func ListFoodForm(usecase usecase.FoodUseCase) {
	var choice string
	for idx, food := range usecase.GetFood() {
		fmt.Printf("%-3d%-20s%-20sRp. %-20d\n", idx+1, food.GetFoodCode(), food.GetFoodName(), food.GetFoodPrice())
	}
	fmt.Print("Kembali ke Main Menu (y/n) ? ")
	fmt.Scan(&choice)
	if choice == "y" {
		MainMenu()
	} else {
		os.Exit(1)
	}
}
