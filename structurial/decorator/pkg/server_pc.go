package pkg

type ServerPC struct {
	CPU     int
	Memory  int
	Wrapper Wrapper
}

func (pc ServerPC) GetPrice() float64 {
	return pc.Wrapper.GetPrice() * float64(pc.CPU) * float64(pc.Memory)
}
