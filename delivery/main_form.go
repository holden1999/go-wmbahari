package delivery

import (
	"fmt"
	"go-wmb/delivery/utils"
)

func MainMenu() {
	utils.CreateHeader()
	fmt.Println("Warung Makan Bahari")
	utils.CreateHeader()
	fmt.Println("1. Pesan Makanan")
	fmt.Println("2. Pilihan Meja")
	fmt.Println("3. Daftar Makanan")
	fmt.Println("4. Pembayaran")
}
