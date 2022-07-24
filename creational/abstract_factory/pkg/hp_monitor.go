package pkg

import "fmt"

type HPMonitor struct {
	Size int
}

func (pc *HPMonitor) PrintDetails() {
	fmt.Printf("hp monitor size:%d\n", pc.Size)
}
