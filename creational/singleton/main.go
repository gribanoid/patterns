package main

import "patterns/creational/singleton/pkg"

var singleton *pkg.Singleton

func init() {
	singleton = &pkg.Singleton{Type: "Одиночка"}
}
func main() {
	for i := 0; i < 3; i++ {
		pkg.NewSingleton(singleton, "Create singleton")
	}
}
