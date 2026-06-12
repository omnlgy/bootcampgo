package latihan3

import (
	"errors"
	"fmt"
)

type DiscountRate struct {
	Rate float64
}

type DiscountManager struct {
	Discounts map[string]DiscountRate
}

func NewDiscountManager() *DiscountManager {
	return &DiscountManager{
		Discounts: make(map[string]DiscountRate),
	}
}

func (dm *DiscountManager) RegisterCoupon(coupon string, rate float64) error {
	if rate < 0 || rate > 1 {
		return ErrRateInvalid
	}
	dm.Discounts[coupon] = DiscountRate{Rate: rate}

	return nil
}

func (dm *DiscountManager) CalculateDiscount(coupon string, price float64) (float64, error) {
	if price < 0 {
		return 0, errors.New("price cannot be negative")
	}
	discountRate, exists := dm.Discounts[coupon]
	if !exists {
		return 0, ErrCouponNotFound
	}
	return price * discountRate.Rate, nil
}

var ErrCouponNotFound = errors.New("coupon not found")
var ErrRateInvalid = errors.New("rate invalid")

func Main() {
	fmt.Println("Latihan 3")

	dm := NewDiscountManager()
	dm.RegisterCoupon("DISKON10", 0.1)
	dm.RegisterCoupon("DISKON20", 0.2)
	dm.RegisterCoupon("DISKON150", 1.5)

	originalPrice := 15000.0

	priceAfterDiscount, err := dm.CalculateDiscount("DISKON10", originalPrice)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Original Price: %.2f, Price after discount (10%%): %.2f with coupon DISKON10\n", originalPrice, priceAfterDiscount)

	priceAfterDiscount, err = dm.CalculateDiscount("DISKON30", originalPrice)
	if err != nil {
		if errors.Is(err, ErrCouponNotFound) {
			fmt.Println("Coupon not found")
		} else {
			fmt.Println("someting went wrong")
		}
		return
	}
	fmt.Printf("Original Price: %.2f, Price after discount: %.2f\n", originalPrice, priceAfterDiscount)

}
