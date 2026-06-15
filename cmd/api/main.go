package main

import (
	"fmt"

	"github.com/omnlgy/bootcampgo/internal/auth"
	"github.com/omnlgy/bootcampgo/internal/logger"
	"github.com/omnlgy/bootcampgo/internal/model"
	"github.com/omnlgy/bootcampgo/internal/service"
)

func init() {
	logger.Init()
}

func main() {
	fmt.Println("==================================================")
	fmt.Println(" APLIKASI KASIR GO-SHOP (MODULAR LAYOUT) ")
	fmt.Println("==================================================")

	auth := auth.NewMerchantAuth("merchant-token")

	if !auth.ValidateToken("merchant-token") {
		fmt.Println("Token tidak valid")
		return
	}

	laptop := model.Product{
		SKU:   "SKU-LAP-01",
		Name:  "Laptop Business Pro",
		Price: 12500000.0,
		Stock: 5,
	}
	mouse := model.Product{
		SKU:   "SKU-MOU-02",
		Name:  "Mouse Wireless Silent",
		Price: 250000.0,
		Stock: 20,
	}

	cart := service.NewCartService()

	fmt.Println("[PROSES] Memasukkan Laptop x2 ke keranjang .")
	if err := cart.AddItem(laptop, 2); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("[PROSES] Memasukkan Mouse x5 ke keranjang .")
	if err := cart.AddItem(mouse, 5); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("[PROSES] Memasukkan Laptop x10 ke keranjang .")
	if err := cart.AddItem(laptop, 10); err != nil {
		fmt.Println("Gagal:", err) // Diharapkan mencetak stok tidak cukup
	}

	fmt.Println("\nRincian Akhir Keranjang Belanja:")
	for sku, qty := range cart.Itmes {
		fmt.Printf("- SKU: %-12s | Kuantitas: %2d unit\n", sku, qty)
	}
	fmt.Println("==================================================")
}
