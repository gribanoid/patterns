package pkg

import "fmt"

type Computer struct {
	Core   int
	Brand  string
	Memory int
}

func (pc *Computer) PrintDetails() {
	fmt.Printf("%s Core:[%d] Mem:[%d]\n", pc.Brand, pc.Core, pc.Memory)
}
