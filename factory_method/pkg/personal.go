package pkg

import "fmt"

type Personal struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

func NewPC() Computer {
	return Personal{
		Type:    PersonalComputerType,
		Core:    8,
		Memory:  16,
		Monitor: true,
	}
}
func (pc Personal) GetType() string {
	return pc.Type
}

func (pc Personal) PrintDetails() {
	fmt.Printf("%s Core:[%d] Mem:[%d] Monitor[%v]\n", pc.Type, pc.Core, pc.Memory, pc.Monitor)
}
