package pkg

import "fmt"

const (
	ServerType           = "server"
	PersonalComputerType = "personal"
	NotebookType         = "notebook"
)

type Computer interface {
	GetType() string
	PrintDetails()
}

func New(typeName string) Computer {
	switch typeName {
	case ServerType:
		return NewServer()
	case PersonalComputerType:
		return NewPC()
	case NotebookType:
		return NewNotebook()
	default:
		fmt.Printf("%s Несуществующий объект\n", typeName)
		return nil
	}
}
