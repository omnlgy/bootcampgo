package latihan2

import "fmt"

type PaymentProcessor interface {
	Pay(amount float64) error
	GetMethodeName() string
}

type CreaditCard struct {
	Number   string
	Provider string
}

func (c CreaditCard) Pay(amount float64) error {
	if c.Number == "" || len(c.Number) < 16 {
		return fmt.Errorf("credit card number is invalid")
	}
	fmt.Printf("[PROSES] Memotong limit kartu kredit %s sebesar Rp %.2f\n", c.Number, amount)
	return nil
}

func (c CreaditCard) GetMethodeName() string {
	return fmt.Sprintf("Credit Card %s", c.Provider)
}

type QrisPayment struct {
	QrisId   string
	Provider string
}

func (q QrisPayment) Pay(amount float64) error {
	fmt.Printf("[PROSES] Membayar via QRIS %s sebesar Rp %.2f\n", q.Provider, amount)
	return nil
}

func (q QrisPayment) GetMethodeName() string {
	return fmt.Sprintf("QRIS %s", q.Provider)
}

type Cart struct {
	CustomerId string
	Items      map[string]float64
}

func (c Cart) New(customerId string) Cart {
	return Cart{
		CustomerId: customerId,
		Items:      make(map[string]float64),
	}
}

func (c *Cart) AddItem(name string, price float64) {
	if _, exists := c.Items[name]; exists {
		return
	}
	c.Items[name] = price
}

func (c *Cart) CheckoutItems(paymentMethod PaymentProcessor) error {
	if len(c.Items) == 0 {
		return fmt.Errorf("cart is empty")
	}

	var total float64

	fmt.Println("\nRincian Keranjang Belanja:")
	for name, price := range c.Items {
		fmt.Printf("  - %s: Rp %.2f\n", name, price)
		total += price
	}
	fmt.Printf("Total Tagihan : Rp %12.2f\n", total)
	fmt.Printf("Metode Pembayaran : %s\n", paymentMethod.GetMethodeName())

	err := paymentMethod.Pay(total)
	if err != nil {
		return fmt.Errorf("pembayaran via %s gagal: %w", paymentMethod.GetMethodeName(), err)
	}

	c.Items = make(map[string]float64)

	return nil
}

func Main() {

	myCart := Cart{}.New("123")
	myCart.AddItem("Buku", 50000)
	myCart.AddItem("Pensil", 2000)
	myCart.AddItem("Penghapus", 1000)

	fmt.Println("==================================================")
	fmt.Println(" APLIKASI CHECKOUT MULTI-GATEWAY ")
	fmt.Println("==================================================")

	creditCard := CreaditCard{
		Number:   "1234567890123456",
		Provider: "BCA",
	}

	err := myCart.CheckoutItems(creditCard)
	if err != nil {
		fmt.Println(err)
	}

	qrisPayment := QrisPayment{
		QrisId:   "1234567890",
		Provider: "GoPay",
	}

	err = myCart.CheckoutItems(qrisPayment)
	if err != nil {
		fmt.Println(err)
	}

}
