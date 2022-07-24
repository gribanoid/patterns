package pkg

import "fmt"

type AsusComputer struct {
	Memory int
	CPU    int
}

func (pc *AsusComputer) PrintDetails() {
	fmt.Printf("asus cpu:%d memory:%d\n", pc.CPU, pc.Memory)
}
