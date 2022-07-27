package main

import (
	"fmt"
	"patterns/behavioral/snapshot/pkg"
)

func main() {
	storage := &pkg.Guardian{
		Items: make([]*pkg.Snapshot, 0),
	}
	creator := &pkg.Creator{State: "A"}
	fmt.Printf("Creator current state %s\n", creator.GetState())
	storage.Add(creator.Create())

	creator.SetState("B")
	fmt.Printf("Creator current state %s\n", creator.GetState())
	storage.Add(creator.Create())

	creator.SetState("C")
	fmt.Printf("Creator current state %s\n", creator.GetState())
	storage.Add(creator.Create())

	creator.Restore(storage.Get(1))
	fmt.Printf("Restored state %s\n", creator.GetState())

	creator.Restore(storage.Get(0))
	fmt.Printf("Restored state %s\n", creator.GetState())

}
