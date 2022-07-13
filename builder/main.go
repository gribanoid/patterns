package main

import "patterns/builder/pkg"

func main() {
	asusCollector := pkg.GetCollector("asus")
	hpCollector := pkg.GetCollector("hp")
	factory := pkg.NewFactory(asusCollector)
	asusPC := factory.CreateComputer()

	factory.SetCollector(hpCollector)
	hpPC := factory.CreateComputer()

	asusPC.PrintDetails()
	hpPC.PrintDetails()
}
