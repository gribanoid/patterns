package pkg

type AsusCollector struct {
	Core   int
	Brand  string
	Memory int
}

func (collector *AsusCollector) SetCore() Collector {
	collector.Core = 16
	return collector
}

func (collector *AsusCollector) SetBrand() Collector {
	collector.Brand = "asus"
	return collector
}

func (collector *AsusCollector) SetMemory() Collector {
	collector.Memory = 256
	return collector
}

func (collector *AsusCollector) GetComputer() Computer {
	return Computer{
		Core:   collector.Core,
		Memory: collector.Memory,
		Brand:  collector.Brand,
	}
}
