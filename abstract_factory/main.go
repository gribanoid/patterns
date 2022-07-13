package main

import "patterns/abstract_factory/pkg"

var brands = []string{"asus", "hp", "huawei"}

func main() {
	for _, brand := range brands {
		factory, err := pkg.GetFactory(brand)
		if err != nil {
			continue
		}
		factory.GetComputer().PrintDetails()
		factory.GetMonitor().PrintDetails()
	}
}
