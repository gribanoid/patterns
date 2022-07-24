package pkg

type HomePC struct {
	CPU     int
	Wrapper Wrapper
}

func (pc HomePC) GetPrice() float64 {
	return pc.Wrapper.GetPrice() * float64(pc.CPU)
}
