package latihan1

import "fmt"

type PaymentGateway interface {
	Pay(amount float64) error
}

type BankTransfer struct {
	Provider string
}

func (b BankTransfer) Pay(amount float64) error {
	fmt.Printf("Paying %f using %s\n", amount, b.Provider)
	return nil
}

type EWallet struct {
	Provider string
}

func (ew EWallet) Pay(amount float64) error {
	fmt.Printf("Paying %f using %s\n", amount, ew.Provider)
	return nil
}

func CheckoutOrder(paymentMethode PaymentGateway, total float64) {
	err := paymentMethode.Pay(total)
	if err != nil {
		fmt.Println("Payment failed:", err)
		return
	}
	fmt.Println("Payment successful")
}

func Main() {
	// map dan interface

	// Cara 1: Menggunakan make (Sangat direkomendasikan)
	productPrices := make(map[string]float64)
	productPrices["Laptop"] = 15000000.0
	productPrices["Mouse"] = 350000.0
	// Cara 2: Menggunakan map literal
	categories := map[string]string{
		"ROG": "Laptop Gaming",
		"MX3": "Mouse Productivity",
	}
	fmt.Println(productPrices["Laptop"])
	fmt.Println(categories["ROG"])

	delete(categories, "ROG")
	if _, exists := categories["ROG"]; !exists {
		fmt.Println("ROG not found")
	}

	paymentMethode1 := BankTransfer{Provider: "BCA"}
	paymentMethode2 := EWallet{Provider: "OVO"}

	CheckoutOrder(paymentMethode1, 100000)
	CheckoutOrder(paymentMethode2, 200000)

	var data any

	data = "asd"

	_, ok := data.(string)
	if ok {
		fmt.Println("Data is a string")
	}

	data = 233

	switch data.(type) {
	case string:
		fmt.Println("Data is a string")
	case int:
		fmt.Println("Data is an int")
	default:
		fmt.Println("Type is not string or int")
	}

}
