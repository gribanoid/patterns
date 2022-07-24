package pkg

import "fmt"

type CPU struct {
	Name        string
	Description string
}

func (cpu CPU) Search(name string) {
	if cpu.Name == name {
		fmt.Printf("Компонент [%s] найден %s\n", cpu.Name, cpu.Description)
	}
}

func (cpu CPU) GetName() string {
	return cpu.Name
}
