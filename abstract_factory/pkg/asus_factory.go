package pkg

type AsusFactory struct {
}

func (factory *AsusFactory) GetComputer() Computer {
	return &AsusComputer{
		Memory: 8,
		CPU:    4,
	}
}

func (factory *AsusFactory) GetMonitor() Monitor {
	return &AsusMonitor{Size: 34}
}
