package pkg

import "fmt"

type PC struct {
	Name        string
	Description string
	Components  []Component
}

func (pc PC) Search(name string) {
	if pc.Name == name {
		fmt.Printf("Компонент [%s] найден %s\n", pc.Name, pc.Description)
		return
	}
	for _, component := range pc.Components {
		component.Search(name)
	}
}

func (pc PC) GetName() string {
	return pc.Name
}
