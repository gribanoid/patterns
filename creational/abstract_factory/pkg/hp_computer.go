package pkg

import "fmt"

type HPComputer struct {
	Memory int
	CPU    int
}

func (pc *HPComputer) PrintDetails() {
	fmt.Printf("hp cpu:%d memory:%d\n", pc.CPU, pc.Memory)
}
