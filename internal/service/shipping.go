package service

type JNEShipper struct {
	weight float64
}

func NewJNEShipper(weight float64) *JNEShipper {
	return &JNEShipper{
		weight: weight,
	}
}

func (j *JNEShipper) CalculateCost() float64 {
	return j.weight * 1000
}

func (j *JNEShipper) GetCourierName() string {
	return "JNE"
}

type GoSendShipper struct {
	weight float64
}

func NewGoSendShipper(weight float64) *GoSendShipper {
	return &GoSendShipper{
		weight: weight,
	}
}

func (g *GoSendShipper) CalculateCost() float64 {
	return g.weight * 1500
}

func (g *GoSendShipper) GetCourierName() string {
	return "GoSend"
}
