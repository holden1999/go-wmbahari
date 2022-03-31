package main

/*Applikasi warung makan

Warung Makan Bahari (WMB) adalah sebuah warteg yg sangat laku dan terkenal di daerah ragunan.
Karena terkenal hingga sering direview oleh youtuber, WMB mendapatkan suntikan dana dari investor. Setelah dana investasi cair, WMB melakukan upgrade terhadap tempat makan nya yaitu meja makan dan sistem pemesanan dan pembayaran.
Kini meja makan sudah terpisah-pisah dan memiliki nomor masing-masing.
Pemesanan makan pun sudah menggunakan aplikasi di kasir, dan pembayaran dilakukan setelah selesai makan.
ketentuan :
-total meja makan ada 30 meja
-sebelum ke meja makan, konsumen harus pesan makanan di kasir, baru akan diarahkan ke meja tertentu
-untuk meja yg sudah ada konsumen, tidak bisa lagi dipilih
-saat akan memilih meja, kasir dapat melihat meja mana yg masih kosong / available
-Kasir Dapat melihat daftar makanan/mencari makanan berdasarkan kode/nama
-saat konsumen memesan makanan, kasir menginput data pesanan, nama konsumen, dan juga nomor meja yg akan ditempati konsumen
-setelah konsumen selesai makan, dapat melakukan pembayaran di kasir dan kemudian meja tersebut sudah available untuk konsumen berikutnya
*/

//func init() {
//	viper.SetConfigFile(`config.yaml`)
//	err := viper.ReadInConfig()
//	if err != nil {
//		panic(err)
//	}
//
//	if viper.GetBool(`debug`) {
//		log.Println("Service RUN DEBUG mode")
//	}
//}

func main() {
	Server().Run()
	//	appconfig := config.NewConfig()
	//	cli.MainMenu()
	//	for {
	//		var choice string
	//		fmt.Println("Pilih menu : ")
	//		fmt.Scan(&choice)
	//		switch choice {
	//		case "1":
	//			cli.OrderForm(appconfig.UseCaseManager.OrderUseCase())
	//		case "2":
	//			cli.TableForm(appconfig.UseCaseManager.TableListUseCase())
	//		case "3":
	//			cli.ListFoodForm(appconfig.UseCaseManager.FoodListUseCase())
	//		case "4":
	//			cli.PaymentForm(appconfig.UseCaseManager.PaymentUseCase())
	//		}
	//	}
}
