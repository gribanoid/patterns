package pkg

type HPFactory struct {
}

func (factory *HPFactory) GetComputer() Computer {
	return &HPComputer{Memory: 4, CPU: 1}
}

func (factory *HPFactory) GetMonitor() Monitor {
	return &HPMonitor{Size: 65}
}
