package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	var cartItems []string
	cartItems = []string{
		"Laptop Asus ROG | 25000000 | 1 | ready",
		"Mouse Wireless | 350000 | 2 | ready",
		"Smartphone Samsung | 12000000 | 1 | out_of_stock",
		"T-Shirt Basic | 150000 | 3 | ready",
	}
	fmt.Println("==================================================")
	fmt.Println(" MEMULAI PROSES CHECKOUT KERANJANG ")
	fmt.Println("==================================================")

	var totalBelanja float64

checkoutProcessing:
	for idx, item := range cartItems {
		parts := strings.Split(item, " | ")

		productName := parts[0]
		var productPrice float64
		fmt.Sscanf(parts[1], "%f", &productPrice)
		productQuantity, _ := strconv.ParseInt(parts[2], 10, 64)
		productStatus := parts[3]

		fmt.Printf("\n[Item #%d] Memproses: %s x%d\n", idx+1, productName, productQuantity)

		switch productStatus {
		case "out_of_stock":
			fmt.Printf("[Item #%d] ERROR: Stok habis untuk %s\n", idx+1, productName)
			continue checkoutProcessing
		case "ready":
			subTotal := productPrice * float64(productQuantity)
			totalBelanja += subTotal

			if subTotal > 10000000.0 {
				fmt.Println(" > Peringatan: Transaksi Risiko Tinggi Terdeteksi")
				fmt.Println(" > Memulai prosedur verifikasi OTP Keamanan")

				maxRetries := 3
				isOtpVerified := false

				for attempt := 1; attempt <= maxRetries; attempt++ {
					fmt.Printf(" > Verifikasi OTP (Percobaan %d/%d): ", attempt, maxRetries)

					if attempt == maxRetries {
						isOtpVerified = true
					}

					time.Sleep(500 * time.Millisecond)

					if isOtpVerified {
						fmt.Println("[OK] Verifikasi OTP Sukses")
						break
					} else {
						fmt.Println("[FAIL] Verifikasi OTP Gagal")
					}
				}

				if !isOtpVerified {
					fmt.Println("\n > FRAUD DETECTED: Verifikasi OTP gagal")
					fmt.Println(" > Membatalkan seluruh transaksi")
					totalBelanja = 0
					break checkoutProcessing
				}
			}
		}

	}

	fmt.Printf("\n==================================================")
	fmt.Printf("\n TOTAL BELANJA: Rp %.2f", totalBelanja)
	fmt.Println("\n==================================================")

	totalPoint := int((totalBelanja / 100000.0) * 10)
	fmt.Printf("\n TOTAL POIN: %d", totalPoint)
	fmt.Println("\n==================================================")

	switch {
	case totalBelanja > 20000000:
		fmt.Println("Anda mendapatkan Cashback Rp.1000.000")
	case totalBelanja > 5000000:
		fmt.Println("Anda mendapatkan Cashback Rp.200.000")
	default:
		fmt.Println("Anda tidak mendapatkan Cashback")
	}

	userMembership := "Premium"
	cartItems = []string{
		"Laptop Asus ROG | Elektronik | 25000000",
		"Mouse Wireless | Elektronik | 350000",
		"T-Shirt Basic | Fashion | 150000",
		"Sepatu Running | Fashion | 800000",
		"Kopi Arabika | Grocery | 100000",
		"Air Mineral 600ml | Grocery | 5000",
	}

	fmt.Println("==================================================")
	fmt.Println(" MEMULAI PROSES CHECKOUT KERANJANG ")
	fmt.Println("==================================================")

	totalBelanja = 0

	for idx, item := range cartItems {
		parts := strings.Split(item, " | ")
		productName := parts[0]
		productCategory := parts[1]
		productPrice, _ := strconv.ParseInt(parts[2], 10, 64)

		var finalPrice int64

		switch productCategory {
		case "Elektronik":
			finalPrice = productPrice * 110 / 100
		case "Fashion":
			finalPrice = productPrice * 120 / 100
		case "Grocery":
			finalPrice = productPrice * 105 / 100
		}

		totalBelanja += float64(finalPrice)
		fmt.Printf("Item %d: %s (Harga: Rp %d, Diskon: %s, Harga Akhir: Rp %d)\n", idx+1, productName, productPrice, productCategory, finalPrice)
	}

	fmt.Printf("\n==================================================")
	fmt.Printf("\n TOTAL BELANJA: Rp %.2f", float64(totalBelanja))
	fmt.Println("\n==================================================")

	switch userMembership {
	case "Premium":
		totalBelanja = float64(totalBelanja) * 0.95
		fmt.Println("Anda mendapatkan tambahan diskon 5% karena membership Premium")
		fmt.Printf("Total belanja setelah diskon: Rp %.2f\n", float64(totalBelanja))
	case "Gold":
		totalBelanja = float64(totalBelanja) * 0.90
		fmt.Println("Anda mendapatkan tambahan diskon 10% karena membership Gold")
		fmt.Printf("Total belanja setelah diskon: Rp %.2f\n", float64(totalBelanja))
	}

	if totalBelanja > 200000 {
		totalBelanja = float64(totalBelanja) * 0.95
		fmt.Println("Selamat! Anda mendapatkan Gratis Ongkir")
	}

	fmt.Printf("\n==================================================")
	fmt.Printf("\n Flash Sale Smartphone 10 unit")
	fmt.Println("\n==================================================")

	stock := 10
	user := []string{
		"Andi",
		"Budi",
		"Citra",
		"Dewi",
		"Eko",
		"Fina",
		"Gina",
		"Hadi",
		"Ika",
		"Joni",
		"Kania",
		"Lina",
		"Mira",
		"Nina",
		"Oki",
		"Pina",
		"Qina",
		"Rina",
		"Sina",
		"Tina",
	}
flashSale:
	for stock >= 0 {
		for _, name := range user {
			// make random no 1-100
			responseTime := rand.Intn(100) + 1
			if stock > 0 && responseTime < 50 {
				fmt.Printf("%s mendapatkan smartphone\n", name)
				stock--
				continue
			} else if responseTime < 50 && stock <= 0 {
				fmt.Printf("Sold Out: %s kehabisan stok!\n", name)
				break flashSale
			}
		}
	}
}
