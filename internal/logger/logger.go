package logger

import "log"

func Init() {
	log.Println("[AUDIT SYSTEM] Menginisialisasi modul pencatatan transaksi terpusat.")
}

func LogTransaction(orderID string, amount float64) {
	log.Printf("[TRANSACTION] OrderID: %s, Amount: %.2f", orderID, amount)
}
