package pkg

import "fmt"

type GraphicsCard struct {
	Name        string
	Description string
}

func (gc GraphicsCard) Search(name string) {
	if gc.Name == name {
		fmt.Printf("Компонент [%s] найден %s\n", gc.Name, gc.Description)
	}
}

func (gc GraphicsCard) GetName() string {
	return gc.Name
}
