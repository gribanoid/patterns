package main

import "patterns/structurial/composite/pkg"

var (
	cpu1        = pkg.CPU{Name: "CP-1", Description: "Процессор 1"}
	cpu2        = pkg.CPU{Name: "CP-2", Description: "Процессор 2"}
	card1       = pkg.GraphicsCard{Name: "Radeon", Description: "Видеокарта 1"}
	card2       = pkg.GraphicsCard{Name: "GeForce", Description: "Видеокарта 2"}
	motherboard = pkg.Motherboard{Name: "Gigabyte", Description: "Материнская плата", Components: []pkg.Component{cpu1, cpu2, card1, card2}}
	pc          = pkg.PC{Name: "CPU-1", Description: "ПК", Components: []pkg.Component{motherboard}}
)

func main() {
	pc.Search("CP-1")
}
