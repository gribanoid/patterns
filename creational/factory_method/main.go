package main

import "patterns/creational/factory_method/pkg"

var types = []string{pkg.NotebookType, pkg.PersonalComputerType, pkg.ServerType, "mono block"}

func main() {
	for _, typeName := range types {
		computer := pkg.New(typeName)
		if computer == nil {
			continue
		}
		computer.PrintDetails()
	}
}
