package pkg

import "fmt"

type Motherboard struct {
	Name        string
	Description string
	Components  []Component
}

func (mb Motherboard) Search(name string) {
	if mb.Name == name {
		fmt.Printf("Компонент [%s] найден %s\n", mb.Name, mb.Description)
		return
	}
	for _, component := range mb.Components {
		component.Search(name)
	}
}

func (mb Motherboard) GetName() string {
	return mb.Name
}
