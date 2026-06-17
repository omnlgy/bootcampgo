package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/omnlgy/bootcampgo/internal/auth"
	"github.com/omnlgy/bootcampgo/internal/logger"
	"github.com/omnlgy/bootcampgo/internal/model"
	"github.com/omnlgy/bootcampgo/internal/service"
)

func init() {
	logger.Init()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		logger.LogError(err)
	}

	port := os.Getenv("PORT")
	envMode := os.Getenv("ENV_MODE")
	secretKey := os.Getenv("SECRET_KEY")

	fmt.Println("==================================================")
	fmt.Printf(" MEMULAI SERVER GO-SHOP DI PORT %s (%s)\n", port, envMode)
	fmt.Println("==================================================")

	auth := auth.NewMerchantAuth(secretKey)

	if !auth.ValidateToken("your-secret-key-here") {
		logger.LogError(fmt.Errorf("Token tidak valid"))
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

	inventory := map[string]model.Product{
		laptop.SKU: laptop,
		mouse.SKU:  mouse,
	}

	cart := service.NewCartService()

	fmt.Println("[PROSES] Memasukkan Laptop x2 ke keranjang .")
	if err := cart.AddItem(laptop, 2); err != nil {
		logger.LogError(err)
	}

	fmt.Println("[PROSES] Memasukkan Mouse x5 ke keranjang .")
	if err := cart.AddItem(mouse, 5); err != nil {
		logger.LogError(err)
	}

	fmt.Println("[PROSES] Memasukkan Laptop x10 ke keranjang .")
	if err := cart.AddItem(laptop, 10); err != nil {
		logger.LogError(err) // Diharapkan mencetak stok tidak cukup
	}

	fmt.Println("\nRincian Akhir Keranjang Belanja:")
	for sku, qty := range cart.Itmes {
		fmt.Printf("- SKU: %-12s | Kuantitas: %2d unit\n", sku, qty)
	}

	// jneShipper := service.NewJNEShipper(5.0)
	goSendShipper := service.NewGoSendShipper(5.0)
	orderId, totalAmount, err := cart.Checkout(&inventory, goSendShipper)
	if err != nil {
		logger.LogError(err)
		return
	}

	fmt.Printf("\nOrder ID: %s\n", orderId)
	fmt.Printf("Total Amount: Rp %.2f\n", totalAmount)
	fmt.Println("==================================================")
}
