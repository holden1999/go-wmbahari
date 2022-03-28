package cli

import (
	"fmt"
	"go-wmb/usecase"
	"os"
)

func TableForm(usecase usecase.ShowListUseCase) {
	var choice string
	fmt.Println("Nomor Meja     Status Meja Kososng")
	for _, table := range usecase.GetTable() {
		fmt.Printf("%d %20t\n", table.GetTableNumber(), table.GetTableStatus())
	}
	fmt.Print("Kembali ke Main Menu (y/n) ? ")
	fmt.Scan(&choice)
	if choice == "y" {
		MainMenu()
	} else {
		os.Exit(1)
	}
}
