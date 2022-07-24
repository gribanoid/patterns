package pkg

import "fmt"

type Singleton struct {
	Type string
}

func (s *Singleton) Print() {
	fmt.Println("Type", s.Type)
}

func NewSingleton(item *Singleton, typeName string) *Singleton {
	if nil == item {
		return &Singleton{typeName}
	}
	fmt.Println("Type", item.Type, "уже создан!")
	return item
}
