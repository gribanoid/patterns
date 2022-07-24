package pkg

import "fmt"

type Factory interface {
	GetComputer() Computer
	GetMonitor() Monitor
}

func GetFactory(brand string) (Factory, error) {
	switch brand {
	case Asus:
		return &AsusFactory{}, nil
	case HP:
		return &HPFactory{}, nil
	default:
		return nil, fmt.Errorf("brand %s not found", brand)
	}
}
