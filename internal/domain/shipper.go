package domain

type Shipper interface {
	CalculateCost() float64
	GetCourierName() string
}
