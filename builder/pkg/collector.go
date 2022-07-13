package pkg

const (
	AsusCollectorType = "asus"
	HPCollectorType   = "hp"
)

type Collector interface {
	SetCore() Collector
	SetBrand() Collector
	SetMemory() Collector
	GetComputer() Computer
}

func GetCollector(collectorType string) Collector {
	switch collectorType {
	case AsusCollectorType:
		return &AsusCollector{}
	case HPCollectorType:
		return &HPCollector{}
	default:
		return nil
	}
}
