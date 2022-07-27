package main

import (
	"fmt"
	"patterns/behavioral/state/pkg"
)

func main() {
	vendingMachine := pkg.NewVendingMachine(1, 10)
	err := vendingMachine.RequestItem()
	if err != nil {
		panic(err)
	}
	err = vendingMachine.InsertMoney(10)
	if err != nil {
		panic(err)
	}
	err = vendingMachine.DispenseItem()
	if err != nil {
		panic(err)
	}
	err = vendingMachine.AddItem(2)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	err = vendingMachine.RequestItem()
	if err != nil {
		panic(err)
	}
	err = vendingMachine.InsertMoney(10)
	if err != nil {
		panic(err)
	}
	err = vendingMachine.DispenseItem()
	if err != nil {
		panic(err)
	}
}
