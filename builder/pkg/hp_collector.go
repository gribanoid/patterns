package pkg

type HPCollector struct {
	Core   int
	Brand  string
	Memory int
}

func (collector *HPCollector) SetCore() Collector {
	collector.Core = 4
	return collector
}

func (collector *HPCollector) SetBrand() Collector {
	collector.Brand = "hp"
	return collector
}

func (collector *HPCollector) SetMemory() Collector {
	collector.Memory = 8
	return collector
}

func (collector *HPCollector) GetComputer() Computer {
	return Computer{
		Core:   collector.Core,
		Memory: collector.Memory,
		Brand:  collector.Brand,
	}
}
