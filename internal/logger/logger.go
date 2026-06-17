package logger

import (
	"log"

	"github.com/fatih/color"
)

var (
	greenPrint  = color.New(color.FgGreen, color.Bold).PrintfFunc()
	redPrint    = color.New(color.FgRed, color.Bold).PrintfFunc()
	yellowPrint = color.New(color.FgYellow, color.Bold).PrintfFunc()
)

func Init() {
	yellowPrint("[AUDIT SYSTEM] Menginisialisasi modul pencatatan transaksi terpusat.\n")
}

func LogTransaction(orderID string, amount float64) {
	log.Printf("[TRANSACTION] OrderID: %s, Amount: %.2f", orderID, amount)
}

func LogError(err error) {
	redPrint("[ERROR] %v\n", err)
}

func LogSuccess(msg string) {
	greenPrint("[SUCCESS] %s\n", msg)
}
